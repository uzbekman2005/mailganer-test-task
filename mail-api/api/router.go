package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/uzbekman2005/mailganer-test-task/mail-api/api/docs" //swag
	"github.com/uzbekman2005/mailganer-test-task/mail-api/api/handlers"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/api/middleware"
	token "github.com/uzbekman2005/mailganer-test-task/mail-api/api/tokens"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/config"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/email"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/pkg/logger"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/storage/postgres"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/storage/repo"
)

type Option struct {
	Conf           *config.Config
	Logger         logger.Logger
	EmailSender    *email.EmailSender
	CasbinEnforcer *casbin.Enforcer
	Redis          repo.InMemoryStorageI
	Postgres       *postgres.Postgres
}

// New ...
// @title           Mailganer-test-task
// @version         1.0
// @description     This test task server

// @contact.name   Azizbek
// @contact.url    https://t.me/azizbek_dev_2005
// @contact.email  azizbekhojimurodov@gmail.com

// @host      localhost:9090

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(opt Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	jwtHandler := token.JWTHandler{
		SigninKey: opt.Conf.SignInKey,
		Log:       opt.Logger,
	}

	router.Use(middleware.NewAuth(opt.CasbinEnforcer, jwtHandler, *opt.Conf))
	handler := handlers.Handler{
		Log:         opt.Logger,
		Cfg:         opt.Conf,
		EmailSender: opt.EmailSender,
		Redis:       opt.Redis,
		Postgres:    opt.Postgres,
	}

	router.POST("/email/tosubscribers", handler.SendNewsToSupscribers)
	router.POST("/user/register", handler.Register)
	router.GET("/user/login", handler.Login)
	router.GET("/user/profile", handler.Profile)
	router.POST("/email/schedule", handler.SendScheduledEmails)
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
