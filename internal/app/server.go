package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	handler *gin.Engine
	logger  *logrus.Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.handler.ServeHTTP(w, req)
}

func newServer(handler *gin.Engine) *Server {
	s := &Server{
		handler: handler,
		logger:  logrus.New(),
	}

	return s
}
