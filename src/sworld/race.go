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
}

func CreateRace() *Race {
	r := new(Race)
	r.territory = CreateTerritory()
	// following special power is switched-off by default
	r.airborne = false
	r.seafaring = false
	r.immortal = false
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

func (r *Race) Seafaring() bool {
	return r.seafaring
}

func (r *Race) Airborne() bool {
	return r.airborne
}

func (r *Race) Immortal() bool {
	return r.immortal
}
