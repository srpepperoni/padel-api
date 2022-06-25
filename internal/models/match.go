package models

type MatchJSON struct {
	TournamentID int     `json:"tournamentId"`
	Status       string  `json:"status"`
	CoupleOne    []int   `json:"coupleOne"`
	CoupleTwo    []int   `json:"coupleTwo"`
	Result       [][]int `json:"result"`
}

type Match struct {
	MatchId int `json:"matchId" gorm:"primaryKey;autoIncrement:true"`
	Attrs   JSONMap
}

type MatchesAndPlayers struct {
	Players []Player
	Matches []Match
}

type Result struct {
	SetsCounter   int
	CoupleOneSets []int
	CoupleTwoSets []int
}

func NewMatch(coupleOne_1 int, coupleOne_2 int, coupleTwo_1 int, coupleTwo_2 int, status string, tournamentID int, result Result) *Match {
	matchAttrs := map[string]interface{}{
		"tournamentID": tournamentID,
		"status":       status,
		"coupleOne": []int{
			coupleOne_1,
			coupleOne_2,
		},
		"coupleTwo": []int{
			coupleTwo_1,
			coupleTwo_2,
		},
		"result": Result{
			result.SetsCounter,
			result.CoupleOneSets,
			result.CoupleTwoSets,
		},
	}

	return &Match{Attrs: JSONMap(matchAttrs)}
}

func (m *Match) GetCoupleOne() []int {
	attrs := m.Attrs
	return attrs["coupleOne"].([]int)
}

func (m *Match) GetCoupleTwo() []int {
	attrs := m.Attrs
	return attrs["coupleTwo"].([]int)
}

func (m *Match) GetStatus() string {
	attrs := m.Attrs
	return attrs["status"].(string)
}

func (m *Match) GetTournamentID() int {
	attrs := m.Attrs
	return attrs["tournamentID"].(int)
}

func (m *Match) GetResult() Result {
	attrs := m.Attrs
	return attrs["result"].(Result)
}

func (m *Match) SetCoupleOne(idPlayerA int, idPlayerB int) {
	attrs := m.Attrs
	attrs["coupleOne"] = []int{idPlayerA, idPlayerB}
}

func (m *Match) SetCoupleTwo(idPlayerA int, idPlayerB int) {
	attrs := m.Attrs
	attrs["coupleTwo"] = []int{idPlayerA, idPlayerB}
}

func (m *Match) SetStatus(status string) {
	attrs := m.Attrs
	attrs["status"] = status
}

func (m *Match) SetTournament(tournamentID int) {
	attrs := m.Attrs
	attrs["tournamentID"] = tournamentID
}

func (m *Match) SetResult(result Result) {
	attrs := m.Attrs
	attrs["result"] = result
}
