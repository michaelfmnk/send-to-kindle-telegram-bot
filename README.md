# SendToKindle [![Go](https://github.com/michaelfmnk/SendToKindleTelegramBot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/michaelfmnk/SendToKindleTelegramBot/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/michaelfmnk/send-to-kindle-telegram-bot)](https://goreportcard.com/report/github.com/michaelfmnk/send-to-kindle-telegram-bot)

#### SendToKindle is a self-hosted telegram bot that converts and sends documents to your kindle. It uses calibre for conversion.

---

You can build docker container like this:

```shell
docker build -t sendtokindle .
```

To start it you need to pass the following environment variables:

| Variable            | Description                                         |
|---------------------|-----------------------------------------------------|
| UBOT_TELEGRAM_TOKEN | Telegram bot token                                  |
| UBOT_EMAIL_FROM     | Email from which bot will send converted books      |
| UBOT_PASSWORD       | Email password                                      |
| UBOT_EMAIL_TO       | Kindle email. Bot will send there converted e-books |
| UBOT_SMTP_HOST      | SMTP mail host                                      |

Now you can use it! 
Just send a message to the bot with the document you want to convert, and it will send it to your kindle.


### Links
You can find dev.to post [here](https://dev.to/michaelfmnk/developing-send-to-kindle-telegram-bot-120c) \
Message me if you have any questions or suggestions: [michael@fomenko.dev](mailto:michael@fomenko.dev)