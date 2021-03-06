package sworld

type Region struct {
	id        int
	terrain   int
	adjacent  *Neighbors
	belonging int
	troop     *Troop
	tokens    *Token
	border    bool
}

func CreateRegion(ter int, bord bool) *Region {
	r := new(Region)
	r.terrain = ter
	r.adjacent = CreateNeighbors()
	r.belonging = -1
	r.troop = nil
	r.border = bord
	return r
}

func (r *Region) IsMountain() bool {
	return r.terrain == Mountain
}

func (r *Region) IsFarm() bool {
	return r.terrain == Farm
}

func (r *Region) IsSwamp() bool {
	return r.terrain == Swamp
}

func (r *Region) IsHill() bool {
	return r.terrain == Hill
}

func (r *Region) IsForest() bool {
	return r.terrain == Forest
}

func (r *Region) IsSea() bool {
	return r.terrain == Sea
}

func (r *Region) Adjacent() []int {
	return r.adjacent.List()
}

func (r *Region) Troop() *Troop {
	return r.troop
}

func (r *Region) Defense() int {
	def := 0
	if r.troop != nil {
		def += r.troop.Population()
	}
	if r.tokens != nil {
		def += r.tokens.def
	}
	return def
}

func (r *Region) Terrain() int {
	return r.terrain
}

func (r *Region) TerrainStr() string {
	return NumToStr(r.terrain)
}

func (r *Region) IsBorder() bool {
	return r.border
}

func (r *Region) SetBelonging(player_id int) {
	r.belonging = player_id
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

func (r *Region) SetTroop(t *Troop) {
	r.troop = t
}
