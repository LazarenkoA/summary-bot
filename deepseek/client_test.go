package deepseek

import (
	"context"
	"github.com/go-deepseek/deepseek/response"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_deepseek "summary-bot/deepseek/mock"
	"testing"
)

func Test_GetMessageCharacteristics(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	ds, err := NewDSClient(context.Background(), "123")
	assert.NoError(t, err)

	client := mock_deepseek.NewMockClient(c)
	ds.client = client

	client.EXPECT().CallChatCompletionsChat(gomock.Any(), gomock.Any()).Return(&response.ChatCompletionsResponse{
		Choices: []*response.Choice{
			{
				Message: &response.Message{Content: "```json" +
					"{" +
					"	\"topics\": [" +
					"{" +
					"	\"topic\": \"Обсуждение вложенных транзакций и их поддержки на уровне платформы\"," +
					"	\"root_message_id\": \"158621\"" +
					"}," +
					"{" +
					"	\"topic\": \"Тестирование реакции дипсика на токсичные или агрессивные сообщения\"," +
					"	\"root_message_id\": \"158626\"" +
					"}]}```"},
			},
		},
	}, nil)

	result, err := ds.GetSummary("test")
	assert.NoError(t, err)
	if assert.NotNil(t, result) {
		assert.Len(t, result.Topics, 2)
	}
}
