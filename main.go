package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"os"
)

func ReverseProxy() gin.HandlerFunc {
	target := os.Getenv("API_HOST")

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Host = target
		}

		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()
	router.GET("/api/*all", ReverseProxy())
	router.POST("/api/*all", ReverseProxy())
	router.PUT("/api/*all", ReverseProxy())
	router.PATCH("/api/*all", ReverseProxy())
	router.DELETE("/api/*all", ReverseProxy())
	router.Use(static.Serve("/", static.LocalFile("public", false)))
	router.NoRoute(func(c *gin.Context) {
		c.File("public/index.html")
	})
	router.Run()
}
