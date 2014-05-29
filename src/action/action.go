package action

import (
	"atlas"
	"errors"
	"race"
)

func ConquerRegion(race *race.Race, atlas *atlas.Atlas, region_id int) error {

	// out of range?
	if region_id < 0 || region_id >= atlas.GetSize() {
		return errors.New("The region ID is out of range")
	}

	// out of arrival?

	// own territory?
	if race.HasTerritory(region_id) {
		return errors.New("Can not conquer your own territory")

	}

	bottom_def := atlas.GetRegionDefense(region_id) + 2

	if race.GetDeployable() > bottom_def {
		race.SetDeployable(race.GetDeployable() - bottom_def)
		race.AddTerritory(region_id)
		atlas.SetRegionBelonging(region_id, race.GetPlayerId())

		// loser surrenders its territory
		loser := atlas.GetRegionLord(region_id)

		if loser != nil {
			loser.SetDeployable(loser.GetDeployable() + bottom_def - 3)
			loser.SurrenderTerritory(region_id)
		}
		atlas.SetRegionLord(region_id, race)

		return nil
	} else {
		return errors.New("The defense is too high to conquer")
	}
}
