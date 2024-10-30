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
				Content: c.prompt(),
			},
			{
				Role:    "user",
				Content: text,
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
	return fmt.Sprintf("Ты нейросеть предназначеная для того что бы анализировать длинный текст и выдавать краткое резюме. \n" +
		"Cоздай краткое резюме переписки из чата, о чем люди общались, выдели основные темы.\n" +
		"Структура такая <имя пользователя>:<сообщение> табуляция определяет темы разговора (кто на какое сообщение отвечал)\n" +
		"ПРИМЕР ПЕРЕПИСКИ: \n" +
		"Иванов Иван: че в бсп нету методов чтобы добавить или убавить отборы?\n" +
		"\tМаша: нет")
	//"Краткое резюме необходимо предоставить на языке %s")
}

func ptr[T any](v T) *T {
	return &v
}
