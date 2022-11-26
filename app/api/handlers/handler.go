package handlers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	t "github.com/uzbekman2005/mailganer-test-task/app/api/tokens"

	"github.com/uzbekman2005/mailganer-test-task/app/config"
	"github.com/uzbekman2005/mailganer-test-task/app/email"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/logger"
	"github.com/uzbekman2005/mailganer-test-task/app/storage/postgres"
	"github.com/uzbekman2005/mailganer-test-task/app/storage/repo"
)

type Handler struct {
	Log         logger.Logger
	Cfg         *config.Config
	EmailSender *email.EmailSender
	Redis       repo.InMemoryStorageI
	JWTHandler  t.JWTHandler
	Postgres    *postgres.Postgres
}

func GetClaims(h Handler, c *gin.Context) (*t.CustomClaims, error) {
	var (
		claims = t.CustomClaims{}
	)
	strToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) { return []byte(h.Cfg.SignInKey), nil })

	if err != nil {
		h.Log.Error("invalid access token")
		return nil, err
	}
	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Sub = rawClaims["sub"].(string)
	claims.Exp = rawClaims["exp"].(float64)
	aud := cast.ToStringSlice(rawClaims["aud"])
	claims.Aud = aud
	claims.Role = rawClaims["role"].(string)
	claims.Sub = rawClaims["sub"].(string)
	claims.Token = token
	return &claims, nil
}
