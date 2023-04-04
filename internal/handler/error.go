package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, tmpl, msg string) {
	logrus.Error(msg)
	c.JSON(401, gin.H{"Error": "Unauthorized"})
}
