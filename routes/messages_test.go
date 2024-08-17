package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/gg-messages/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMessages(t *testing.T) {
	db.InitDb()
	server := gin.Default()

	RegisterRoutes(server)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateMessagesSuccess(t *testing.T) {
	db.InitDb()
	server := gin.Default()

	RegisterRoutes(server)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/messages", strings.NewReader("example msg"))
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateMessages_NoInput_ReturnsBadRequest(t *testing.T) {
	db.InitDb()
	server := gin.Default()

	RegisterRoutes(server)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/messages", strings.NewReader(""))
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetMessage_IdNotFound_ReturnsNotFound(t *testing.T) {
	db.InitDb()
	server := gin.Default()

	RegisterRoutes(server)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages/1234", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetMessage_IdFound_ReturnsSuccess(t *testing.T) {
	db.InitDb()
	server := gin.Default()

	RegisterRoutes(server)

	w := httptest.NewRecorder()
	http.NewRequest("POST", "/messages", strings.NewReader("example msg"))
	req, _ := http.NewRequest("GET", "/messages/1", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
