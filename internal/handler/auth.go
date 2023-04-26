package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"smth/internal/model"
)

const (
	defaultRedirectURL = "/registration"
	maxAgeCookie       = 300
)

func (h *Handler) registerUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var u = model.User{Role: 0}
		if err := c.BindJSON(&u); err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		if u.Password != u.RePassword {
			c.JSON(http.StatusOK, gin.H{
				"errorResponse": "Passwords do not match",
			})
			return
		}

		if err := u.Validate(); err != true {
			c.JSON(http.StatusOK, gin.H{
				"errorResponse": "Invalid email or password",
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

type signInInput struct {
	Login    string `json:"login" binding:"required,min=4,max=15"`
	Password string `json:"password,omitempty" binding:"required,min=5,max=25"`
}

func (h *Handler) loginUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var signInInput signInInput
		if err := c.BindJSON(&signInInput); err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		u, err := h.store.User().FindByLogin(signInInput.Login)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		if u.CheckUserPassword(signInInput.Password) != nil {
			logrus.Error(errors.New("password do not match"))
			c.JSON(http.StatusOK, gin.H{
				"errorResponse": "Password do not match",
			})
			return
		}

		token, err := h.tokenManager.GenerateJWT(u.ID, u.Role)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
		})
	}
}
