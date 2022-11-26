package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/uzbekman2005/mailganer-test-task/app/api/models"
	"github.com/uzbekman2005/mailganer-test-task/app/config"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/logger"
)

type EmailSender struct {
	Cfg    *config.Config
	Logger logger.Logger
}

type Message struct {
	FirstName string
	LastName  string
	News      string
}

func NewEmailSender(cfg *config.Config, log logger.Logger) *EmailSender {
	return &EmailSender{
		Cfg:    cfg,
		Logger: log,
	}
}

func (e *EmailSender) SendEmailWithSupscibers(ecfg *models.SendEmailConfig, req *models.SendNewsToSupscribersReq) error {
	for _, el := range req.To {
		body := new(bytes.Buffer)
		t, err := template.ParseFiles("/home/azizbek/go/src/github.com/uzbekman2005/mailganer-test-task/app/email/html_templates/news.html")
		if err != nil {
			e.Logger.Error("Error while parsing HTML template", logger.Error(err))
			return err
		}
		mInfo := &Message{
			FirstName: el.FirstName,
			LastName:  el.LastName,
			News:      req.News,
		}

		t.Execute(body, mInfo)
		fmt.Println(body.String())
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		msg := []byte("Subject: News" + mime + body.String())

		auth := smtp.PlainAuth("", ecfg.Email, ecfg.Passwrod, "smtp.gmail.com")

		err = smtp.SendMail("smtp.gmail.com:587", auth, ecfg.Email, []string{el.Email}, msg)
		if err != nil {
			e.Logger.Error("Error while sending mail", logger.Error(err))
			return err
		}
	}
	return nil
}
