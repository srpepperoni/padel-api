package models

type Match struct {
	MatchId           int    `json:"matchId" gorm:"primaryKey;autoIncrement:true"`
	PlayerOne         int    `json:"playerOne"`
	PlayerTwo         int    `json:"playerTwo"`
	PlayerThree       int    `json:"playerThree"`
	PlayerFour        int    `json:"playerFour"`
	Status            string `json:"status"`
	Result            string `json:"result"`
	MatchType         string `json:"matchType"`
	TournamentID      int    `json:"tournamentID"`
	ResultSetA1       int    `json:"resultSetA1"`
	ResultSetA2       int    `json:"resultSetA2"`
	ResultSetA3       int    `json:"resultSetA3"`
	ResultSetB1       int    `json:"resultSetB1"`
	ResultSetB2       int    `json:"resultSetB2"`
	ResultSetB3       int    `json:"resultSetB3"`
	CurrentSetResultA int    `json:"currentSetResultA"`
	CurrentSetResultB int    `json:"currentSetResultB"`
}

type MatchesAndPlayers struct {
	Players []Player
	Matches []Match
}
