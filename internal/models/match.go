package models

type Match struct {
	MatchId     int    `json:"matchId" gorm:"primaryKey;autoIncrement:true"`
	PlayerOne   int    `json:"playerOne"`
	PlayerTwo   int    `json:"playerTwo"`
	PlayerThree int    `json:"playerThree"`
	PlayerFour  int    `json:"playerFour"`
	Status      string `json:"status"`
	Result      string `json:"result"`
}

type MatchesAndPlayers struct {
	Players []Player
	Matches []Match
}
