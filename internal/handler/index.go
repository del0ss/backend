package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) helloPage() gin.HandlerFunc {
	isAuth := true

	return func(c *gin.Context) {
		_, err := c.Cookie("auth-token")
		if err != nil {
			isAuth = false
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":          "Home",
			"isAuthenticate": isAuth,
		})

	}

}

func (h *Handler) handlePe() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "zxc.html", gin.H{})
	}

}
