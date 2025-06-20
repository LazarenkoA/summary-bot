package giga

import (
	"context"
	"fmt"
	"github.com/paulrzcz/go-gigachat"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock.go
type IGigaClient interface {
	AuthWithContext(ctx context.Context) error
	ChatWithContext(ctx context.Context, in *gigachat.ChatRequest) (*gigachat.ChatResponse, error)
}

type Client struct {
	ctx    context.Context
	client IGigaClient
}

func NewGigaClient(ctx context.Context, authKey string) (*Client, error) {
	client, err := gigachat.NewInsecureClientWithAuthKey(authKey)
	if err != nil {
		return nil, errors.Wrap(err, "newGigaClient error")
	}

	return &Client{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *Client) GetSummary(text string) (string, error) {
	err := c.client.AuthWithContext(c.ctx)
	if err != nil {
		return "", errors.Wrap(err, "auth error")
	}

	if text == "" {
		return "", errors.New("message is not defined")
	}

	req := &gigachat.ChatRequest{
		Model:  "GigaChat",
		Stream: ptr(false),
		Messages: []gigachat.Message{
			{
				Role:    "system",
				Content: "Ты анализируешь переписку людей в чате.",
			},
			{
				Role:    "user",
				Content: c.prompt(),
			},
		},
	}

	resp, err := c.client.ChatWithContext(c.ctx, req)
	if err != nil {
		return "", errors.Wrap(err, "request error")
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("response does not contain data")
	}

	return resp.Choices[0].Message.Content, nil
}

func (c *Client) prompt() string {
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
`)
}

func ptr[T any](v T) *T {
	return &v
}
