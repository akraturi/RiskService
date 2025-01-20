package server

import "github.com/gin-gonic/gin"

func configureRoutes(r *gin.Engine) {
	v1Router := r.Group("/v1")

	v1Router.GET("/risks", getRisks)
	v1Router.POST("/risks", postRisks)
}
