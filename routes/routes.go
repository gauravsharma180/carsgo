package routes

import (
	"cars/config"
	"cars/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		var cars []models.Car
		config.DB.Find(&cars)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title":    "Dashboard",
			"cars":     cars,
			"carCount": len(cars),
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})
	r.POST("/auth/login", func(c *gin.Context) {
		// Handle login logic here
		c.JSON(http.StatusOK, gin.H{"status": "Logged in"})
	})
	// Add more routes here

	return r
}
