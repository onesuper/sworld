package region

import (
	"fmt"
)

type Region struct {
	id        int
	terrain   int
	adjacent  *Neighbors
	defense   int
	belonging int
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

func (r *Region) AdjacentTo(id int) {
	if r == nil {
		return
	}

	r.adjacent.Add(id)
}
