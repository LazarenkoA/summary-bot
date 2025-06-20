module summary-bot

go 1.23.4

toolchain go1.24.1

//replace github.com/paulrzcz/go-gigachat => ../go-gigachat

require (
	github.com/garyburd/redigo v1.6.4
	github.com/go-deepseek/deepseek v0.8.0
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/golang/mock v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/paulrzcz/go-gigachat v0.1.2
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
