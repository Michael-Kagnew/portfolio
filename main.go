package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {

	// HomeController := controllers.HomeController{}
	router := gin.Default()
	router.LoadHTMLGlob("./web_dist/build/*.html")
	router.Static("static", "./web_dist/build/static")
	// router.Use(static.Serve("/", static.LocalFile("./web_dist/build", true)))
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusConflict, "index.html", nil)
	})

	router.Run()
}
