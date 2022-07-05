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
			bool 
			int	int8  int16  int32  int64 
			uint	uint8 uint16 uint32 uint64
				byte         rune   uintptr
			float32 float64 complex64 complex128 
			string
			error
			any comparable

		Constants:
			true false iota

		Zero value:
			nil

		Functions:
			append cap close complex copy delete imag len
			make new panic print println real recover	
	`)
}

func TestIdentifierNil(t *testing.T) {
	println(`
		各种类型在没有赋值之前都会是一个0值, 在builtin.go中是预定义好的一个变量
			type Type int
			var nil Type

		bool      -> false                              
		numbers -> 0                                 
		string    -> ""      
		
		pointers -> nil
		slices -> nil
		maps -> nil
		channels -> nil
		functions -> nil
		interfaces -> nil
	`)
}
