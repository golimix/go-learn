package typing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypingArraySpec(t *testing.T) {
	println(`
		// 1. 数组类型 

		// 2. 数组类型规范
		ArrayType   = "[" ArrayLength "]" ElementType .
		ArrayLength = Expression .
		ElementType = Type .

		Production  = production_name "=" [ Expression ] "." .
		Expression  = Alternative { "|" Alternative } .
		Alternative = Term { Term } .
		Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
		Group       = "(" Expression ")" .
		Option      = "[" Expression "]" .
		Repetition  = "{" Expression "}" .
	`)
}

func TestTypingArray(t *testing.T) {
	println(`
		ArrayType   = "[" ArrayLength "]" ElementType .
		ArrayLength = Expression .
		ElementType = Type .`)
	arr := [3]string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())
	arr = [...]string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())

	// TODO 向数据类型中添加一个元素
}
