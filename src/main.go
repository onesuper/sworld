package main

import (
	"action"
	"atlas"
	"fmt"
	"player"
	"race"
	"strconv"
)

func main() {

	tiny_atlas := atlas.CreateTinyAtlas()

	var players [2]*player.Player
	players[0] = player.CreatePlayer(0, "PLAYER 1")
	players[1] = player.CreatePlayer(1, "PLAYER 2")

	for round := 1; round <= 3; round++ {

		fmt.Printf("====================== Round %d ======================\n", round)

		for _, player := range players {

			fmt.Printf("%s's turn\n", player.GetName())

			// if race is not chosen, chose it!
			for player.GetRace() == nil {

				fmt.Println("Please choose a race. Insert the combo ID")
				fmt.Println("1: HUMAN")
				fmt.Println("2: ORC")
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
							player.SetRace(race.CreateHumans(player.GetId()))
						case 2:
							player.SetRace(race.CreateOrcs(player.GetId()))
						default:
							fmt.Println("Wrong race code!")
						}
					} else {
						fmt.Println("Wrong race code!")
					}

				}

			}

			// conquer region
		ConquerStage:
			for {

				fmt.Printf("%s %d\n", player.GetRace().GetName(), player.GetRace().GetDeployable())

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
						err = action.ConquerRegion(player.GetRace(), tiny_atlas, region_id)
						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Printf("Conquering %d\n", region_id)

						}
					} else {
						fmt.Println("Wrong region code!")
					}
				}

			}
			fmt.Println("redeploy")
			fmt.Println("make money")

		}
	}

}
