package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

				race_code := ReadCommand()
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

				tiny_atlas.Show()

				AlertInfo(fmt.Sprintf("%s %d", player.Race().Name(), player.Race().Deployable()))

				fmt.Println("Please choose a region to conquer. Insert the region ID")
				fmt.Println("Or insert 'r' to redeploy troops")

				region_code := ReadCommand()

				switch region_code {
				case "r":
					break ConquerStage // move to redeployment stage
				// add other options
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

		RedeployStage:
			for {
				tiny_atlas.Show()

				AlertInfo(fmt.Sprintf("%s %d", player.Race().Name(), player.Race().Deployable()))

				fmt.Println("Please redeploy your troop. Insert a Region ID pair (e.g. 3 5)")
				fmt.Println("Or insert 'd' to finish your job")

				deploy_code := ReadCommand()

				switch deploy_code {
				case "d":
					break RedeployStage // move to redeployment stage
				// add other options
				default:

					words := strings.Fields(deploy_code)
					if len(words) < 2 {
						AlertError("Too less region")
						continue
					}

					fmt.Printf("%d %d\n\n", words[0], words[1])
					region_id_A, err1 := strconv.Atoi(words[0])
					region_id_B, err2 := strconv.Atoi(words[1])

					if err1 == nil && err2 == nil {

						if RedeployRegion(player.Race(), tiny_atlas, region_id_A, region_id_B) {
							AlertSuccess(fmt.Sprintf("Succesfully deployed your troops"))
						}

					} else {
						AlertError("Wrong deploy code")
					}
				}

			}

			fmt.Println("make money")

		}
	}

}

func ClearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func ReadCommand() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
