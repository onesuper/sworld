package sworld

import (
	"fmt"
)

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
		a.regions[0].TerrainStr(), a.regions[1].TerrainStr(), a.regions[2].TerrainStr(),
		a.regions[0].LordSymbol(), a.regions[0].Troops(),
		a.regions[1].LordSymbol(), a.regions[1].Troops(),
		a.regions[2].LordSymbol(), a.regions[2].Troops(),

		a.regions[3].TerrainStr(), a.regions[4].TerrainStr(), a.regions[6].TerrainStr(),
		a.regions[3].LordSymbol(), a.regions[3].Troops(),
		a.regions[4].LordSymbol(), a.regions[4].Troops(),
		a.regions[6].LordSymbol(), a.regions[6].Troops(),

		a.regions[5].TerrainStr(), a.regions[7].TerrainStr(),
		a.regions[5].LordSymbol(), a.regions[5].Troops(),
		a.regions[7].LordSymbol(), a.regions[7].Troops(),

		a.regions[8].TerrainStr(), a.regions[9].TerrainStr(),
		a.regions[10].TerrainStr(), a.regions[11].TerrainStr(),
		a.regions[8].LordSymbol(), a.regions[8].Troops(),
		a.regions[9].LordSymbol(), a.regions[9].Troops(),
		a.regions[10].LordSymbol(), a.regions[10].Troops(),
		a.regions[11].LordSymbol(), a.regions[11].Troops())
}

func CreateTinyAtlas() *Atlas {

	// create an atlas with four regions
	a := CreateAtlas(12)

	a.regions[0] = CreateRegion(Farm, true)

	a.regions[0].AdjacentTo(1)
	a.regions[0].AdjacentTo(3)
	a.regions[0].AdjacentTo(4)

	a.regions[1] = CreateRegion(Plain, true)
	a.regions[1].AdjacentTo(0)
	a.regions[1].AdjacentTo(2)
	a.regions[1].AdjacentTo(4)
	a.regions[1].AdjacentTo(5)
	a.regions[1].AdjacentTo(6)

	a.regions[2] = CreateRegion(Forest, true)
	a.regions[2].AdjacentTo(1)
	a.regions[2].AdjacentTo(6)
	a.regions[2].AdjacentTo(7)

	a.regions[3] = CreateRegion(Forest, true)
	a.regions[3].AdjacentTo(0)
	a.regions[3].AdjacentTo(4)
	a.regions[3].AdjacentTo(8)

	a.regions[4] = CreateRegion(Mountain, false)
	a.regions[4].AdjacentTo(0)
	a.regions[4].AdjacentTo(1)
	a.regions[4].AdjacentTo(3)
	a.regions[4].AdjacentTo(5)
	a.regions[4].AdjacentTo(8)
	a.regions[4].AdjacentTo(9)
	a.regions[4].AddToken(CreateToken("MOUNTAIN"))

	a.regions[5] = CreateRegion(Sea, false)
	a.regions[5].AdjacentTo(1)
	a.regions[5].AdjacentTo(4)
	a.regions[5].AdjacentTo(6)
	a.regions[5].AdjacentTo(9)
	a.regions[5].AdjacentTo(10)

	a.regions[6] = CreateRegion(Mountain, false)
	a.regions[6].AdjacentTo(1)
	a.regions[6].AdjacentTo(2)
	a.regions[6].AdjacentTo(5)
	a.regions[6].AdjacentTo(7)
	a.regions[6].AdjacentTo(10)
	a.regions[6].AddToken(CreateToken("MOUNTAIN"))

	a.regions[7] = CreateRegion(Swamp, true)
	a.regions[7].AdjacentTo(2)
	a.regions[7].AdjacentTo(6)
	a.regions[7].AdjacentTo(10)
	a.regions[7].AdjacentTo(11)

	a.regions[8] = CreateRegion(Swamp, true)
	a.regions[8].AdjacentTo(3)
	a.regions[8].AdjacentTo(4)
	a.regions[8].AdjacentTo(9)

	a.regions[9] = CreateRegion(Plain, true)
	a.regions[9].AdjacentTo(4)
	a.regions[9].AdjacentTo(5)
	a.regions[9].AdjacentTo(8)
	a.regions[9].AdjacentTo(10)

	a.regions[10] = CreateRegion(Farm, true)
	a.regions[10].AdjacentTo(5)
	a.regions[10].AdjacentTo(6)
	a.regions[10].AdjacentTo(7)
	a.regions[10].AdjacentTo(9)
	a.regions[10].AdjacentTo(11)

	a.regions[11] = CreateRegion(Sea, true)
	a.regions[11].AdjacentTo(7)
	a.regions[11].AdjacentTo(10)

	return a
}
