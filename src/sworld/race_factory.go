package sworld

func CreateHumans(player_id int) *Race {
	r := CreateRace()
	r.name = "Humans"
	r.symbol = "HM"
	r.deployable = 30
	r.player_id = player_id
	return r
}

func CreateOrcs(player_id int) *Race {
	r := CreateRace()
	r.name = "Orcs"
	r.symbol = "OR"
	r.deployable = 30
	r.player_id = player_id

	return r
}

func CreateHalflings(player_id int) *Race {
	r := CreateRace()
	r.name = "Halflings"
	r.symbol = "HF"
	r.deployable = 30
	r.player_id = player_id
	r.airborne = true
	return r
}
