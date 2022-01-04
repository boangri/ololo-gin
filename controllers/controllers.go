package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/boangri/ololo-gin/models"
)

func HomePage(ctx *gin.Context) {
	
	data := gin.H{
		"title": "Ололоша тим",
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func ServicesPage(ctx *gin.Context) {
	data := gin.H{
		"title": "Сервисы",
	}
	ctx.HTML(http.StatusOK, "services.html", data)
}

func PortfolioPage(ctx *gin.Context) {
	portfolio := models.GetPortfolio()
	data := gin.H{
		"title": "Портфолио",
		"portfolio": portfolio,
	}
	ctx.HTML(http.StatusOK, "portfolio.html", data)
}

func ContactsPage(ctx *gin.Context) {
	
	data := gin.H{
		"title": "Контакты",
	}
	ctx.HTML(http.StatusOK, "contacts.html", data)
}