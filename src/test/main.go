package main

import (
	"fmt"
	"strconv"
	. "sworld"
)

func main() {

	tiny_atlas := CreateTinyAtlas()

	var players [2]*Player
	players[0] = CreatePlayer(0, "PLAYER 1")
	players[1] = CreatePlayer(1, "PLAYER 2")

	for round := 1; round <= 10; round++ {

		fmt.Printf("====================== Round %d ======================\n", round)

		for _, player := range players {

			fmt.Printf("%s's turn\n", player.Name())

			// if race is not chosen, chose it!
			for player.Race() == nil {

				fmt.Println("Please choose a race. Insert the combo ID")
				fmt.Println("1: Humans")
				fmt.Println("2: Orcs")
				fmt.Println("To see the map, insert 'm'")

				var race_code string
				fmt.Scanf("%s", &race_code)

				switch race_code {
				case "m":
					tiny_atlas.Show()
				default:
					race_id, err := strconv.Atoi(race_code)

					if err == nil {
						switch race_id {
						case 1:
							player.SetRace(CreateHumans(player.Id()))
						case 2:
							player.SetRace(CreateOrcs(player.Id()))
						default:
							AlertError("Wrong race code")
						}
					} else {
						AlertError("Wrong race code")
					}

				}

			}

			// conquer region
		ConquerStage:
			for {

				AlertInfo(fmt.Sprintf("%s %d", player.Race().Name(), player.Race().Deployable()))

				fmt.Println("Please choose a region to conquer. Insert the region ID")
				fmt.Println("Or insert 'r' to redeploy troops")
				fmt.Println("To see the map, insert 'm'")

				var region_code string
				fmt.Scanf("%s", &region_code)

				switch region_code {
				case "r":
					break ConquerStage // move to redeployment stage
				case "m":
					tiny_atlas.Show()

				default:
					region_id, err := strconv.Atoi(region_code)

					if err == nil {
						if ConquerRegion(player.Race(), tiny_atlas, region_id) {
							AlertSuccess(fmt.Sprintf("Succesfully conquered region %d", region_id))
						}
					} else {
						AlertError("Wrong region code")
					}
				}

			}
			fmt.Println("redeploy")
			fmt.Println("make money")

		}
	}

}
