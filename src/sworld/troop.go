package sworld

type Troop struct {
	lord       *Race
	population int
}

func (t *Troop) Lord() *Race {
	if t == nil {
		return nil
	}
	return t.lord
}

func (t *Troop) Population() int {
	return t.population
}

func (t *Troop) SetPopulation(pop int) {
	t.population = pop
}

func CreateTroop(race *Race, pop int) *Troop {
	t := new(Troop)
	t.lord = race
	t.population = pop
	return t
}
