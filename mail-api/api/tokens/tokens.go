package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/uzbekman2005/mailganer-test-task/mail-api/pkg/logger"
)

// JWTHandler ...
type JWTHandler struct {
	Sub       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	SigninKey string
	Log       logger.Logger
	Token     string
}

type CustomClaims struct {
	*jwt.Token
	Sub  string   `json:"sub"`
	Exp  float64  `json:"exp"`
	Iat  float64  `json:"iat"`
	Aud  []string `json:"aud"`
	Role string   `json:"role"`
}

// GenerateAuthJWT ...
func (jwtHandler *JWTHandler) GenerateAuthJWT() ([]string, error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = jwtHandler.Sub
	claims["exp"] = time.Now().Add(time.Hour * 500).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role
	claims["aud"] = jwtHandler.Aud
	access, err := accessToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		jwtHandler.Log.Error("error generating access token", logger.Error(err))
		return []string{access, ""}, err
	}

	refresh, err := refreshToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		jwtHandler.Log.Error("error generating refresh token", logger.Error(err))
		return []string{access, refresh}, err
	}
	return []string{access, refresh}, nil
}

// ExtractClaims ...
func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.SigninKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		jwtHandler.Log.Error("invalid jwt token")
		return nil, err
	}
	return claims, nil
}

// ExtractClaim extracts claims from given token
func ExtractClaim(tokenStr string, signinigKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	fmt.Println(tokenStr)
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signinigKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil
}
