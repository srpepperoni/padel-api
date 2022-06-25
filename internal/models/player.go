package models

import "fmt"

type PlayerJSON struct {
	PlayerName string `json:"playerName"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
}

type Player struct {
	PlayerId int `json:"PlayerId" gorm:"primaryKey"`
	Attrs    JSONMap
}

func NewPlayer(name string, lastName string, playerName string) *Player {
	playerAttrs := map[string]interface{}{
		"playerName": playerName,
		"name":       name,
		"lastName":   lastName,
	}

	return &Player{Attrs: JSONMap(playerAttrs)}
}

func (p *Player) GetName() string {
	attrs := p.Attrs
	name := fmt.Sprint(attrs["name"])
	return name
}

func (p *Player) GetLastName() string {
	attrs := p.Attrs
	name := fmt.Sprint(attrs["lastName"])
	return name
}

func (p *Player) GetPlayerName() string {
	attrs := p.Attrs
	name := fmt.Sprint(attrs["playerName"])
	return name
}

func (p *Player) SetName(name string) {
	attrs := p.Attrs
	attrs["name"] = name
	p.Attrs = JSONMap(attrs)
}

func (p *Player) SetLastName(lastName string) {
	attrs := p.Attrs
	attrs["lastName"] = lastName
	p.Attrs = JSONMap(attrs)
}

func (p *Player) SetPlayerName(playerName string) {
	attrs := p.Attrs
	attrs["playerName"] = playerName
	p.Attrs = JSONMap(attrs)
}
