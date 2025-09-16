package main

import (
	"github.com/arashiaslan/forum-app-go/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}