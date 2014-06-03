package sworld

import "errors"

/**
 * Redeploy Region A to Region B
 * Region A and B must your own territory
 * Region A must contain at least 1 soldier
 */

func DeployRegions(race *Race, atlas *Atlas, region_id_A int, region_id_B int) error {

	if region_id_A < 0 || region_id_A >= atlas.Size() ||
		region_id_B < 0 || region_id_B >= atlas.Size() {
		return errors.New("Region ID out of range")
	}

	region_A := atlas.Region(region_id_A)
	region_B := atlas.Region(region_id_B)

	if !race.Territory().Has(region_id_A) ||
		!race.Territory().Has(region_id_B) {
		return errors.New("Must deploy region between the regions of your territory")
	}

	if region_A.Troop().Population() == 1 {
		return errors.New("Each region must contain at least one soldier")
	}

	region_A.Troop().SetPopulation(region_A.Troop().Population() - 1)
	region_B.Troop().SetPopulation(region_B.Troop().Population() + 1)

	return nil
}

func DeployIdle(race *Race, atlas *Atlas, region_id int) error {

	if region_id < 0 || region_id >= atlas.Size() {
		return errors.New("Region ID out of range")
	}

	if race.Deployable() == 0 {
		return errors.New("Have no idle soldier")
	}

	region := atlas.Region(region_id)

	if !race.Territory().Has(region_id) {
		return errors.New("Must deploy to your own territory")
	}

	region.Troop().SetPopulation(region.Troop().Population() + 1)
	race.SetDeployable(race.Deployable() - 1)

	return nil
}
