package deepseek

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
	"github.com/pkg/errors"
	"regexp"
)

//go:generate mockgen -destination=./mock/ds.go github.com/go-deepseek/deepseek Client

type Client struct {
	ctx    context.Context
	client deepseek.Client
}

func NewDSClient(ctx context.Context, apiKey string) (*Client, error) {
	client, err := deepseek.NewClient(apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "newGigaClient error")
	}

	return &Client{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *Client) GetSummary(text string) (*Summary, error) {
	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "system",
				Content: "Ты — языковая модель, анализирующая сообщения из чата",
			},
			{
				Role:    "user",
				Content: c.prompt(text),
			},
		},
	}

	chatResp, err := c.client.CallChatCompletionsChat(c.ctx, chatReq)
	if err != nil {
		return nil, errors.Wrap(err, "CallChatCompletionsChat error")
	}

	return c.postProcessing(chatResp.Choices[0].Message.Content)
}

func (c *Client) prompt(text string) string {
	return fmt.Sprintf(`Каждое сообщение имеет уникальный messageID, имя автора и текст. Табуляция (\t) обозначает, что сообщение является ответом на другое сообщение (вложенность тредов).
Твоя задача:
Выделить основные темы обсуждений, которые происходили в чате. Для каждой темы:
Дай краткое описание сути обсуждения.
Укажи messageID, с которого тема началась (root_message_id).
Формат вывода: JSON:
{
  "topics": [
    {
      "topic": "Краткое описание темы",
      "root_message_id": "ID сообщения, откуда началась тема"
    },
    ...
  ]
}
🔹 Игнорируй фразы вне обсуждений (шутки, междометия и пр.), если они не формируют отдельную тему.
🔹 Если ветки обсуждений идут параллельно, считаем их разными темами.
🔹 Используй только информацию из переписки, ничего не выдумывай.
🔹 Не выводи сообщения и авторов отдельно — только описание темы и ID, где она началась.

ВОТ СООБЩЕНИЕ:
%s
`, text)
}

func (c *Client) postProcessing(answer string) (*Summary, error) {
	var re = regexp.MustCompile("(?s)```json(.*)```")
	if sm := re.FindStringSubmatch(answer); len(sm) > 1 {
		answer = sm[1]
	}

	var s Summary
	if err := json.Unmarshal([]byte(answer), &s); err != nil {
		return nil, errors.Wrap(err, "json unmarshal error")
	}

	return &s, nil
}
