package sworld

/**
 * Given a region number, return false if this action fails
 */
func ConquerRegion(race *Race, atlas *Atlas, region_id int) bool {

	// out of range checking
	if region_id < 0 || region_id >= atlas.Size() {
		AlertError("The region ID is out of range")
		return false
	}

	// own territory checking
	if race.HasTerritory(region_id) {
		AlertError("Can not conquer your own territory")
		return false
	}

	target_region := atlas.Region(region_id)

	// sea region checking
	if target_region.Terrain() == Sea && !race.Seafaring() {
		AlertError("Can not conquer sea region")
		return false
	}

	// enter from border checking
	if !race.HasAnyTerritory() && !race.Airborne() && !target_region.Border() {
		AlertError("Must enter from a border region")
		return false
	}

	// if a race has conquered any regions before, it has to conquer the
	// regions surrounding to its territory
	if race.HasAnyTerritory() {
		within := false
		for _, id := range target_region.Adjacent() {
			if race.HasTerritory(id) {
				within = true // find at least one way to come
				break
			}
		}
		if within == false {
			AlertError("Must conquer a region adjacent to your territory")
			return false
		}
	}

	bottom_def := target_region.Defense() + 2

	if race.Deployable() < bottom_def {
		AlertError("The defense is too high to conquer")
		return false
	}

	race.SetDeployable(race.Deployable() - bottom_def)
	race.AddTerritory(region_id)
	//tlas.SetRegionBelonging(region_id, race.PlayerId())

	// loser surrenders its territory
	loser := target_region.Lord()
	if loser != nil {
		loser.SetDeployable(loser.Deployable() + bottom_def - 3)
		loser.SurrenderTerritory(region_id)
	}

	target_region.SetLord(race)

	return true

}
