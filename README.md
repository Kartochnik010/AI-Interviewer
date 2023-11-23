# AI-interviewer 

## About
Веб-сайт для грейда, который использует ИИ.

Cайт попросит вас ввести информацию о вашем текущем опыте, о вашей желаемой должности. Далее начнется процесс интервью для работника на этой должности будь то Пилот гражданского самолета или разработчика ПО. 

## Requirements

- Go version 1.15

## Usage



1. Клонируйте репозиторий:
```Bash
   git clone https://github.com/Kartochnik010/ai-interviewer.git
```
2. Перейдите в директорию проекта:
```Bash
    cd ai-interviewer
```
3. Скачайте зависимости:
```Go
    go mod tidy
```
4. Заполните `.env` Файл`
```Bash
    # .env
    GPT_TOKEN="sk-******"
    GPT_URL="https://api.openai.com/v1/chat/completions"
    PORT="8080"
    GPT_MAX_TOKENS="256"
    GPT_MODEL="gpt-4"
```
5. Запустите приложение:
```Go
    make run
```

Все готово! Теперь вы можете перейти по ссылке localhost:`$PORT`

## Project structure
```
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── internal
│   ├── api
│   │   └── api.go
│   ├── handler
│   │   ├── handler.go
│   │   └── routes.go
│   └── models
│       ├── gpt.go
│       ├── models.go
│       └── user.go
├── pkg
│   └── server
│       └── server.go
├── ui
│   ├── static
│   │   ├── css
│   │   │   └── style.css
│   │   └── js
│   │       └── script.js
│   ├── templates
│   │   └── home.html
│   └── ui.go
├── utils
│   └── utils.go
├── Dockerfile
├── Makefile
├── README.md
├── exampleRequestWithResponse.txt
├── go.mod
└── go.sum
```
`/cmd` - точка входа в приложение.

`/config` - конфиги для приложения

`/internal/api` - пакет, содержащий функционал для работы с API.

`/internal/models` - пакет, определяющий структуры данных.

`/ui` - пакет, необходимый для поддержки и взаимодействия с веб интефейсом.

`/pkg` - пакеты, которые можно переиспользовать в других проектах.

`/util` - пакет с вспомогательными функциями для обработки данных.

### API
Приложение делает запросы к API [GPT Completions](https://api.openai.com/v1/chat/completions), взаимодействует с моделями. Данные обрабатываются и возвращаются в приложение, после чего показывается на экране у пользователя.

### Error handling
Все сетевые запросы и операции обработки данных сопровождаются проверками на наличие ошибок. В случае возникновения ошибок, приложение предоставляет информативное сообщение об ошибке.

### Contribute to project
## Todos
- [ ] Add storage for roadmap's
- [ ] Add storage for best conversations
- [ ] Export conversation
- [ ] Publish conversation on the web
- [ ] 3rd party logger?