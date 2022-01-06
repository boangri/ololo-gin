package main

import (
	"net/http"
	"github.com/boangri/ololo-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Static("/css", "./static/css")
	server.Static("/js", "./static/js")
	server.Static("/images", "./static/images")
	server.StaticFile("/favicon.ico", "./static/favicon.ico")
	server.LoadHTMLGlob("templates/*.html")

	server.GET("/", controllers.HomePage)
	server.GET("/main", controllers.HomePage)
	server.GET("/services", controllers.ServicesPage)
	server.GET("/portfolio", controllers.PortfolioPage)
	server.GET("/contacts", controllers.ContactsPage)
	server.NoRoute(func(c *gin.Context) {
		// c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		data := gin.H{
			"title": "404",
		}
		c.HTML(http.StatusNotFound, "404.html", data)
	})

	server.Run(":8080")
}