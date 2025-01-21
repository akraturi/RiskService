package server

import "github.com/gin-gonic/gin"

func (s *Server) configureRoutes(r *gin.Engine) {
	v1Router := r.Group("/v1")

	v1Router.GET("/risks", s.getRisks)
	v1Router.POST("/risks", s.addRisk)
	v1Router.GET("/risks/:id", s.getRiskById)
}
