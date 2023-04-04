package handler

import (
	"github.com/gin-gonic/gin"
	"smth/internal/store"
	"smth/internal/store/sqlstore"
	"smth/pkg/auth"
)

type Handler struct {
	store        store.Store
	tokenManager auth.TokenManager
}

func New(store *sqlstore.Store, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		store:        store,
		tokenManager: tokenManager,
	}
}

func (h *Handler) ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", h.helloPage())
	router.GET("/sing-up", h.singUpPage())
	router.GET("/sing-in", h.singInPage())
	router.POST("/sing-up", h.handlerRegisterUser())
	router.POST("/sing-in", h.handlerLoginUser())

	authGroup := router.Group("/auth", h.userIdentity)
	{
		authGroup.GET("/pepepe", h.handlePe())
	}
	postsGroup := router.Group("/posts", h.userIdentity)
	{
		postsGroup.GET("/", h.GetPosts())
		postsGroup.GET("/create", h.CreatePosts())
		postsGroup.POST("/create", h.HandlerCreatePost())
		postsGroup.GET("/:id", h.GetPost())
		postsGroup.DELETE("/:id")
	}
	return router
}
