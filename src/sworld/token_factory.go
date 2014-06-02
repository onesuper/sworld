package sworld

func CreateMountainToken() *Token {
	r := new(Token)
	r.name = "Mountain"
	r.symbol = "MT"
	r.def = 1
	r.profit = 0
	r.movable = false
	return r
}
