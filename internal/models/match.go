package models

import (
	"encoding/json"

	"k8s.io/klog"
)

// Struct for Request from client side used for API
type MatchJSON struct {
	TournamentID int     `json:"tournamentId"`
	Status       string  `json:"status"`
	CoupleOne    []int   `json:"coupleOne"`
	CoupleTwo    []int   `json:"coupleTwo"`
	Result       [][]int `json:"result"`
}

// Struct for Database
type Match struct {
	MatchId int `json:"matchId" gorm:"primaryKey;autoIncrement:true"`
	Attrs   JSONMap
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

// Used to map attrs interface into usable object
type MatchAttrs struct {
	TournamentID int    `json:"tournamentID"`
	Status       string `json:"status"`
	CoupleOne    []int  `json:"coupleOne"`
	CoupleTwo    []int  `json:"coupleTwo"`
	Result       Result `json:"result"`
}

// Used for templating injection into new-match.html
type MatchForTemplate struct {
	TournamentName string
	Status         string
	Result         Result
	CoupleOne      []PlayerJSON
	CoupleTwo      []PlayerJSON
}

type Result struct {
	SetsCounter   int   `json:"tournamentID"`
	CoupleOneSets []int `json:"coupleOneSets"`
	CoupleTwoSets []int `json:"coupleTwoSets"`
}

func (m *Match) GetAttrs() *MatchAttrs {
	var matchAttrs MatchAttrs
	j, err := m.Attrs.MarshalJSON()

	if err != nil {
		klog.Error(err)
	}

	err = json.Unmarshal(j, &matchAttrs)

	if err != nil {
		klog.Error(err)
	}

	return &matchAttrs
}
