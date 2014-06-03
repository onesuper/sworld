package sworld

func Scoring(player *Player, atlas *Atlas) int {

	income := 0
	race := player.Race()

	for _, id := range race.Territory().List() {
		region := atlas.Region(id)

		// Scoring via troop
		income += 1

		if race.IsMechant() {
			income += 1
		}

		if region.IsForest() && race.IsForester() {
			income += 1
		}

		if region.IsFarm() && race.IsFarmer() {
			income += 1
		}

	}

	player.GainCoins(income)

	return income
}
