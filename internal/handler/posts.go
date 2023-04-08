package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		fmt.Println(p)
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

func (h *Handler) HandlerCreatePost() gin.HandlerFunc {

	return func(c *gin.Context) {
		var p model.Post
		if err := c.BindJSON(&p); err != nil {
			newErrorMessage(c, http.StatusUnauthorized, err.Error())
			return
		}

		userId, err := h.parseAuthHeader(c)
		if err != nil {
			newErrorMessage(c, http.StatusUnauthorized, err.Error())
			return
		}

		if err := h.store.Post().CreatePost(&p, userId); err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}
