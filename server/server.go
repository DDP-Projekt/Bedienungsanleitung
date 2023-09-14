package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("server/html/*")
	r.StaticFS("/Artikel", http.Dir("Artikel"))
	r.StaticFS("/static", http.Dir("static"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "de.gohtml", gin.H{
			"page": "Startseite",
		})
	})

	r.GET("/p/*page", func(c *gin.Context) {
		p, _ := url.PathUnescape(c.Param("page"))
		if p == "/" {
			p = "Startseite"
		}
		c.HTML(http.StatusOK, "de.gohtml", gin.H{
			"page": strings.TrimLeft(p, "/"),
		})
	})

	r.GET("/EN/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "en.gohtml", gin.H{
			"page": "Startseite",
		})
	})

	r.GET("/EN/p/*page", func(c *gin.Context) {
		p, _ := url.PathUnescape(c.Param("page"))
		if p == "/" {
			p = "Startseite"
		}
		c.HTML(http.StatusOK, "en.gohtml", gin.H{
			"page": strings.TrimLeft(p, "/"),
		})
	})

	r.Run()
}
