package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/uzbekman2005/mailganer-test-task/app/api/models"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/etc"
	"github.com/uzbekman2005/mailganer-test-task/app/pkg/logger"
)

// @Summary 		SendNewsToSupscribers
// @Description 	Through this api, news and informatin of all subscirbers will be sent
// @Tags 			Email
// @Security        BearerAuth
// @Accept 			json
// @Produce         json
// @Param           SendNewsToSupscribersReq      body  	  models.SendNewsToSupscribersReq true "SendNewsToSupscribersReq"
// @Success         200					  {object} 	models.StatusInfo
// @Failure         500                   {object}  models.StatusInfo
// @Failure         409                   {object}  models.StatusInfo
// @Router          /email/tosubscribers [post]
func (h *Handler) SendNewsToSupscribers(c *gin.Context) {
	claims, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusInfo{
			Message: "Ooops something went wrong",
		})
		h.Log.Error("Error while getting claims of user", logger.Error(err))
		return
	}

	user := &models.SaveUserInRedis{}
	resRedis, err := h.Redis.Get(claims.Sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.StatusInfo{
			Message: "You haven't regitered before, please Sign up first",
		})
		h.Log.Error("Error while getting from redis", logger.Error(err))
		return
	}

	resRedisString := cast.ToString(resRedis)

	err = json.Unmarshal([]byte(resRedisString), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Oops something went wrong",
		})
		h.Log.Error("Error while unmarshaling from json", logger.Error(err))
		return
	}

	body := &models.SendNewsToSupscribersReq{}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusConflict, models.StatusInfo{
			Message: "Please check your data",
		})
		h.Log.Error("Error while binding from request", logger.Error(err))
		return
	}

	for _, e := range body.To {
		if _, ok := etc.ValidMailAddress(e.Email); !ok {
			c.JSON(http.StatusBadRequest, models.StatusInfo{
				Message: "Invalid mailing address" + e.Email,
			})
			return
		}
	}

	err = h.EmailSender.SendEmailWithSupscibers(&models.SendEmailConfig{
		Email:    user.Email,
		Passwrod: user.EmailPassword,
	}, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusInfo{
			Message: "Something went wrong",
		})
		h.Log.Error("Error while sending email with subsribers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.StatusInfo{
		Message: "All Emails send successfully",
	})
}

// @Summary 		SendScheduledEmails
// @Description 	Through this api, news and informatin of all subscirbers will be sent
// @Tags 			Email
// @Security        BearerAuth
// @Accept 			json
// @Produce         json
// @Param           SendScheduledEmailsApiReq      body  	  models.SendScheduledEmailsApiReq true "SendScheduledEmailsApiReq"
// @Success         200					  {object} 	models.StatusInfo
// @Failure         500                   {object}  models.StatusInfo
// @Failure         409                   {object}  models.StatusInfo
// @Router          /email/schedule [post]
func (h *Handler) SendScheduledEmails(c *gin.Context) {
	claims, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.StatusInfo{
			Message: "Ooops something went wrong",
		})
		h.Log.Error("Error while getting claims of user", logger.Error(err))
		return
	}

	user := &models.SaveUserInRedis{}
	resRedis, err := h.Redis.Get(claims.Sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.StatusInfo{
			Message: "You haven't regitered before, please Sign up first",
		})
		h.Log.Error("Error while getting from redis", logger.Error(err))
		return
	}

	resRedisString := cast.ToString(resRedis)

	err = json.Unmarshal([]byte(resRedisString), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Oops something went wrong",
		})
		h.Log.Error("Error while unmarshaling from json", logger.Error(err))
		return
	}

	body := &models.SendScheduledEmailsApiReq{}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusConflict, models.StatusInfo{
			Message: "Please check your data",
		})
		h.Log.Error("Error while binding from request", logger.Error(err))
		return
	}

	for _, e := range body.To {
		if email, ok := etc.ValidMailAddress(e.Email); !ok {
			c.JSON(http.StatusConflict, models.StatusInfo{
				Message: "Invalid mail address " + email,
			})
			return
		}
	}

	message := &models.SendScheduledEmailsReq{
		News:           body.News,
		MinutsAfter:    body.MinutsAfter,
		SenderEmail:    user.Email,
		EmailPaassword: user.EmailPassword,
		To:             body.To,
	}

	err = h.Postgres.WriteMessagesToDb(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.StatusInfo{
			Message: "Oops something went wrong",
		})
		h.Log.Error("Error while writing to database", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &models.StatusInfo{
		Message: "Sending emailes Successfulley scheduled",
	})
}
