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
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (h *Handler) CreatePosts() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_post.html", gin.H{})
	}
}

func (h *Handler) GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		fmt.Println(id)
		p, err := h.store.Post().GetPost(id)
		fmt.Println(err)
		if err != nil {
			c.HTML(http.StatusOK, "error.html", gin.H{})
		}
		c.HTML(http.StatusOK, "post.html", gin.H{
			"Title":          "Post",
			"isAuthenticate": true,
			"Post":           p,
		})

	}
}

func (h *Handler) HandlerCreatePost() gin.HandlerFunc {

	return func(c *gin.Context) {
		post := model.Post{Title: c.PostForm("title"), Content: c.PostForm("content")}
		userId, _ := h.parseAuthHeader(c)
		err := h.store.Post().CreatePost(&post, userId)

		if err != nil {
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/posts")
	}
}
