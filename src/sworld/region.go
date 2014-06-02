package sworld

type Region struct {
	id        int
	terrain   int
	adjacent  *Neighbors
	belonging int
	lord      *Race
	troops    int
	tokens    *Token
	border    bool
}

func CreateRegion(ter int, bord bool) *Region {
	r := new(Region)
	r.terrain = ter
	r.adjacent = CreateNeighbors()
	r.belonging = -1
	r.lord = nil
	r.troops = 0
	r.border = bord
	return r
}

func (r *Region) Adjacent() []int {
	return r.adjacent.List()
}

func (r *Region) Troops() int {
	return r.troops
}

func (r *Region) Defense() int {
	return 3
}

func (r *Region) Terrain() int {
	return r.terrain
}

func (r *Region) TerrainStr() string {
	return NumToStr(r.terrain)
}

func (r *Region) Lord() *Race {
	return r.lord
}

func (r *Region) Border() bool {
	return r.border
}

func (r *Region) LordSymbol() string {
	if r.lord == nil {
		return "  "
	} else {
		return r.lord.Symbol()
	}
}

func (r *Region) SetBelonging(player_id int) {
	r.belonging = player_id
}

func (r *Region) SetLord(race *Race) {
	r.lord = race
}

func (r *Region) AddToken(t *Token) {
	r.tokens = t
}

func (r *Region) AdjacentTo(id int) {
	if r == nil {
		return
	}

	r.adjacent.Add(id)
}
