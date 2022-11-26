package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/uzbekman2005/mailganer-test-task/app/api/docs" //swag
	"github.com/uzbekman2005/mailganer-test-task/app/api/handlers"
	"github.com/uzbekman2005/mailganer-test-task/app/api/middleware"
	token "github.com/uzbekman2005/mailganer-test-task/app/api/tokens"
	"github.com/uzbekman2005/mailganer-test-task/app/config"
	"github.com/uzbekman2005/mailganer-test-task/app/email"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/logger"
	"github.com/uzbekman2005/mailganer-test-task/app/storage/postgres"
	"github.com/uzbekman2005/mailganer-test-task/app/storage/repo"
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
	}

	router.POST("/email/tosubscribers", handler.SendNewsToSupscribers)
	router.POST("/user/register", handler.Register)
	router.GET("/user/login", handler.Login)
	router.GET("/user/profile", handler.Profile)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
