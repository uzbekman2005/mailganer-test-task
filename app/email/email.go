package email

import (
	"bytes"
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

func NewEmailSender(cfg *config.Config, log logger.Logger) *EmailSender {
	return &EmailSender{
		Cfg:    cfg,
		Logger: log,
	}
}

func (e *EmailSender) SendEmailWithSupscibers(ecfg *models.SendEmailConfig, req *models.SendNewsToSupscribersReq) error {
	var body bytes.Buffer

	t, err := template.ParseFiles("./kafka/consumer/Email/email_temp/email.html")
	if err != nil {
		e.Logger.Error("Error while parsing HTML template", logger.Error(err))
		return err
	}

	for _, el := range req.To {
		t.Execute(&body, struct{ FirstName string }{FirstName: el.FirstName})
		t.Execute(&body, struct{ LastName string }{LastName: el.LastName})

		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		msg := []byte("SEND HTML TEMPLATE" + mime + body.String())

		auth := smtp.PlainAuth("", ecfg.Email, ecfg.Passwrod, "smtp.gmail.com")

		err = smtp.SendMail("smtp.gmail.com:587", auth, ecfg.Email, []string{el.Email}, msg)
		if err != nil {
			e.Logger.Error("Error while sending mail", logger.Error(err))
			return err
		}
	}
	return nil
}
