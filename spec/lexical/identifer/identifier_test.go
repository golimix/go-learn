package identifer

import (
	"fmt"
	"testing"
)

func TestIdentifier(t *testing.T) {

}

// 预定义标识符
func TestIdentifierPredeclared(t *testing.T) {
	// 预定义类型标识符
	fmt.Println("预定义标识符")
	fmt.Println("=========================================")
	fmt.Println("any comparable\nbool byte complex64 complex128 error float32 float64\n int int8 int16 int32 int64 rune string\n uint uint8 uint16 uint32 uint64 uintptr")
	fmt.Println("=========================================")
}
