package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	//app starting point
	app := gin.Default()

	//middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"enter the url you want to be allowed"},
		AllowMethods:     []string{"PUT", "PATCH", "POST","GET","DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	//default or test route
	app.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"Your server is up and running",
		})
	})

	//app starting point
	app.Run(":3000")
}