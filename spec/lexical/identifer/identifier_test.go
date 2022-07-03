package identifer

import (
	"fmt"
	"testing"
)

func TestIdentifier(t *testing.T) {

}

// 预定义标识符
func TestIdentifierPredeclared(t *testing.T) {
	fmt.Println("预定义标识符")
	fmt.Println(`
		Types:
			any bool byte comparable
			complex64 complex128 error float32 float64
			int int8 int16 int32 int64 rune string
			uint uint8 uint16 uint32 uint64 uintptr

		Constants:
			true false iota

		Zero value:
			nil

		Functions:
			append cap close complex copy delete imag len
			make new panic print println real recover	
	`)
}
