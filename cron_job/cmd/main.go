package main

import (
	"database/sql"
	"time"

	"github.com/uzbekman2005/mailganer-test-task/cron_job/config"
	"github.com/uzbekman2005/mailganer-test-task/cron_job/email"
	"github.com/uzbekman2005/mailganer-test-task/cron_job/pkg/logger"
	"github.com/uzbekman2005/mailganer-test-task/cron_job/pkg/models"
	p "github.com/uzbekman2005/mailganer-test-task/cron_job/storage/postgres"
)

type CronJob struct {
	Cfg         *config.Config
	Postgres    *p.Postgres
	EmailSender *email.EmailSender
	Logger      logger.Logger
}

func main() {
	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel, "mailganer-cronJob-task")
	emailSender := email.NewEmailSender(cfg, log)
	db, err := p.NewPostgres(*cfg)
	if err != nil {
		panic(err)
	}
	defer db.Db.Close()

	cronJob := &CronJob{
		Logger:      log,
		Cfg:         cfg,
		EmailSender: emailSender,
		Postgres:    db,
	}
	cronJob.Logger.Info("Cron job is working: ")
	for {
		scheduledMessages, err := cronJob.Postgres.GetScheduledMessages()
		if err != sql.ErrNoRows && err != nil {
			cronJob.Logger.Error("Error while getting from database", logger.Error(err))
		}

		for _, msg := range scheduledMessages {
			err = cronJob.EmailSender.SendEmailToSupscibers(
				&models.SendEmailConfig{
					Email:    msg.SenderEmail,
					Passwrod: msg.EmailPaassword,
				},
				&models.SendNewsToSupscribersReq{
					To: []*models.Subscriber{
						msg.To,
					},
					News: msg.News,
				})
			if err != nil {
				cronJob.Logger.Error("Error sending emails", logger.Error(err))
			}
			cronJob.Logger.Info("Messages is sent to " + msg.To.FirstName)
		}

		time.Sleep(time.Minute * 1)
	}
}
