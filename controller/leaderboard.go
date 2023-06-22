package controller

func GetScores(c *gin.Context) {
	scores := []models.Leaderboard{}
	config.DB.Find(&scores)
	c.JSON(200, &scores)
}

func CreateScore(c *gin.Context) {
	var score = models.Leaderboard
	c.BindJSON(&score)
	config.DB.Create(&score)
	c.JSON(200, &score)
}