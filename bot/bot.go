package bot

import (
	"errors"
	"fmt"
	"github.com/scorredoira/email"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	defaultSmtpPort = "587"
	tmpFilesPath    = "/files/"
)

var (
	ErrTokenNotSet     = errors.New("token for telegram bot not set")
	ErrPasswordNotSet  = errors.New("password for email not set")
	ErrEmailFromNotSet = errors.New("emailfrom not set")
	ErrEmailToNotSet   = errors.New("emailto not set")

	ErrConversion = errors.New("could not convert file")
	ErrStartup    = errors.New("could not create telebot instance")

	supportedFormats = []string{"doc", "docx", "rtf", "htm", "html", "txt", "mobi", "pdf"}
)

// SendToKindleBot stores bot configuration
type SendToKindleBot struct {
	Token     string
	EmailFrom string
	EmailTo   string
	SmtpHost  string
	SmtpPort  string
	Password  string
}

// Start starts bot. It is blocking.
// If there is an error during startup, returns it. Otherwise blocks
func (b *SendToKindleBot) Start() error {
	if err := b.verifyConfig(); err != nil {
		return err
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  b.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return ErrStartup
	}

	bot.Handle(tb.OnDocument, b.documentHandler(bot))
	bot.Start()

	return nil
}

func (b *SendToKindleBot) documentHandler(bot *tb.Bot) func(msg *tb.Message) {
	return func(msg *tb.Message) {
		doc := msg.Document
		nameParts := strings.Split(doc.FileName, ".")
		fileNameWithoutExtension := strings.Join(nameParts[:len(nameParts)-1], "")
		extension := nameParts[len(nameParts)-1]

		originalFilePath := tmpFilesPath + doc.FileName
		if err := bot.Download(&doc.File, originalFilePath); err != nil {
			log.Println("could not download file", err)
			respond(bot, msg, "Sorry. I could not download file")
		}
		defer removeSilently(originalFilePath)

		fileToSend := originalFilePath
		if needToConvert(extension) {
			outputFilePath := tmpFilesPath + fileNameWithoutExtension + ".mobi"
			if err := convert(originalFilePath, outputFilePath); err != nil {
				log.Println("could not convert file", err)
				respond(bot, msg, "Sorry. I could not convert file")
			}
			fileToSend = outputFilePath
			defer removeSilently(outputFilePath)
		}

		if err := b.sendFileViaEmail(fileToSend); err != nil {
			log.Println("could not send file", err)
			respond(bot, msg, "Sorry. I could not send file")
		}
	}
}

func needToConvert(extension string) bool {
	for _, format := range supportedFormats {
		if format == extension {
			return false
		}
	}
	return true
}

func respond(bot *tb.Bot, m *tb.Message, text string) {
	if _, err := bot.Send(m.Sender, text); err != nil {
		log.Println(fmt.Sprintf("could not send a message to %d", m.Sender.ID), err)
	}
}

func convert(in, out string) error {
	cmd := exec.Command("ebook-convert", in, out)
	if err := cmd.Run(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	if _, err := os.Stat(out); errors.Is(err, os.ErrNotExist) {
		return ErrConversion
	}
	return nil
}

func removeSilently(path string) {
	if err := os.Remove(path); err != nil {
		log.Println(fmt.Sprintf("could not delete file %s", path), err)
	}
}

func (b *SendToKindleBot) verifyConfig() error {
	if b.Token == "" {
		return ErrTokenNotSet
	}
	if b.Password == "" {
		return ErrPasswordNotSet
	}
	if b.EmailFrom == "" {
		return ErrEmailFromNotSet
	}
	if b.EmailTo == "" {
		return ErrEmailToNotSet
	}
	if b.SmtpPort == "" {
		b.SmtpPort = defaultSmtpPort
	}
	return nil
}

func (b *SendToKindleBot) sendFileViaEmail(path string) error {
	msg := email.NewMessage("", "")
	msg.From = mail.Address{Name: "From", Address: b.EmailFrom}
	msg.To = []string{b.EmailTo}

	if err := msg.Attach(path); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", b.EmailFrom, b.Password, b.SmtpHost)
	addr := fmt.Sprintf("%s:%s", b.SmtpHost, b.SmtpPort)
	if err := email.Send(addr, auth, msg); err != nil {
		return err
	}
	return nil
}
