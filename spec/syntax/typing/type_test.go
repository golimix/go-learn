package typing

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestTypingSpec(t *testing.T) {
	fmt.Println(`
		// 1. 类型的具体含义
		类型确定一组值以及特定于这些值的操作和方法,内部确定了类型的字节长度,类型的结构
		类型是一组相似值的集合
		在这一组值上我们可以进行相同的操作
		
		// 2. 类型的逻辑部分
		name, args, literal

		// 3. 类型的基础规范
		Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
		TypeName  = identifier | QualifiedIdent .
		TypeArgs  = "[" TypeList [ "," ] "]" .
		TypeList  = Type { "," Type } .
		TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
				SliceType | MapType | ChannelType .		
	`)
}

func TestTypingBool(t *testing.T) {
	println(true)
	println(false)
}

// 整型的字节长度决定了其所能存储的数据范围
func TestTypingInteger(t *testing.T) {
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

	// 检测int类型的字节数
	fmt.Println("Int的大小=%i", strconv.IntSize)
}

func TestTypingFloat(t *testing.T) {
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

func TestTypingString(t *testing.T) {
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

// Values of string type are immutable
func changeString(email string) string {
	fmt.Printf("email: %v\n", email)
	email = "normal"
	return email
}

func changeStringPoint(email *string) string {
	*email = "point"
	return *email
}

// 测试字符串参数传递
func TestTypingStringParams(t *testing.T) {
	fmt.Println("type string 的值是immutable")
	user := "init"
	user1 := changeString(user)
	fmt.Printf("%p,%p\n", &user, &user1)
	user2 := changeStringPoint(&user)
	fmt.Printf("%p,%p,%p\n", &user, &user1, &user2)

	user3 := ([]byte)("init")
	println(&user3[0], "\n")
}
