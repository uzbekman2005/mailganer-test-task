package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/uzbekman2005/mailganer-test-task/app/api/models"
	"github.com/uzbekman2005/mailganer-test-task/app/config"
)

type Postgres struct {
	Conn *sqlx.DB
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
	return &Postgres{Conn: conn}, nil
}

func (p *Postgres) WriteMessagesToDb(req *models.SendNewsToSupscribersReq) {
	
}
