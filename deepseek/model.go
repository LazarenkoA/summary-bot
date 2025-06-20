package deepseek

type Topic struct {
	Topic         string `json:"topic"`
	RootMessageId string `json:"root_message_id"`
}

type Summary struct {
	Topics []Topic `json:"topics"`
}
