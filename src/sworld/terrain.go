package sworld

const (
	Farm     = 0
	Sea      = 1
	Plain    = 2
	Mountain = 3
	Forest   = 4
	Swamp    = 5
)

func NumToStr(num int) string {
	var t string
	switch num {
	case 0:
		t = "##"
	case 1:
		t = "~~"
	case 2:
		t = "=="
	case 3:
		t = "AA"
	case 4:
		t = "YY"
	case 5:
		t = "**"
	}
	return t
}
