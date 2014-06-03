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

func (p *Player) Race() *Race {
	return p.race
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Id() int {
	return p.id
}

func (p *Player) Coins() int {
	return p.coins
}

func (p *Player) SetRace(r *Race) {
	p.race = r
}

func (p *Player) GainCoins(c int) {
	p.coins += c
}
