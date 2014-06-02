package sworld

type Atlas struct {
	regions []*Region
}

func CreateAtlas(num int) *Atlas {
	a := new(Atlas)
	a.regions = make([]*Region, num)
	return a
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

func (a *Atlas) Size() int {
	return len(a.regions)
}

func (a *Atlas) Region(region_id int) *Region {

	if region_id < 0 || region_id >= a.Size() {
		return nil
	}

	return a.regions[region_id]
}
