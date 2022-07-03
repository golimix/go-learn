package typing

import (
	"reflect"
	"testing"
)

func TestTypingPointerSpec(t *testing.T) {
	println(`
		PointerType = "*" BaseType .
		BaseType    = Type .`)
}

func TestTypingPointerSample(t *testing.T) {
	var a *emptyStruct = &emptyStruct{}
	var b *[4]int
	println(reflect.TypeOf(a).String())
	println(reflect.TypeOf(b).String())
}
