package sworld

import (
	. "ct"
	"fmt"
)

func AlertSuccess(s string) {
	ChangeColor(Green, true, None, false)
	fmt.Println(s)
	ResetColor()
}

func AlertError(s string) {
	ChangeColor(Red, true, None, false)
	fmt.Println(s)
	ResetColor()
}

func AlertInfo(s string) {
	ChangeColor(Blue, true, None, false)
	fmt.Println(s)
	ResetColor()
}
