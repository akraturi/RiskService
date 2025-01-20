package server

import (
	"github.com/gin-gonic/gin"
)

func Run() error {
	configureLogging()

	r := gin.Default()
	configureRoutes(r)

	r.Use(gin.Recovery())

	return r.Run()
}
