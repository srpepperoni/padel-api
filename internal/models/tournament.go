package models

type TournamentJSON struct {
	Icon         string `json:"icon"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Rounds       int    `json:"rounds"`
	ActualRounds int    `json:"actualRound"`
}

type Tournament struct {
	TournamentID int `json:"tournamentID" gorm:"primaryKey;autoIncrement:true"`
	Attrs        JSONMap
}

func NewTournament(icon string, name string, description string, rounds int, actualRound int) *Tournament {
	tournamentAttrs := map[string]interface{}{
		"icon":        icon,
		"name":        name,
		"description": description,
		"rounds":      rounds,
		"actualRound": actualRound,
	}

	return &Tournament{Attrs: JSONMap(tournamentAttrs)}
}

func (t *Tournament) GetName() string {
	attrs := t.Attrs
	return attrs["name"].(string)
}

func (t *Tournament) GetDescription() string {
	attrs := t.Attrs
	return attrs["description"].(string)
}

func (t *Tournament) GetRounds() int {
	attrs := t.Attrs
	return attrs["rounds"].(int)
}

func (t *Tournament) GetActualRound() int {
	attrs := t.Attrs
	return attrs["actualRound"].(int)
}

func (t *Tournament) SetName(name string) {
	attrs := t.Attrs
	attrs["name"] = name
	t.Attrs = JSONMap(attrs)
}

func (t *Tournament) SetDescription(description string) {
	attrs := t.Attrs
	attrs["description"] = description
	t.Attrs = JSONMap(attrs)
}

func (t *Tournament) SetRounds(rounds int) {
	attrs := t.Attrs
	attrs["rounds"] = rounds
	t.Attrs = JSONMap(attrs)
}

func (t *Tournament) SetActualRound(actualRound int) {
	attrs := t.Attrs
	attrs["actualRound"] = actualRound
	t.Attrs = JSONMap(attrs)
}
