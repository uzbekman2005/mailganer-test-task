package handlers

import (
	"github.com/uzbekman2005/mailganer-test-task/app/config"
	"github.com/uzbekman2005/mailganer-test-task/app/email"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/logger"
	"github.com/uzbekman2005/mailganer-test-task/app/storage/repo"
)

type Handler struct {
	Log         logger.Logger
	Cfg         *config.Config
	EmailSender *email.EmailSender
	Redis       repo.InMemoryStorageI
}
