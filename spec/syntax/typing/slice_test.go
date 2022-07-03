package typing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypingSliceStruct(t *testing.T) {
	println(`
		type slice struct {
			array unsafe.Pointer  //指针
			len   int             //长度
			cap   int             //容量
		}
	`)
	demo := []int{1, 2, 3}
	println(len(demo), cap(demo), demo[0], demo[1], demo[2])
	println(&demo)
}

/**
 * 结论: 切片具有len和cap两个属性
 * 结论: 切片的len和cap可以通过内置函数len()和cap()获取
 * 结论: 切片可通过,切片字面量的方式定义, `[]string{"hello", "cool", "limix"}`
 * 结论: 切片可通过,make函数定义,需要指定长度,会被初始化为0值 `make([]int, 3)` `make([]int,3,10)`
 * 结论: 切片新增元素可以通过 `append()` 本质是追加元素而不是扩展容量
 * 结论: 切片新增元素可以通过 `copy()`
 */
func TestTypingSlice(t *testing.T) {
	println(`
		SliceType = "[" "]" ElementType .
	`)
	arr := []string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())

	// 通过new构造一个切片
	newarr := new([100]int)[0:50]
	fmt.Printf("%d, 长度:%d, 容量:%d 类型:%s\n", newarr, len(newarr), cap(newarr), reflect.TypeOf(newarr).String())

	// 通过make构造一个切片
	arr1 := make([]int, 3)
	fmt.Printf("%d, 长度:%d, 容量:%d 类型:%s\n", arr1, len(arr1), cap(arr1), reflect.TypeOf(arr1).String())
	arr = make([]string, 3, 10)
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())
	println(arr[0], arr[1], arr[2])
	var err interface{}
	defer func() {
		if err = recover(); err != nil {
			fmt.Println("main ==》 ", err)
			arr = append(arr, "4")
			fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())

			// 向切片插入数据
			// 向切片插入数据-append链式操作
			theNew := append([]string{"insert"}, arr[1:]...)
			arr = append(arr[:1], theNew...)
			fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())

			// 向切片插入数据-copy
			i := 1
			arr = append(arr, "")
			copy(arr[i+1:], arr[i:])
			arr[i] = "新增"
			fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())
		}
	}()
	println(arr[0], arr[1], arr[2], arr[3])

}
