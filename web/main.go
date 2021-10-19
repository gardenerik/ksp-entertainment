package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zahradnik.xyz/ksp-entertainment/web/api"
)

func RunWebServer(port int) {
	r := gin.Default()

	api.RegisterRoutes(r.Group("/api"))

	err := r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}
}
