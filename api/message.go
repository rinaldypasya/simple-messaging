package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rinaldypasya/simple-messaging/message"
)

func (idb *InDB) Send(c *gin.Context) {
	var (
		msg    message.Message
		result gin.H
	)
	text := c.PostForm("text")
	msg.Text = text
	idb.DB.Create(&msg)
	result = gin.H{
		"result": msg,
		"status": "success",
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetMessages(c *gin.Context) {
	var (
		messages []message.Message
		result   gin.H
	)

	idb.DB.Find(messages)
	if len(messages) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": messages,
			"count":  len(messages),
		}
	}

	c.JSON(http.StatusOK, result)
}
