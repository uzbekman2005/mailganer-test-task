package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/uzbekman2005/mailganer-test-task/app/api/models"
	"github.com/uzbekman2005/mailganer-test-task/app/config"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sqlx.DB
}

func NewPostgres(cfg config.Config) (*Postgres, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	conn, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return &Postgres{}, err
	}
	return &Postgres{Db: conn}, nil
}

func (p *Postgres) WriteMessagesToDb(req *models.SendScheduledEmailsReq) error {
	newsId := uuid.New().String()
	err := p.Db.QueryRow(`INSERT INTO news (
		id,
		content, 
		sender_email, 
		sender_email_password
	) values ($1, $2, $3, $4)`,
		newsId,
		req.News,
		req.SenderEmail,
		req.EmailPaassword,
	).Err()

	if err != nil {
		return err
	}

	for _, e := range req.To {
		err = p.Db.QueryRow(`INSERT INTO messages (
			id,
			news_id,
			first_name, 
			last_name,
			email,
			minutes_after 
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		`,
			uuid.New().String(),
			newsId,
			e.FirstName,
			e.LastName,
			e.Email,
			req.MinutsAfter,
		).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetScheduledMessages() ([]*models.GetScheduledMessagesRes, error) {
	response := []*models.GetScheduledMessagesRes{}
	rows, err := p.Db.Query(
		`SELECT 
			n.content, 
			n.sender_email,
			n.sender_email_password,
			m.id,
			m.first_name,
			m.last_name, 
			m.email
		FROM 
			messages as m INNER JOIN news as n ON n.id=m.news_id
		WHERE 
			EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - scheduled_at)) >= (m.minutes_after * 60)
		`,
	)
	if err != nil {
		return response, err
	}
	defer rows.Close()
	ids := []string{}
	for rows.Next() {
		msg := &models.GetScheduledMessagesRes{}
		msg.To = &models.Subscriber{}
		id := ""
		err = rows.Scan(
			&msg.News,
			&msg.SenderEmail,
			&msg.EmailPaassword,
			&id,
			&msg.To.FirstName,
			&msg.To.LastName,
			&msg.To.Email,
		)
		if err != nil {
			return response, err
		}

		ids = append(ids, id)
		response = append(response, msg)
	}

	err = p.DeleteMessagesByIds(ids)
	if err != nil {
		return []*models.GetScheduledMessagesRes{}, err
	}
	return response, nil
}

func (p *Postgres) DeleteMessagesByIds(ids []string) error {
	for _, id := range ids {
		_, err := p.Db.Exec(`DELETE FROM messages where id=$1`, id)
		if err != nil {
			return err
		}
	}
	return nil
}
