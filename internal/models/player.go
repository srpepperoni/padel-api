package models

import (
	"encoding/json"

	"k8s.io/klog"
)

type PlayerJSON struct {
	PlayerName string `json:"playerName"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
}

type Player struct {
	PlayerId int `json:"PlayerId" gorm:"primaryKey"`
	Attrs    JSONMap
}

type PlayerAttrs struct {
	PlayerName string `json:"playerName"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
}

func NewPlayer(name string, lastName string, playerName string) *Player {
	playerAttrs := map[string]interface{}{
		"playerName": playerName,
		"name":       name,
		"lastName":   lastName,
	}

	return &Player{Attrs: JSONMap(playerAttrs)}
}

func (p *Player) GetAttrs() *PlayerAttrs {
	var playerAttrs PlayerAttrs
	j, err := p.Attrs.MarshalJSON()

	if err != nil {
		klog.Error(err)
	}

	err = json.Unmarshal(j, &playerAttrs)

	if err != nil {
		klog.Error(err)
	}

	return &playerAttrs
}

func (p *Player) ToPlayerJSON() *PlayerJSON {
	playerAttrs := p.GetAttrs()
	return &PlayerJSON{PlayerName: playerAttrs.PlayerName,
		Name:     playerAttrs.Name,
		LastName: playerAttrs.LastName}
}
