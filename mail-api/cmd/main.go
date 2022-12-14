package main

import (
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"github.com/gomodule/redigo/redis"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/api"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/config"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/email"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/pkg/logger"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/storage/postgres"
	r "github.com/uzbekman2005/mailganer-test-task/mail-api/storage/redis"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel, "mailganer-test-task")
	emailSender := email.NewEmailSender(cfg, log)

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Error("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.RedisHost+":"+cfg.RedisPort)
		},
	}

	postgres, err := postgres.NewPostgres(*cfg)
	if err != nil {
		log.Error("Error while connecting to postgres")
		panic(err)
	}
	defer postgres.Db.Close()
	server := api.NewRouter(
		api.Option{
			Conf:           cfg,
			Logger:         log,
			EmailSender:    emailSender,
			CasbinEnforcer: casbinEnforcer,
			Redis:          r.NewRedisRepo(pool),
			Postgres:       postgres,
		},
	)

	if err := server.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatal("failed to run the server", logger.Error(err))
		panic(err)
	}
}
