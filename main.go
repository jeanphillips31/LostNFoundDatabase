package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.New()
	config.Connect()
	routes.LeaderboardRoute(router)
	router.Run(":8080")
}