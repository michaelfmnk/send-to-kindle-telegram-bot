# SendToKindle [![Go](https://github.com/michaelfmnk/SendToKindleTelegramBot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/michaelfmnk/SendToKindleTelegramBot/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/michaelfmnk/send-to-kindle-telegram-bot)](https://goreportcard.com/report/github.com/michaelfmnk/send-to-kindle-telegram-bot)

## Building the Docker Container

To build the Docker container, execute the following command:

```shell
docker build -t sendtokindle .
```

## Configuring Environment Variables

Before starting the bot, you need to configure the following environment variables:

| Variable            | Description                                         |
|---------------------|-----------------------------------------------------|
| UBOT_TELEGRAM_TOKEN | Telegram bot token                                  |
| UBOT_EMAIL_FROM     | Email address that the bot will use to send books   |
| UBOT_PASSWORD       | Email password                                      |
| UBOT_EMAIL_TO       | Kindle email address to which books will be sent    |
| UBOT_SMTP_HOST      | SMTP mail host                                      |

## Usage

After starting the bot and configuring the necessary environment variables, you can use it by sending a message to the bot containing the document you want to convert. The bot will then send the converted document to your Kindle email address.

## Links

For more information, you can check out the following links:

- [Dev.to post](https://dev.to/michaelfmnk/developing-send-to-kindle-telegram-bot-120c)
- If you have any questions or suggestions, feel free to message Michael at [michael@fomenko.dev](mailto:michael@fomenko.dev).
