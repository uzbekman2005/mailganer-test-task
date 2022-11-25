package email

import (
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

func (e *EmailSender) SendEmailWithSupscibers(req *models.SendEmailWithSupscribersReq) error {
	return nil
}
