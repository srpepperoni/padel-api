package models

type Player struct {
	PlayerId   int    `json:"PlayerId" gorm:"primaryKey"`
	PlayerName string `json:"playerName"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
}
