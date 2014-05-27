package player

import (
	"race"
)

type Player struct {
	id    int
	name  string
	money int
	race  *race.Race
}

func CreatePlayer(id int, name string) *Player {
	p := new(Player)
	p.id = id
	p.name = name
	p.money = 0
	return p
}

func (p *Player) GetRace() *race.Race {
	return p.race
}

func (p *Player) SetRace(r *race.Race) {
	p.race = r
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetId() int {
	return p.id
}

func (p *Player) GetMoney() int {
	return p.money
}
