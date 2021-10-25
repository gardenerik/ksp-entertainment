package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zahradnik.xyz/ksp-entertainment/web/api"
	"zahradnik.xyz/ksp-entertainment/web/frontend"
)

func RunWebServer(port int) {
	r := gin.Default()
	r.LoadHTMLGlob("assets/templates/*")
	r.Static("static/", "assets/static/")

	frontend.RegisterRoutes(r.Group("/"))
	api.RegisterRoutes(r.Group("/api"))

	err := r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}
}
