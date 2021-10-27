package web

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"zahradnik.xyz/ksp-entertainment/web/api"
	"zahradnik.xyz/ksp-entertainment/web/frontend"
)

func customRenderer(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	templates, err := filepath.Glob(path.Join(templatesDir, "*.gohtml"))
	if err != nil {
		panic(err)
	}

	for _, template := range templates {
		base := filepath.Base(template)
		if base == "_base.gohtml" {
			continue
		}

		r.AddFromFiles(base, path.Join(templatesDir, "_base.gohtml"), template)
	}

	return r
}

func RunWebServer(port int) {
	if !viper.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.HTMLRender = customRenderer("assets/templates/")
	r.Static("static/", "assets/static/")

	frontend.RegisterRoutes(r.Group("/"))
	api.RegisterRoutes(r.Group("/api"))

	err := r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}
}
