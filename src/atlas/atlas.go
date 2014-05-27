package atlas

import (
	"fmt"
	"region"
)

type Atlas struct {
	regions []*region.Region
}

func CreateAtlas(num int) *Atlas {
	a := new(Atlas)
	a.regions = make([]*region.Region, num)
	return a
}

func (a *Atlas) Show() {
	fmt.Println("- * - * - * - * - * - * - * - * - * - * - * - * - * -")
	fmt.Println("ID\tType\tDef\tAdjacent\tBelonging")
	for i := 0; i < len(a.regions); i++ {
		fmt.Printf("%d", i)
		a.regions[i].Show()
		fmt.Println()
	}
	fmt.Println("- * - * - * - * - * - * - * - * - * - * - * - * - * -")
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

func (a *Atlas) SetRegionDefense(region_id int, d int) {
	a.regions[region_id].SetDefense(d)
}

func (a *Atlas) SetRegionBelonging(region_id int, player_id int) {
	a.regions[region_id].SetBelonging(player_id)
}

func CreateTinyAtlas() *Atlas {

	// create an atlas with four regions
	a := CreateAtlas(4)

	a.regions[0] = region.CreateRegion(region.Plain)

	a.regions[0].AdjacentTo(3)
	a.regions[0].AdjacentTo(1)

	a.regions[1] = region.CreateRegion(region.Swamp)
	a.regions[1].AdjacentTo(0)
	a.regions[1].AdjacentTo(2)

	a.regions[2] = region.CreateRegion(region.Hill)
	a.regions[2].AdjacentTo(1)
	a.regions[2].AdjacentTo(3)

	a.regions[3] = region.CreateRegion(region.Farm)
	a.regions[3].AdjacentTo(0)
	a.regions[3].AdjacentTo(2)

	return a
}
