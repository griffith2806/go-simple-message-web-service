package routes

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/gg-messages/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreatedResponse struct {
	Id int64 `json:"id"`
}

// @BasePath /messages

// GetMessages godoc
// @Summary Retrieve all messages
// @Description Fetches a list of all messages.
// @Tags messages
// @Produce json
// @Success 200 {array} string "List of messages"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /messages [get]
func getMessages(context *gin.Context) {
	messages, err := models.GetAllMessages()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch messages"})
		return
	}

	context.JSON(http.StatusOK, messages)
}

// GetMessage godoc
// @Summary Retrieve a message by ID
// @Description Fetches a single message by its ID.
// @Tags messages
// @Produce plain
// @Param id path int64 true "Message ID"
// @Success 200 {string} string "Message text"
// @Failure 400 {object} ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} ErrorResponse "Message not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /messages/{id} [get]
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

// CreateMessage godoc
// @Summary Create a new message
// @Description Creates a new message with the provided content.
// @Tags messages
// @Accept plain
// @Produce json
// @Param message body string true "Message content"
// @Success 201 {object} CreatedResponse "Created message ID"
// @Failure 400 {object} ErrorResponse "Bad request: Empty message"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /messages [post]
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
