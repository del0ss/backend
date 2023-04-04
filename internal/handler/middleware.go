package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	_, err := h.parseAuthCookie(c)
	if err != nil {
		return
	}
}

func (h *Handler) parseAuthCookie(c *gin.Context) (interface{}, error) {
	cookie, err := c.Cookie("auth-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return nil, err
	}
	if cookie == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(cookie, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}
	return h.tokenManager.Parse(headerParts[1])
}
