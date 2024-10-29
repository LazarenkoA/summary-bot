package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	mock_bot "summary-bot/bot/mock"
	"testing"
)

func Test_getDataForStore(t *testing.T) {
	sb := &SummaryBot{}

	t.Run("text message is empty", func(t *testing.T) {
		msg := &tgbotapi.Message{}

		data, err := sb.getDataForStore(msg)
		assert.Nil(t, data)
		assert.EqualError(t, err, "text message is empty")
	})
	t.Run("test1", func(t *testing.T) {
		msg := &tgbotapi.Message{
			Text: "test",
			From: &tgbotapi.User{UserName: "user"},
			Chat: &tgbotapi.Chat{},
		}

		data, err := sb.getDataForStore(msg)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, `{"messageID":0,"replyToMessageID":0,"txt":"test","dateTime":"1970-01-01T00:00:00Z","user":"user","chatID":0}`, string(data))
	})
	t.Run("test2", func(t *testing.T) {
		msg := &tgbotapi.Message{
			MessageID: 123,
			Text:      "test",
			From:      &tgbotapi.User{UserName: "user", LastName: "Иванов"},
			Chat:      &tgbotapi.Chat{ID: 321},
		}

		data, err := sb.getDataForStore(msg)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, `{"messageID":123,"replyToMessageID":0,"txt":"test","dateTime":"1970-01-01T00:00:00Z","user":"Иванов","chatID":321}`, string(data))
	})
	t.Run("test3", func(t *testing.T) {
		msg := &tgbotapi.Message{
			Text: "test",
			From: &tgbotapi.User{UserName: "user", FirstName: "Иван", LastName: "Иванов"},
			Chat: &tgbotapi.Chat{},
		}

		data, err := sb.getDataForStore(msg)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, `{"messageID":0,"replyToMessageID":0,"txt":"test","dateTime":"1970-01-01T00:00:00Z","user":"Иван Иванов","chatID":0}`, string(data))
	})
}

func Test_getMessages(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_bot.NewMockIAdapter(c)
	sb := &SummaryBot{
		adapter: db,
	}

	t.Run("error", func(t *testing.T) {
		db.EXPECT().GetMessageData(msgStoreKey, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

		result := sb.getMessages(0)
		assert.Equal(t, "", result)
	})
	t.Run("other chatID", func(t *testing.T) {
		testData := [][]byte{
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 53, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 209, 130, 208, 181, 209, 129, 209, 130, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 48, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 54, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 208, 176, 208, 178, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 50, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 55, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 208, 176, 208, 178, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 50, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 56, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 54, 44, 34, 116, 120, 116, 34, 58, 34, 209, 131, 209, 131, 209, 131, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 54, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 57, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 56, 44, 34, 116, 120, 116, 34, 58, 34, 209, 131, 208, 186, 209, 131, 208, 186, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 57, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 49, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 57, 44, 34, 116, 120, 116, 34, 58, 34, 103, 103, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 50, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 50, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 56, 44, 34, 116, 120, 116, 34, 58, 34, 103, 102, 103, 102, 103, 103, 103, 103, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 55, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 51, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 103, 102, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 56, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 52, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 103, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 57, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 53, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 55, 44, 34, 116, 120, 116, 34, 58, 34, 103, 102, 103, 102, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 49, 51, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 50, 57, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 50, 44, 34, 116, 120, 116, 34, 58, 34, 99, 99, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 48, 58, 50, 51, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 48, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 50, 57, 44, 34, 116, 120, 116, 34, 58, 34, 99, 99, 99, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 48, 58, 50, 54, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 49, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 50, 44, 34, 116, 120, 116, 34, 58, 34, 50, 50, 50, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 48, 58, 51, 49, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 52, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 115, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 49, 58, 48, 54, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 54, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 52, 44, 34, 116, 120, 116, 34, 58, 34, 119, 119, 119, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 49, 58, 49, 51, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 57, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 101, 119, 119, 119, 101, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 55, 58, 49, 49, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 208, 144, 209, 128, 209, 130, 209, 145, 208, 188, 32, 208, 155, 208, 176, 208, 183, 208, 176, 209, 128, 208, 181, 208, 189, 208, 186, 208, 190, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 52, 49, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 57, 44, 34, 116, 120, 116, 34, 58, 34, 115, 115, 115, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 55, 58, 50, 52, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 208, 144, 209, 128, 209, 130, 209, 145, 208, 188, 32, 208, 155, 208, 176, 208, 183, 208, 176, 209, 128, 208, 181, 208, 189, 208, 186, 208, 190, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
		}
		db.EXPECT().GetMessageData(msgStoreKey, gomock.Any(), gomock.Any()).Return(testData, nil)

		result := sb.getMessages(0)
		assert.Equal(t, "", result)
	})
	t.Run("pass", func(t *testing.T) {
		testData := [][]byte{
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 53, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 209, 130, 208, 181, 209, 129, 209, 130, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 48, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 54, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 208, 176, 208, 178, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 50, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 55, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 208, 176, 208, 178, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 50, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 56, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 54, 44, 34, 116, 120, 116, 34, 58, 34, 209, 131, 209, 131, 209, 131, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 54, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 57, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 56, 44, 34, 116, 120, 116, 34, 58, 34, 209, 131, 208, 186, 209, 131, 208, 186, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 48, 58, 53, 57, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 49, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 57, 44, 34, 116, 120, 116, 34, 58, 34, 103, 103, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 50, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 50, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 56, 44, 34, 116, 120, 116, 34, 58, 34, 103, 102, 103, 102, 103, 103, 103, 103, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 55, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 51, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 103, 102, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 56, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 52, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 103, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 48, 57, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 53, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 56, 57, 55, 44, 34, 116, 120, 116, 34, 58, 34, 103, 102, 103, 102, 103, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 55, 58, 52, 56, 58, 49, 51, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 50, 57, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 50, 44, 34, 116, 120, 116, 34, 58, 34, 99, 99, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 48, 58, 50, 51, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 48, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 50, 57, 44, 34, 116, 120, 116, 34, 58, 34, 99, 99, 99, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 48, 58, 50, 54, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 49, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 48, 50, 44, 34, 116, 120, 116, 34, 58, 34, 50, 50, 50, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 48, 58, 51, 49, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 52, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 115, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 49, 58, 48, 54, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 54, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 52, 44, 34, 116, 120, 116, 34, 58, 34, 119, 119, 119, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 49, 58, 49, 51, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 76, 97, 122, 97, 114, 101, 110, 107, 111, 65, 78, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 57, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 48, 44, 34, 116, 120, 116, 34, 58, 34, 101, 119, 119, 119, 101, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 55, 58, 49, 49, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 208, 144, 209, 128, 209, 130, 209, 145, 208, 188, 32, 208, 155, 208, 176, 208, 183, 208, 176, 209, 128, 208, 181, 208, 189, 208, 186, 208, 190, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
			{123, 34, 109, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 52, 49, 44, 34, 114, 101, 112, 108, 121, 84, 111, 77, 101, 115, 115, 97, 103, 101, 73, 68, 34, 58, 49, 57, 51, 57, 44, 34, 116, 120, 116, 34, 58, 34, 115, 115, 115, 34, 44, 34, 100, 97, 116, 101, 84, 105, 109, 101, 34, 58, 34, 50, 48, 50, 52, 45, 49, 48, 45, 50, 57, 84, 48, 56, 58, 51, 55, 58, 50, 52, 90, 34, 44, 34, 117, 115, 101, 114, 34, 58, 34, 208, 144, 209, 128, 209, 130, 209, 145, 208, 188, 32, 208, 155, 208, 176, 208, 183, 208, 176, 209, 128, 208, 181, 208, 189, 208, 186, 208, 190, 34, 44, 34, 99, 104, 97, 116, 73, 68, 34, 58, 45, 49, 48, 48, 50, 49, 51, 49, 48, 56, 49, 51, 56, 53, 125},
		}
		db.EXPECT().GetMessageData(msgStoreKey, gomock.Any(), gomock.Any()).Return(testData, nil)

		result := sb.getMessages(-1002131081385)
		assert.Equal(t, "LazarenkoAN: тест\nLazarenkoAN: ав\n\tLazarenkoAN: ууу\n\t\tLazarenkoAN: укук\n\t\t\tLazarenkoAN: ggg\n\t\tLazarenkoAN: gfgfggggg\n\t\t\tLazarenkoAN: cc\n\t\t\t\tLazarenkoAN: ccc\n\t\t\tLazarenkoAN: 222\nLazarenkoAN: ав\n\tLazarenkoAN: gfgfg\nLazarenkoAN: gfg\nLazarenkoAN: gg\nLazarenkoAN: s\n\tLazarenkoAN: www\nАртём Лазаренко: ewwwe\n\tАртём Лазаренко: sss\n", result)
	})
}