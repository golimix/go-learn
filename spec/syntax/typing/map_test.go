package typing

import (
	"fmt"
	"testing"
)

func TestTypingMapSpec(t *testing.T) {
	println(`
		MapType     = "map" "[" KeyType "]" ElementType .
		KeyType     = Type .
		ElementType = Type .
	`)
}

func TestTypingMapSample(t *testing.T) {
	println(`
		map[string]int
		map[*T]struct{ x, y float64 }
		map[string]interface{}
	`)
}

func TestTypingMapInit(t *testing.T) {
	demoMap := make(map[string]int)
	demoMap["a"] = 0
	demoMap["b"] = 0
	fmt.Printf("demoMap: %v\n", demoMap)
}

func TestTypingMapExist(t *testing.T) {
	demoMap := make(map[string]int)
	demoMap["a"] = int('a')
	demoMap["b"] = int('b')
	if v, ok := demoMap["a"]; ok {
		fmt.Printf("v: %v\n", v)
	} else {
		fmt.Printf("v: %v\n", v)
	}
}
