package race

import (
	"fmt"
	"territory"
)

type Race struct {
	name       string
	deployable int
	territory  *territory.Territory
}

func CreateHuman() *Race {
	r := new(Race)
	r.name = "HUMAN"
	r.deployable = 10
	r.territory = territory.CreateTerritory()
	return r
}

func CreateOrc() *Race {
	r := new(Race)
	r.name = "ORC"
	r.deployable = 10
	r.territory = territory.CreateTerritory()
	return r
}

func (r *Race) Show() {
	fmt.Printf("%s", r.name)
	fmt.Printf("\t%d\t", r.deployable)
	r.territory.Show()
	fmt.Println()
}

func (r *Race) AddTerritory(region_id int) {
	r.territory.Add(region_id)

}

func (r *Race) HasTerritory(region_id int) bool {
	return r.territory.HasRegion(region_id)
}

func (r *Race) GetDeployable() int {
	return r.deployable
}

func (r *Race) SetDeployable(n int) {
	r.deployable = n
}
