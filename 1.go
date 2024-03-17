package main

import (
	. "fmt"
	"strings"
)

func calculate(expr string) int {

	parts := strings.Split(expr, "")

	_ = parseNumber(parts[0])
	_ = parseNumber(parts[1])

	switch parts[1] {

	case "+":

		return 'a' + 'b'

	case "-":

		return 'a' - 'b'

	case "*":

		return 'a' * 'b'

	case "/":

		return 'a' / 'b'

	default:

		panic("unknown operator")

	}

}

func parseNumber(num string) int {

	switch num {

	case "I":

		return 1

	case "II":

		return 2

	case "III":

		return 3

	case "IV":

		return 4

	case "V":

		return 5

	case "VI":

		return 6

	case "VII":

		return 7

	case "VIII":

		return 8

	case "IX":

		return 9

	case "X":

		return 10

	default:
		panic("unknown number:" + num)
	}
	return 0
}

func main() {

	var expr string
	Println("ввод:")
	var _, err = Scan(expr)
	if err != nil {
		return
	}

	_ = calculate(expr)
	Println("равно:")
}
