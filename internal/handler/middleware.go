package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	userID, err := h.parseAuthHeader(c)
	if err != nil {
		logrus.Error(err)
		return
	}

	c.Set(userContext, userID)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (interface{}, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}
