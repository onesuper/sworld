package race

func CreateHumans(player_id int) *Race {
	r := CreateRace()
	r.name = "HUMANS"
	r.symbol = "HM"
	r.deployable = 30
	r.player_id = player_id
	return r
}

func CreateOrcs(player_id int) *Race {

	r := CreateRace()
	r.name = "ORCS"
	r.symbol = "OR"
	r.deployable = 30
	r.player_id = player_id
	return r
}
