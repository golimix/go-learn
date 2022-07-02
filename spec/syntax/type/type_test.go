package typing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTyping(t *testing.T) {
	fmt.Println(`
	Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
	TypeName  = identifier | QualifiedIdent .
	TypeArgs  = "[" TypeList [ "," ] "]" .
	TypeList  = Type { "," Type } .
	TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
			SliceType | MapType | ChannelType .`)
}

func TestBool(t *testing.T) {
	println(true)
	println(false)
}

func TestInteger(t *testing.T) {
	// 8位无符号整型
	var a uint8
	println(a)
	a = 100
	// a = 256 // 超出8位所能表达的最大范围
	fmt.Printf("%s超出%s范围\n", "a = 256\t\t", "uint8")
	println(a)
	// 16位无符号整型
	var b uint16
	println(b)
	b = 256
	print(b)
	// b = 65536 // 超出16位所能表达的最大范围
	fmt.Printf("%s超出%s范围\n", "a = 65536\t", "uint16")
	// 32位无符号整型
	var c uint32
	println(c)
	c = 65536
	c = 0x1p31
	// c = 0x1p32 // 超出32位所能表达的最大范围
	fmt.Printf("%s超出%s范围\n", "a = 0x1p32\t", "uint32")
	var d uint64
	println(d)
	d = 0x1p32
	println(d)
	// d = 0x1p64
	fmt.Printf("%s超出%s范围\n", "a = 0x1p64\t", "uint64")
}

func TestFloat(t *testing.T) {
	var a float32
	print(a)
	a = 1.0
	fmt.Printf("%f\n", a)
	a = 3.4e38
	fmt.Printf("%s超出%s范围\n", "a = 3.4e38\t", "float32")
	var b float64
	fmt.Printf("%.10f", b)
	b = 1.8e307
	// b = 1.8e308
	fmt.Printf("%.10f\n", b)
	fmt.Printf("%.20f\n", b)
	fmt.Printf("%.30f\n", b)
	fmt.Printf("%.40f\n", b)

	// 默认的自动推断的浮点类型是float64
	c := 3.0
	println(reflect.TypeOf(c).String())
}

func TestString(t *testing.T) {
	var err interface{}
	defer func() {
		if err = recover(); err != nil {
			fmt.Println("main ==》 ", err)
		}
	}()
	a := "abcde"
	fmt.Printf("%c\t%c\t%c\t%c\t%c\t\n", a[0], a[1], a[2], a[3], a[4])
	fmt.Printf("%c\t%c\t%c\t%c\t%c\t%c\n", a[0], a[1], a[2], a[3], a[4], a[5])
}

func TestArray(t *testing.T) {
	println(`
		ArrayType   = "[" ArrayLength "]" ElementType .
		ArrayLength = Expression .
		ElementType = Type .`)
	arr := [3]string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())
	arr = [...]string{"hello", "cool", "limix"}
	fmt.Printf("%s, 长度:%d, 容量:%d 类型:%s\n", arr, len(arr), cap(arr), reflect.TypeOf(arr).String())

	// TODO 向数据添加一个元素

}

/**
 * 结论: 切片具有len和cap两个属性
 * 结论: 切片的len和cap可以通过内置函数len()和cap()获取
 * 结论: 切片可通过,切片字面量的方式定义, `[]string{"hello", "cool", "limix"}`
 * 结论: 切片可通过,make函数定义,需要指定长度,会被初始化为0值 `make([]int, 3)` `make([]int,3,10)`
 * 结论: 切片新增元素可以通过 `append()` 本质是追加元素而不是扩展容量
 * 结论: 切片新增元素可以通过 `copy()`
 */
func TestSlice(t *testing.T) {
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
