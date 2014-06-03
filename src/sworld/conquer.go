package sworld

import "errors"

/**
 * Given a region number, return false if this action fails
 */
func ConquerRegion(race *Race, atlas *Atlas, region_id int) error {

	// out of range checking
	if region_id < 0 || region_id >= atlas.Size() {
		return errors.New("The region ID is out of range")
	}

	// own territory checking
	if race.Territory().Has(region_id) {
		return errors.New("Can not conquer your own territory")

	}

	target_region := atlas.Region(region_id)

	// sea region checking
	if target_region.IsSea() && !race.IsSeafaring() {
		return errors.New("Can not conquer sea region")
	}

	// enter from border checking
	if race.Territory().IsEmpty() && !race.IsAirborne() && !target_region.IsBorder() {
		return errors.New("Must enter from a border region")
	}

	// if a race has conquered any regions before, it has to conquer the
	// regions surrounding to its territory
	if !race.Territory().IsEmpty() {
		within := false
		for _, id := range target_region.Adjacent() {
			if race.Territory().Has(id) {
				within = true // find at least one way to come
				break
			}
		}
		if within == false {
			return errors.New("Must conquer a region adjacent to your territory")

		}
	}

	bottom_def := target_region.Defense() + 2

	if race.Deployable() < bottom_def {
		return errors.New("The defense is too high to conquer")
	}

	// loser surrenders its territory fist
	loser := target_region.Troop().Lord()

	if loser != nil {
		loser.Territory().Remove(region_id)
		remains := target_region.Troop().Population()
		// loser's death
		if !loser.IsImmortal() {
			remains -= 1
		}
		// all the remains back to hand
		loser.SetDeployable(loser.Deployable() + remains)
	}

	race.Territory().Add(region_id)
	race.SetDeployable(race.Deployable() - bottom_def)
	target_region.SetTroop(CreateTroop(race, bottom_def))

	return nil

}
