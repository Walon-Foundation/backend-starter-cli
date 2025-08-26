package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	app := gin.Default()

	app.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"Your server is up and running",
		})
	})

	app.Run(":3000")
}