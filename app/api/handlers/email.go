package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/uzbekman2005/mailganer-test-task/app/api/models"
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
	body := &models.SendNewsToSupscribersReq{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusConflict, models.StatusInfo{
			Message: "Please check your data",
		})
		h.Log.Error("Error while binding from request", logger.Error(err))
		return
	}

	err = h.EmailSender.SendEmailWithSupscibers(body)
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