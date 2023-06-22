package models

import "gorm.io/gorm"

type Leaderboard struct {
	gorm.Model
	Id int 'json:"id" gorm:"primary_key"'
	Name string 'json:"name"'
	Score string 'json:"score"'
	Time string 'json:"time"'
}