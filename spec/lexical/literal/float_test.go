package literal

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {

	println(`
		float_lit         = decimal_float_lit | hex_float_lit .

		decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
							decimal_digits decimal_exponent |
							"." decimal_digits [ decimal_exponent ] .
		decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .
		
		hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
		hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
							[ "_" ] hex_digits |
							"." hex_digits .
		hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
	`)

	// 标准十进制
	fmt.Printf("%f\n", 72.40)
	// 二进制
	fmt.Printf("%f\n", 0x1p-2) // =1/2*2
	fmt.Printf("%f\n", 0x2p10) // = 2*2的10次方=2048
	// 十进制
	fmt.Printf("%f\n", 1.e+0)
	fmt.Printf("%f\n", 1.e+0)
}
