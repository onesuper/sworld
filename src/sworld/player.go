package sworld

type Player struct {
	id    int
	name  string
	coins int
	race  *Race
}

func CreatePlayer(id int, name string) *Player {
	p := new(Player)
	p.id = id
	p.name = name
	p.coins = 0
	return p
}

func (p *Player) GetRace() *Race {
	return p.race
}

func (p *Player) SetRace(r *Race) {
	p.race = r
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetId() int {
	return p.id
}

func (p *Player) GetCoins() int {
	return p.coins
}
