package region

import (
	"fmt"
	"race"
)

type Region struct {
	id        int
	terrain   int
	adjacent  *Neighbors
	defense   int
	belonging int
	lord      *race.Race
}

func CreateRegion(ter int) *Region {
	r := new(Region)
	r.terrain = ter
	r.adjacent = CreateNeighbors()
	if ter == Hill {
		r.defense = 1
	} else {
		r.defense = 0
	}
	r.belonging = -1
	r.lord = nil
	return r
}

func (r *Region) Show() {
	if r == nil {
		return
	}
	fmt.Printf("\t%s", NumToType(r.terrain))
	fmt.Printf("\t%d\t", r.defense)
	r.adjacent.Show()
	fmt.Printf("\t%d", r.belonging)
	if r.lord != nil {
		fmt.Printf("\t%s", r.lord.GetName())
	} else {
		fmt.Printf("\tnil")
	}
}

func (r *Region) GetDefense() int {
	return r.defense
}

func (r *Region) SetDefense(d int) {
	r.defense = d
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

func (r *Region) AdjacentTo(id int) {
	if r == nil {
		return
	}

	r.adjacent.Add(id)
}
