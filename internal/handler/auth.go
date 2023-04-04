package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"smth/internal/model"
)

const (
	defaultRedirectURL = "/registration"
	maxAgeCookie       = 300
)

func (h *Handler) singUpPage() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}

}

func (h *Handler) singInPage() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}

}

func (h *Handler) handlerRegisterUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var u = model.User{Role: 0}
		if err := c.BindJSON(&u); err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		if u.Password != u.RePassword {
			c.JSON(http.StatusOK, gin.H{
				"Error":   "Validation password",
				"Message": "Passwords do not match",
			})
			return
		}

		if err := u.Validate(); err != true {
			c.JSON(http.StatusOK, gin.H{
				"Error":   "Validation error",
				"Message": "Invalid email or password",
			})
			return
		}

		if err := h.store.User().CreateUser(&u); err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		u.Sanitize()

		token, err := h.tokenManager.GenerateJWT(u.ID, u.Role)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
		})

	}
}

func (h *Handler) handlerLoginUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var u model.User
		if err := c.BindJSON(&u); err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}

		if u.CheckUserPassword(c.PostForm("password")) != nil {
			c.JSON(http.StatusOK, gin.H{
				"Error":   "Validation password",
				"Message": "Password do not match",
			})
			return
		}

		token, err := h.tokenManager.GenerateJWT(u.ID, u.Role)

		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
		})
	}
}
