package bot

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"strings"
	"summary-bot/giga"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock.go

type IAdapter interface {
	AppendTimeData(key string, t time.Time, data []byte) error
	GetMessageData(key string, tstart, tfinish time.Time) ([][]byte, error)
	GetMessageDataForClear(key string, tfinish time.Time) ([][]byte, error)
	DeleteMessageDataByTime(key string, t time.Time) error
	DeleteMessageData(key string, t time.Time, data []byte) error
}
type AI interface {
	GetSummary(text string) (string, error)
}

type SummaryBot struct {
	adapter  IAdapter
	botAPI   *tgbotapi.BotAPI
	wdUpdate tgbotapi.UpdatesChannel
	logger   *slog.Logger
	ai       AI
}

type storeMsgData struct {
	MessageID        int             `json:"messageID"`
	ReplyToMessageID int             `json:"replyToMessageID"`
	Txt              string          `json:"txt"`
	DateTime         time.Time       `json:"dateTime"`
	User             string          `json:"user"`
	ChatID           int64           `json:"chatID"`
	answers          []*storeMsgData `json:"-"`
	parent           *storeMsgData   `json:"-"`
}

const (
	msgStoreKey = "messages"
	storageDays = 2
)

func NewSummaryBot(botToken, authKey string, adapter IAdapter) (*SummaryBot, error) {
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	// bot.Debug = true
	if err != nil {
		return nil, errors.Wrap(err, "create new bot error")
	}

	bot := &SummaryBot{
		botAPI:  botAPI,
		adapter: adapter,
		logger:  slog.New(slog.NewTextHandler(os.Stdout, nil)).With("name", "bot"),
	}

	bot.ai, err = giga.NewGigaClient(context.Background(), authKey)
	if err != nil {
		return nil, errors.Wrap(err, "create gigachat client error")
	}

	botAPI.Request(&tgbotapi.DeleteWebhookConfig{}) // на всякий случай удаляем веб хук

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.AllowedUpdates = []string{"message", "edited_message"}
	bot.wdUpdate = botAPI.GetUpdatesChan(u) // полинг

	return bot, nil
}

func (sb *SummaryBot) Run() {
	sb.logger.Info("bot running")

	go sb.garbageCleaning() // контроль и удаления старых записей в БД

	for update := range sb.wdUpdate {
		if update.EditedMessage != nil {
			if err := sb.editMessage(update.EditedMessage); err != nil {
				sb.logger.Error("store edit message error", "error", err.Error())
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		command := update.Message.Command()
		switch strings.ToLower(command) {
		case "summarytoday":

			sb.summary(sb.getMessages(update.Message.Chat.ID), update.Message.Chat.ID)
			sb.DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			continue
		case "summarythread":
			if update.Message.ReplyToMessage != nil {
				sb.summary(sb.getThreadMessages(update.Message.ReplyToMessage), update.Message.Chat.ID)
			} else {
				sb.SendMessage("Данную команду необходимо использовать в ответ на любое сообщение из треда", update.Message.Chat.ID, update.Message.MessageID, time.Second*30)
			}

			sb.DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			continue
		}

		if err := sb.storeMessage(update.Message); err != nil {
			sb.logger.Error(err.Error())
		}
	}
}

func (sb *SummaryBot) summary(rawTxt string, chatID int64) {
	summary, err := sb.ai.GetSummary(rawTxt)
	if err != nil {
		sb.logger.Error(errors.Wrap(err, "get summary error").Error())
		sb.SendMessage("ой, произошла ошибка. Нужно смотреть логи.", chatID, 0, time.Second*30)
	} else {
		sb.SendMessage(fmt.Sprintf("Краткий пересказ переписки:\n\n%s", summary), chatID, 0, 0)
	}
}

func (sb *SummaryBot) editMessage(msg *tgbotapi.Message) error {
	bdata, err := sb.getDataForStore(msg)
	if err != nil {
		return err
	}

	if err := sb.updateMessageData(msg.Time(), msg.MessageID, bdata); err != nil {
		return err
	}

	return nil
}

func (sb *SummaryBot) updateMessageData(t time.Time, messageID int, data []byte) error {
	// получаем сообщения, смотрим ID, удаляем, добавляем новое
	rows, err := sb.adapter.GetMessageDataForClear(msgStoreKey, t)
	if err != nil {
		return err
	}

	for _, row := range rows {
		var sdata storeMsgData

		if err := json.Unmarshal(row, &sdata); err != nil {
			sb.logger.Error(errors.Wrap(err, "json unmarshal error").Error())
			continue
		}

		if sdata.MessageID == messageID {
			sb.adapter.DeleteMessageData(msgStoreKey, t, row)
			sb.adapter.AppendTimeData(msgStoreKey, t, data)
		}
	}

	return nil
}

func (sb *SummaryBot) storeMessage(msg *tgbotapi.Message) error {
	bdata, err := sb.getDataForStore(msg)
	if err != nil {
		return err
	}

	return sb.adapter.AppendTimeData(msgStoreKey, msg.Time(), bdata)
}

func (sb *SummaryBot) getDataForStore(msg *tgbotapi.Message) ([]byte, error) {
	text := msg.Text
	if msg.Document != nil {
		text = "<document>"
	}
	if msg.Voice != nil {
		text = "<voice>"
	}
	if msg.Photo != nil {
		text = fmt.Sprintf("<photo> %s", msg.Caption)
	}
	if msg.Poll != nil {
		text = "<poll>"
	}
	if msg.Animation != nil {
		text = fmt.Sprintf("<gif> %s", msg.Caption)
	}

	if strings.TrimSpace(text) == "" {
		return nil, errors.New("text message is empty")
	}

	sdata := &storeMsgData{
		MessageID: msg.MessageID,
		Txt:       text,
		DateTime:  msg.Time().UTC(),
		User:      msg.From.UserName,
		ChatID:    msg.Chat.ID,
	}

	if msg.From.FirstName != "" || msg.From.LastName != "" {
		sdata.User = strings.TrimSpace(msg.From.FirstName + " " + msg.From.LastName)
	}

	if msg.ReplyToMessage != nil {
		sdata.ReplyToMessageID = msg.ReplyToMessage.MessageID
	}

	bdata, err := json.Marshal(sdata)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal error")
	}

	return bdata, nil
}

func (sb *SummaryBot) getMessages(chatID int64) string {
	list, graph := sb.buildGraph(chatID)
	return sb.printGraph(list, graph, nil)
}

func (sb *SummaryBot) getThreadMessages(msg *tgbotapi.Message) string {
	if msg == nil {
		return ""
	}

	list, graph := sb.buildGraph(msg.Chat.ID)

	// определяем корневое сообщение и от него получаем весь тред
	root := graph[msg.MessageID]
	for m, ok := graph[msg.MessageID]; ok && m.parent != nil; m, ok = graph[m.parent.MessageID] {
		root = m.parent
	}

	return sb.printGraph(list, graph, root)
}

func (sb *SummaryBot) printGraph(list []int, graph map[int]*storeMsgData, root *storeMsgData) string {
	var result strings.Builder

	sb.printGraphHelper(0, &result, list, graph, root)
	return result.String()
}

func (sb *SummaryBot) printGraphHelper(level int, result *strings.Builder, list []int, graph map[int]*storeMsgData, root *storeMsgData) {
	for _, msgID := range list {
		msg, ok := graph[msgID]
		if !ok {
			continue
		}

		if root != nil && msgID != root.MessageID && level == 0 {
			delete(graph, msgID)
			continue
		}

		delete(graph, msgID)

		result.WriteString(fmt.Sprintf("%s%s: %s", strings.Repeat("\t", level), msg.User, msg.Txt))
		result.WriteString("\n")

		answersMsgID := make([]int, len(msg.answers))

		for i, answer := range msg.answers {
			answersMsgID[i] = answer.MessageID
		}

		if len(answersMsgID) > 0 {
			sb.printGraphHelper(level+1, result, answersMsgID, graph, root)
		}

		// вернулись, значит обошли все подчиненные узлы, выходим
		if root != nil && msgID == root.MessageID {
			break
		}
	}
}

func (sb *SummaryBot) buildGraph(chatID int64) (resultList []int, resultMap map[int]*storeMsgData) {
	resultMap = make(map[int]*storeMsgData)

	data, err := sb.adapter.GetMessageData(msgStoreKey, startDay(time.Now().UTC()), endDay(time.Now().UTC()))
	if err != nil {
		return
	}

	for _, item := range data {
		var sdata storeMsgData

		if err := json.Unmarshal(item, &sdata); err != nil {
			sb.logger.Error(errors.Wrap(err, "json unmarshal error").Error())
			return
		}

		if sdata.ChatID != chatID {
			continue
		}

		resultList = append(resultList, sdata.MessageID)
		resultMap[sdata.MessageID] = &sdata
		if sdata.ReplyToMessageID != 0 {
			if parent, ok := resultMap[sdata.ReplyToMessageID]; ok {
				parent.answers = append(parent.answers, &sdata)
				resultMap[sdata.MessageID].parent = parent
			}
		}
	}

	return
}

func (sb *SummaryBot) getSummary() string {
	return ""
}

func (sb *SummaryBot) garbageCleaning() {
	for {
		sb.logger.Info("start GC")

		data, err := sb.adapter.GetMessageDataForClear(msgStoreKey, startDay(time.Now().UTC()).Add(-time.Second))
		if err != nil {
			sb.logger.Error(errors.Wrap(err, "garbage cleaning error").Error())
		}

		for _, item := range data {
			var sdata storeMsgData

			if err := json.Unmarshal(item, &sdata); err != nil {
				sb.logger.Error(errors.Wrap(err, "json unmarshal error").Error())
				continue
			}

			if time.Now().UTC().Sub(sdata.DateTime).Hours()/24 > storageDays {
				sb.logger.Info(fmt.Sprintf("deleted message from DB, message data %v", sdata.DateTime))
				if err := sb.adapter.DeleteMessageDataByTime(msgStoreKey, sdata.DateTime); err != nil {
					sb.logger.Error(errors.Wrap(err, "DeleteMessageData error").Error())
				}
			}

		}

		time.Sleep(time.Hour)
	}
}

func (sb *SummaryBot) SendMessage(txt string, chatID int64, replyTo int, ttl time.Duration) error {
	if txt == "" {
		return errors.New("text message is empty")
	}

	newmsg := tgbotapi.NewMessage(chatID, txt)
	newmsg.ParseMode = "HTML"
	if replyTo > 0 {
		newmsg.ReplyToMessageID = replyTo
	}

	msg, err := sb.botAPI.Send(newmsg)
	go func() {
		if ttl > 0 {
			time.Sleep(ttl)
			sb.DeleteMessage(chatID, msg.MessageID)
		}
	}()

	return err
}

func (sb *SummaryBot) DeleteMessage(chatID int64, messageID int) error {
	conf := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}

	_, err := sb.botAPI.Request(conf)
	return err
}

func startDay(now time.Time) time.Time {
	return now.Truncate(24 * time.Hour)
}

func endDay(now time.Time) time.Time {
	return now.Add(time.Hour * 24).Truncate(24 * time.Hour).Add(-time.Second)
}
