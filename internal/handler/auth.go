package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smth/internal/model"
)

const (
	defaultRedirectURL = "/registration"
	maxAgeCookie       = 300
)

func (h *Handler) singUpPage() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "singup.html", gin.H{})
	}

}

func (h *Handler) singInPage() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.HTML(200, "singin.html", gin.H{})
	}

}

func (h *Handler) handlerRegisterUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		password := c.PostForm("password")
		rePassword := c.PostForm("rePassword")
		if password != rePassword {
			c.Redirect(http.StatusMovedPermanently, defaultRedirectURL)
			c.JSON(http.StatusOK, gin.H{
				"Error": "Passwords do not match",
			})
			return
		}
		role, err := h.store.User().GetRole()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		u := model.User{
			Login:    c.PostForm("login"),
			Email:    c.PostForm("email"),
			Password: password,
			Role:     role,
		}

		if err := u.Validate(); err != true {
			c.Redirect(http.StatusMovedPermanently, defaultRedirectURL)
			c.JSON(http.StatusOK, gin.H{
				"Error":   "Validation error",
				"Message": "Invalid email or password",
			})
			return
		}

		if err := h.store.User().CreateUser(&u); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		u.Sanitize()
		token, err := h.tokenManager.GenerateJWT(u.ID, u.Role)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.SetCookie("auth-token", token, maxAgeCookie, "/", "localhost", false, true)

		if _, err = h.tokenManager.RefreshJWT(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/")

	}
}

func (h *Handler) handlerLoginUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		u, err := h.store.User().FindByLogin(c.PostForm("login"))

		if err != nil || u.CheckUserPassword(c.PostForm("password")) != nil {
			c.Redirect(http.StatusMovedPermanently, defaultRedirectURL)
			c.JSON(http.StatusOK, gin.H{
				"Error":   "Validation error",
				"Message": "Invalid password",
			})
			return
		}

		token, err := h.tokenManager.GenerateJWT(u.ID, u.Role)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.SetCookie("auth-token", "Bearer "+token, maxAgeCookie, "/", "localhost", false, true)

		c.Redirect(http.StatusMovedPermanently, "/")

	}

}
