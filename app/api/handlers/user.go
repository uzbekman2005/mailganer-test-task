package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"github.com/uzbekman2005/mailganer-test-task/app/api/models"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/etc"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/logger"
)

// @Summary 		SendEmailWithSupscribers
// @Description 	Through this api, news and informatin of all subscirbers will be sent
// @Tags 			User
// @Accept 			json
// @Produce         json
// @Param           User      body  	  models.User true "User"
// @Success         200					  {object} 	models.SaveUserInRedis
// @Failure         500                   {object}  models.StatusInfo
// @Failure         409                   {object}  models.StatusInfo
// @Router          /user/register [post]
func (h *Handler) Register(c *gin.Context) {
	body := &models.User{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.StatusInfo{
			Message: "Please check your data",
		})
		h.Log.Error("Error while registering user", logger.Error(err))
		return
	}

	isUniqueEmail, err := h.Redis.Exists(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Please check your data",
		})
		h.Log.Error("Error while checking email uniqeness", logger.Error(err))
		return
	}

	if cast.ToString(isUniqueEmail) == "1" {
		c.JSON(http.StatusConflict, &models.StatusInfo{
			Message: "This email is already registered",
		})
		return
	}

	body.Password, err = etc.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Please check your data",
		})
		h.Log.Error("Error while hashing password", logger.Error(err))
		return
	}

	bodyWithId := models.SaveUserInRedis{
		Id:            uuid.New().String(),
		Name:          body.Name,
		Password:      body.Password,
		Email:         body.Email,
		EmailPassword: body.EmailPassword,
	}

	saveInRedisByte, err := json.Marshal(bodyWithId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Something went wrong",
		})
		h.Log.Error("Error while marshaling", logger.Error(err))
		return
	}

	err = h.Redis.Set(bodyWithId.Email, string(saveInRedisByte))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Something went wrong",
		})
		h.Log.Error("Error while saving in redis", logger.Error(err))
		return
	}

	h.JWTHandler.Sub = bodyWithId.Email
	h.JWTHandler.Role = "authorized"
	h.JWTHandler.Aud = []string{"mailganer"}
	h.JWTHandler.SigninKey = h.Cfg.SignInKey
	h.JWTHandler.Log = h.Log
	tokens, err := h.JWTHandler.GenerateAuthJWT()
	accessToken := tokens[0]
	bodyWithId.AccessToken = accessToken
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusInfo{
			Message: "Something went wrong",
		})
		h.Log.Error("Error while generating token", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, bodyWithId)
}

// @Summary 		Login users
// @Description 	Through this api, User can login
// @Tags 			User
// @Accept 			json
// @Produce         json
// @Param        	email     query  		  string    true 	"email"
// @Param        	password  query  		  string    true 	"password"
// @Success         200					  {object} 	models.SaveUserInRedis
// @Failure         500                   {object}  models.StatusInfo
// @Failure         409                   {object}  models.StatusInfo
// @Router          /user/login [get]
func (h *Handler) Login(c *gin.Context) {
	var (
		email    = c.Query("email")
		password = c.Query("password")
		body     = &models.SaveUserInRedis{}
	)

	resRedis, err := h.Redis.Get(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.StatusInfo{
			Message: "You haven't regitered before, please Sign up first",
		})
		h.Log.Error("Error while getting from redis", logger.Error(err))
		return
	}

	resRedisString := cast.ToString(resRedis)

	err = json.Unmarshal([]byte(resRedisString), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Oops something went wrong",
		})
		h.Log.Error("Error while unmarshaling from json", logger.Error(err))
		return
	}

	if !etc.CheckPasswordHash(password, body.Password) {
		c.JSON(http.StatusConflict, &models.StatusInfo{
			Message: "Password is incorrect",
		})
		return
	}

	h.JWTHandler.Sub = body.Email
	h.JWTHandler.Role = "authorized"
	h.JWTHandler.Aud = []string{"mailganer"}
	h.JWTHandler.SigninKey = h.Cfg.SignInKey
	h.JWTHandler.Log = h.Log
	tokens, err := h.JWTHandler.GenerateAuthJWT()
	accessToken := tokens[0]
	body.AccessToken = accessToken
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusInfo{
			Message: "Something went wrong",
		})
		h.Log.Error("Error while generating token", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, body)
}

// @Summary 		Login users
// @Description 	Through this api, User can login
// @Tags 			User
// @Security        BearerAuth
// @Accept 			json
// @Produce         json
// @Success         200					  {object} 	models.SaveUserInRedis
// @Failure         500                   {object}  models.StatusInfo
// @Failure         409                   {object}  models.StatusInfo
// @Router          /user/profile [get]
func (h *Handler) Profile(c *gin.Context) {
	claims, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusInfo{
			Message: "Ooops something went wrong",
		})
		h.Log.Error("Error while getting claims of user", logger.Error(err))
		return
	}

	body := &models.SaveUserInRedis{}
	resRedis, err := h.Redis.Get(claims.Sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.StatusInfo{
			Message: "You haven't regitered before, please Sign up first",
		})
		h.Log.Error("Error while getting from redis", logger.Error(err))
		return
	}

	resRedisString := cast.ToString(resRedis)

	err = json.Unmarshal([]byte(resRedisString), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Oops something went wrong",
		})
		h.Log.Error("Error while unmarshaling from json", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, body)
}
