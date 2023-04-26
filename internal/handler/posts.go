package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"smth/internal/model"
	"strconv"
)

func (h *Handler) GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get(userContext)
		if ok == false {
			newErrorMessage(c, http.StatusUnauthorized, "invalid header")
			return
		}
		p, err := h.store.Post().GetPosts()
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) CreatePosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func (h *Handler) GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := h.store.Post().GetPost(id)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := h.store.Post().DeletePost(id)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "OK")
	}
}

func (h *Handler) HandlerCreatePost() gin.HandlerFunc {

	return func(c *gin.Context) {
		var p model.Post
		if err := c.BindJSON(&p); err != nil {
			newErrorMessage(c, http.StatusUnauthorized, err.Error())
			return
		}

		//userId, err := h.parseAuthHeader(c)
		//if err != nil {
		//	newErrorMessage(c, http.StatusUnauthorized, err.Error())
		//	return
		//}

		id, err := h.store.Post().CreatePost(p, 1)
		if err != nil {
			logrus.Error(err)
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, id)
	}
}
