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

func TestTypingArraySample(t *testing.T) {
	arr := [3]string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())
	arr = [...]string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())

	// 数据下标访问时,不能超出其范围
	// arr[3] = "limix"
	fmt.Printf("%s 数组下标访问越界\n", `arr[3]="limix"`)

	// TODO 向数据类型中添加一个元素

}
