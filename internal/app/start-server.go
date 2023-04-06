package app

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"smth/cmd/config"
	"smth/internal/handler"
	"smth/internal/store/sqlstore"
	"smth/pkg/auth"
	"time"
)

func Start(config *config.Config) error {
	db, err := newDB(config.DataBase.Name, config.DataBase.Address)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	manager := auth.NewManager(config.App.SecretKey)
	corsSettings := CorsSettings()
	handlers := handler.New(store, manager, corsSettings)
	srv := newServer(handlers.ConfigureRouter())

	s := &http.Server{
		Addr:           config.App.Address,
		Handler:        srv,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()

}

func newDB(databaseName, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(databaseName, databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CorsSettings() gin.HandlerFunc {
	c := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowHeaders: []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
	})
	return c
}
