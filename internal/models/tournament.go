package models

import (
	"encoding/json"

	"k8s.io/klog"
)

type TournamentJSON struct {
	Icon         string `json:"icon"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Rounds       int    `json:"rounds"`
	ActualRounds int    `json:"actualRound"`
	Players      []int  `json:"players"`
}

type Tournament struct {
	TournamentID int `json:"tournamentID" gorm:"primaryKey;autoIncrement:true"`
	Attrs        JSONMap
}

// Used to map attrs interface into usable object
type TournamentAttrs struct {
	Icon         string    `json:"icon"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Rounds       int       `json:"rounds"`
	ActualRounds int       `json:"actualRound"`
	Players      []PlayerT `json:"players"`
}

type PlayerT struct {
	PlayerID     int // ID Player
	PlayerScore  int // Actual Player's score in this tournament
	RoundsPlayed int
	Couples      []int // ID Players this player already play with
}

func NewPlayerT(playerID int) *PlayerT {
	return &PlayerT{PlayerID: playerID, PlayerScore: 0, RoundsPlayed: 0, Couples: []int{}}
}

func NewTournament(icon string, name string, description string, rounds int, actualRound int, players []PlayerT) *Tournament {
	tournamentAttrs := map[string]interface{}{
		"icon":        icon,
		"name":        name,
		"description": description,
		"rounds":      rounds,
		"actualRound": actualRound,
		"players":     players,
	}

	return &Tournament{Attrs: JSONMap(tournamentAttrs)}
}

func (t *Tournament) GetAttrs() *TournamentAttrs {
	var tournamentAttrs TournamentAttrs
	j, err := t.Attrs.MarshalJSON()

	if err != nil {
		klog.Error(err)
	}

	err = json.Unmarshal(j, &tournamentAttrs)

	if err != nil {
		klog.Error(err)
	}

	return &tournamentAttrs
}
