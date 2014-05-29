package race

import (
	"territory"
)

type Race struct {
	name       string
	symbol     string
	deployable int
	territory  *territory.Territory
	player_id  int
}

func CreateRace() *Race {
	r := new(Race)
	r.territory = territory.CreateTerritory()
	return r
}

func (r *Race) GetName() string {
	return r.name
}

func (r *Race) GetSymbol() string {
	return r.symbol
}

func (r *Race) GetPlayerId() int {
	return r.player_id
}

func (r *Race) AddTerritory(region_id int) {
	r.territory.Add(region_id)

}

func (r *Race) HasTerritory(region_id int) bool {
	return r.territory.HasRegion(region_id)
}

func (r *Race) SurrenderTerritory(region_id int) bool {
	return r.territory.Remove(region_id)
}

func (r *Race) GetDeployable() int {
	return r.deployable
}

func (r *Race) SetDeployable(n int) {
	r.deployable = n
}
