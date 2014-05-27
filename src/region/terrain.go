package region

const (
	Farm  = 0
	Sea   = 1
	Plain = 2
	Hill  = 3
	Wood  = 4
	Swamp = 5
)

func NumToType(num int) string {
	var t string
	switch num {
	case 0:
		t = "Farm"
	case 1:
		t = "Sea"
	case 2:
		t = "Plain"
	case 3:
		t = "Hill"
	case 4:
		t = "Wood"
	case 5:
		t = "Swamp"
	}
	return t
}
