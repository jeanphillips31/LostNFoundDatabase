package routes

func LeaderboardRoute(router *gin.Engine) {
	router.GET("/", controller.GetScores)
	router.POST("/", controller.CreateScore)
}