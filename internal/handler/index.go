package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) helloPage() gin.HandlerFunc {

	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		p, _ := h.store.Post().GetPosts()
		if header == "" {

		}
		c.JSON(http.StatusOK, p)
	}

}

func (h *Handler) handlePe() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{})
	}

}
