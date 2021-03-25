package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func RespondJSON(w *gin.Context, status_code int, data interface{}) {
	log.Println("status code: ", status_code)
	var msg Message

	msg.StatusCode = status_code
	msg.Data = data
	w.JSON(200, msg)
}
