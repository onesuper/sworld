package token

type Token struct {
	name   string
	def    int
	profit int
	belong int
}

func CreateToken(name string) *Token {
	t := new(Token)
	t.name = name
	t.def = 1
	t.profit = 0
	t.belong = -1
	return t
}

func CreateRaceToken(name string) *Token {
	t := new(Token)
	t.name = name
	t.def = 1
	t.profit = 0
	t.belong = -1
	return t
}

func CreateInvicableToken(name string) *Token {
	t := new(Token)
	t.name = name
	t.def = 1000
	t.profit = 0
	t.belong = -1
	return t
}
