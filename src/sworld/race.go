package sworld

type Race struct {
	name       string
	symbol     string
	deployable int
	territory  *Territory
	player_id  int

	// special power
	airborne  bool
	seafaring bool
	immortal  bool
	merchant  bool
	forester  bool
	farmer    bool
}

func CreateRace() *Race {
	r := new(Race)
	r.territory = CreateTerritory()

	// following special power is switched-off by default
	r.airborne = false
	r.seafaring = false
	r.immortal = false
	r.merchant = false
	r.forester = false
	r.farmer = false
	return r
}

func (r *Race) Name() string {
	return r.name
}

func (r *Race) Symbol() string {
	return r.symbol
}

func (r *Race) PlayerId() int {
	return r.player_id
}

func (r *Race) Deployable() int {
	return r.deployable
}

func (r *Race) Territory() *Territory {
	return r.territory
}

// setter

func (r *Race) SetDeployable(n int) {
	r.deployable = n
}

func (r *Race) IsSeafaring() bool {
	return r.seafaring
}

func (r *Race) IsAirborne() bool {
	return r.airborne
}

func (r *Race) IsImmortal() bool {
	return r.immortal
}

func (r *Race) IsMechant() bool {
	return r.merchant
}

func (r *Race) IsForester() bool {
	return r.forester
}

func (r *Race) IsFarmer() bool {
	return r.farmer
}
