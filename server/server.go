package server

import (
	"RiskService/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

type Server struct {
	database         service.Database
	requestValidator *validator.Validate
}

func (s Server) Run() error {
	configureLogging()

	r := gin.Default()

	s.configureRoutes(r)

	r.Use(gin.Recovery())

	return r.Run()
}

func NewServer() *Server {
	database := service.NewDatabase()
	if database == nil {
		log.Println("failed to connect to database")
		return nil
	}

	requestValidator := NewRequestValidator()
	if requestValidator == nil {
		log.Println("failed to register request validator")
		return nil
	}

	return &Server{
		database:         database,
		requestValidator: requestValidator,
	}
}
