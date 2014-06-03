package sworld

/**
 * Redeploy Region A to Region B
 * Region A and B must your own territory
 * Region A must contain at least 1 soldier
 */

func RedeployRegion(race *Race, atlas *Atlas, region_id_A int, region_id_B int) bool {

	if region_id_A < 0 || region_id_A >= atlas.Size() ||
		region_id_B < 0 || region_id_B >= atlas.Size() {
		AlertError("Region ID out of range")
		return false
	}

	region_A := atlas.Region(region_id_A)
	region_B := atlas.Region(region_id_B)

	if !race.Territory().Has(region_id_A) ||
		!race.Territory().Has(region_id_B) {
		AlertError("Must deploy region between the regions of your territory")
		return false
	}

	if region_A.Troop().Population() == 1 {
		AlertError("Each region must contain at least one soldier")
		return false
	}

	region_A.Troop().SetPopulation(region_A.Troop().Population() - 1)
	region_B.Troop().SetPopulation(region_B.Troop().Population() + 1)

	return true
}