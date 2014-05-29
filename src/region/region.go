package region

import (
	"race"
	"token"
)

type Region struct {
	id        int
	terrain   int
	adjacent  *Neighbors
	belonging int
	lord      *race.Race
	troops    int
	tokens    *token.Token
}

func CreateRegion(ter int) *Region {
	r := new(Region)
	r.terrain = ter
	r.adjacent = CreateNeighbors()
	r.belonging = -1
	r.lord = nil
	r.troops = 0
	return r
}

func (r *Region) GetTroops() int {
	return r.troops
}

func (r *Region) GetDefense() int {
	return 3
}

func (r *Region) GetTerrain() string {
	return NumToStr(r.terrain)
}

func (r *Region) SetBelonging(player_id int) {
	r.belonging = player_id
}

func (r *Region) SetLord(race *race.Race) {
	r.lord = race
}

func (r *Region) GetLord() *race.Race {
	return r.lord
}

func (r *Region) GetLordSymbol() string {
	if r.lord == nil {
		return "  "
	} else {
		return r.lord.GetSymbol()
	}
}

func (r *Region) AddToken(t *token.Token) {
	r.tokens = t
}

func (r *Region) AdjacentTo(id int) {
	if r == nil {
		return
	}

	r.adjacent.Add(id)
}
