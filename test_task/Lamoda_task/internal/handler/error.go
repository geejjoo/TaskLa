package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusСode int, message string) {
	log.Errorf(c, message)
	c.AbortWithStatusJSON(statusСode, errorResponse{message})
}
