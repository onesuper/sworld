package sworld

import (
	"fmt"
)

type Atlas struct {
	regions []*Region
}

func CreateAtlas(num int) *Atlas {
	a := new(Atlas)
	a.regions = make([]*Region, num)
	return a
}

func (a *Atlas) Show() {

	fmt.Printf(`
+------------------------------------------------------------------------------------+
|                             XX                        X                            |
|     0.%s                      XX       1.%s           X      2.%s                  |
|                %s x %d           X                    X                             |
|                                XX      %s x %d        X             %s x %d          |
|XXXXXXXX                       XX                     X                             |
|        X                    XXX                      X                             |
|        X                  XX                          XX                           |
|3.%s    XX           XXXXXXXX                         X  XXX XXXX                   |
|          XX      XXX         XX                     XX          XX                X|
|           XXXXXXX    4.%s      X                   XX            XX            XXX |
|          XX                     X     XXXXXX    XXX      6.%s     XXXXXXXXXXXXX    |
|  %s x %d  X     %s x %d           XXXXXX      XXXX                          X        |
|         X                       XX            XXX          %s x %d         X        |
|        X                    XXXX     5.%s       XX                       XX        |
|       X                  XXX                     XXXXXXXX               XX  7.%s   |
|      XX                  X        %s  x %d                X             XX          |
|      XX                 X                               XX            XX           |
|       XXXXX             XX                             X             XX   %s x %d   |
|         X X XXX          XXXX                      XXXXXXXXXX    XXX               |
|       XX      XXX         X  XXX              XXXXX          XXXXX                 |
|    XXX          XXXXXXXXXX     XX           XXX                 X                XX|
| XXXX    8.%s      XX            XXX        XX                   X            XXXX  |
|XX                   XX    9.%s    XXX XXXXXX        10.%s       X          XX      |
|                      X                    XXX                    X       XX   11.%s|
|    %s x %d            X                       XX                  XX     XX         |
|                      X      %s x %d            XX      %s x %d       XXXXXX   %s x %d |
|                      X                         X                       XX          |
|                      X                         X                      XX           |
|                       X                       XX                     XX            |
+------------------------------------------------------------------------------------+
`,
		a.regions[0].GetTerrain(), a.regions[1].GetTerrain(), a.regions[2].GetTerrain(),
		a.regions[0].GetLordSymbol(), a.regions[0].GetTroops(),
		a.regions[1].GetLordSymbol(), a.regions[1].GetTroops(),
		a.regions[2].GetLordSymbol(), a.regions[2].GetTroops(),

		a.regions[3].GetTerrain(), a.regions[4].GetTerrain(), a.regions[6].GetTerrain(),
		a.regions[3].GetLordSymbol(), a.regions[3].GetTroops(),
		a.regions[4].GetLordSymbol(), a.regions[4].GetTroops(),
		a.regions[6].GetLordSymbol(), a.regions[6].GetTroops(),

		a.regions[5].GetTerrain(), a.regions[7].GetTerrain(),
		a.regions[5].GetLordSymbol(), a.regions[5].GetTroops(),
		a.regions[7].GetLordSymbol(), a.regions[7].GetTroops(),

		a.regions[8].GetTerrain(), a.regions[9].GetTerrain(), a.regions[10].GetTerrain(), a.regions[11].GetTerrain(),
		a.regions[8].GetLordSymbol(), a.regions[8].GetTroops(),
		a.regions[9].GetLordSymbol(), a.regions[9].GetTroops(),
		a.regions[10].GetLordSymbol(), a.regions[10].GetTroops(),
		a.regions[11].GetLordSymbol(), a.regions[11].GetTroops())
}

/*
func (a *Atlas) CanReach(from int, to int) bool {
	for i := 0; i < len(a.regions[from].Adjacent); i++ {
		if a.regions[from].Adjacent[i] == to {
			return true
		}
	}
	return false
}
*/

func (a *Atlas) GetSize() int {
	return len(a.regions)
}

func (a *Atlas) GetRegionDefense(region_id int) int {
	return a.regions[region_id].GetDefense()
}

func (a *Atlas) SetRegionBelonging(region_id int, player_id int) {
	a.regions[region_id].SetBelonging(player_id)
}

func (a *Atlas) SetRegionLord(region_id int, race *Race) {
	a.regions[region_id].SetLord(race)
}

func (a *Atlas) GetRegionLord(region_id int) *Race {
	return a.regions[region_id].GetLord()
}

func CreateTinyAtlas() *Atlas {

	// create an atlas with four regions
	a := CreateAtlas(12)

	a.regions[0] = CreateRegion(Farm)

	a.regions[0].AdjacentTo(1)
	a.regions[0].AdjacentTo(3)
	a.regions[0].AdjacentTo(4)

	a.regions[1] = CreateRegion(Plain)
	a.regions[1].AdjacentTo(0)
	a.regions[1].AdjacentTo(2)
	a.regions[1].AdjacentTo(4)
	a.regions[1].AdjacentTo(5)
	a.regions[1].AdjacentTo(6)

	a.regions[2] = CreateRegion(Forest)
	a.regions[2].AdjacentTo(1)
	a.regions[2].AdjacentTo(6)
	a.regions[2].AdjacentTo(7)

	a.regions[3] = CreateRegion(Forest)
	a.regions[3].AdjacentTo(0)
	a.regions[3].AdjacentTo(4)
	a.regions[3].AdjacentTo(8)

	a.regions[4] = CreateRegion(Mountain)
	a.regions[4].AdjacentTo(0)
	a.regions[4].AdjacentTo(1)
	a.regions[4].AdjacentTo(3)
	a.regions[4].AdjacentTo(5)
	a.regions[4].AdjacentTo(8)
	a.regions[4].AdjacentTo(9)
	a.regions[4].AddToken(CreateToken("MOUNTAIN"))

	a.regions[5] = CreateRegion(Sea)
	a.regions[5].AdjacentTo(1)
	a.regions[5].AdjacentTo(4)
	a.regions[5].AdjacentTo(6)
	a.regions[5].AdjacentTo(9)
	a.regions[5].AdjacentTo(10)

	a.regions[6] = CreateRegion(Mountain)
	a.regions[6].AdjacentTo(1)
	a.regions[6].AdjacentTo(2)
	a.regions[6].AdjacentTo(5)
	a.regions[6].AdjacentTo(7)
	a.regions[6].AdjacentTo(10)
	a.regions[6].AddToken(CreateToken("MOUNTAIN"))

	a.regions[7] = CreateRegion(Swamp)
	a.regions[7].AdjacentTo(2)
	a.regions[7].AdjacentTo(6)
	a.regions[7].AdjacentTo(10)
	a.regions[7].AdjacentTo(11)

	a.regions[8] = CreateRegion(Swamp)
	a.regions[8].AdjacentTo(3)
	a.regions[8].AdjacentTo(4)
	a.regions[8].AdjacentTo(9)

	a.regions[9] = CreateRegion(Plain)
	a.regions[9].AdjacentTo(4)
	a.regions[9].AdjacentTo(5)
	a.regions[9].AdjacentTo(8)
	a.regions[9].AdjacentTo(10)

	a.regions[10] = CreateRegion(Farm)
	a.regions[10].AdjacentTo(5)
	a.regions[10].AdjacentTo(6)
	a.regions[10].AdjacentTo(7)
	a.regions[10].AdjacentTo(9)
	a.regions[10].AdjacentTo(11)

	a.regions[11] = CreateRegion(Sea)
	a.regions[11].AdjacentTo(7)
	a.regions[11].AdjacentTo(10)

	return a
}
