package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setup_config() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("useHTTPS", false)
	viper.SetDefault("certPath", "")
	viper.SetDefault("keyPath", "")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	json_string, err := json.MarshalIndent(viper.AllSettings(), "", "\t")
	if err == nil {
		log.Printf("using config: \n%s\n", json_string)
	}
}

func main() {
	setup_config()

	r := gin.Default()

	r.LoadHTMLGlob("server/html/*")

	g := r.Group("Bedienungsanleitung")

	g.StaticFS("/Artikel", http.Dir("Artikel"))
	g.StaticFS("/static", http.Dir("static"))

	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "de.gohtml", gin.H{
			"page": "Startseite",
		})
	})

	g.GET("/DE/*page", func(c *gin.Context) {
		p, _ := url.PathUnescape(c.Param("page"))
		if p == "/" {
			p = "Startseite"
		}
		c.HTML(http.StatusOK, "de.gohtml", gin.H{
			"page": strings.TrimLeft(p, "/"),
		})
	})

	g.GET("/EN/*page", func(c *gin.Context) {
		p, _ := url.PathUnescape(c.Param("page"))
		if p == "/" {
			p = "Startseite"
		}
		c.HTML(http.StatusOK, "en.gohtml", gin.H{
			"page": strings.TrimLeft(p, "/"),
		})
	})

	// run the server
	if viper.GetBool("useHTTPS") {
		if viper.GetString("certPath") == "" || viper.GetString("keyPath") == "" {
			log.Fatal("certPath and keyPath can not be empty!")
		}
		log.Fatal(r.RunTLS(":"+viper.GetString("port"), viper.GetString("certPath"), viper.GetString("keyPath")))
	} else {
		log.Fatal(r.Run(":" + viper.GetString("port")))
	}
}
