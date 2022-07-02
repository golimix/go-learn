package literal

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	println(`
		整型字面量,十进制、二进制、八进制、十六进制
		int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
		decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
		binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
		octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
		hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .
		
		decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
		binary_digits  = binary_digit { [ "_" ] binary_digit } .
		octal_digits   = octal_digit { [ "_" ] octal_digit } .
		hex_digits     = hex_digit { [ "_" ] hex_digit } .`)

	println()
	PrintInteger(0b100100)
	PrintInteger(36)
	PrintInteger(0o44)
	PrintInteger(0x24)
}

func PrintInteger(a int) {
	fmt.Printf("二进制=%#b\t十进制=%d\t八进制=%#o\t十六进制=%#x\n", a, a, a, a)
}
