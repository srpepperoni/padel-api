package models

type Tournament struct {
	TournamentID int    `json:"tournamentID" gorm:"primaryKey;autoIncrement:true"`
	Icon         string `json:"icon"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Rounds       int    `json:"rounds"`
}
