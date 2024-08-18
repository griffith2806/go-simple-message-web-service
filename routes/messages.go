package routes

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/gg-messages/models"
	"github.com/gin-gonic/gin"
)

func getMessages(context *gin.Context) {
	messages, err := models.GetAllMessages()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch messages"})
		return
	}

	context.JSON(http.StatusOK, messages)
}

func getMessage(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch message"})
		return
	}

	messageText, err := models.GetMessage(id)
	if err != nil {
		if err == models.ErrMessageNotFound {
			context.JSON(http.StatusNotFound, gin.H{"message": "Message not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch message"})
		}
		return
	}

	context.String(http.StatusOK, messageText)
}

func createMessage(context *gin.Context) {
	var message models.Message
	bytes, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse content"})
		return
	}
	if len(bytes) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "You must enter a message"})
		return
	}

	messageText := string(bytes)
	message.Message = strings.TrimSpace(messageText)
	message.CreatedAt = time.Now()
	err = message.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create message"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": message.Id})
}
