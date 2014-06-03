package sworld

func CreateHumans(player_id int) *Race {
	r := CreateRace()
	r.name = "Humans"
	r.symbol = "HM"
	r.deployable = 10
	r.player_id = player_id
	return r
}

func CreateOrcs(player_id int) *Race {
	r := CreateRace()
	r.name = "Orcs"
	r.symbol = "OR"
	r.deployable = 10
	r.player_id = player_id

	return r
}

func CreateHalflings(player_id int) *Race {
	r := CreateRace()
	r.name = "Halflings"
	r.symbol = "HF"
	r.deployable = 10
	r.player_id = player_id
	r.airborne = true
	return r
}

func CreatesElves(player_id int) *Race {
	r := CreateRace()
	r.name = "Elves"
	r.symbol = "EL"
	r.deployable = 11
	r.player_id = player_id
	r.immortal = true
	return r
}
