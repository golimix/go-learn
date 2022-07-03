package typing

import (
	"fmt"
	"strconv"
	"testing"
)

type emptyStruct struct{}

// 所有的内存大小为空,都指向了同一个地址 uintptr, 全局变量,占8个字节
// 涉及到所有内存size为0的内存分配,那么就用同一个地址 &zerobase
func TestTypingStructEmpty(t *testing.T) {
	a := struct{}{}
	b := struct{}{}
	c := emptyStruct{}

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
}

// TODO 字节填充和内存对齐等相关知识
// https://jishuin.proginn.com/p/763bfbd35690

type DemoStruct struct {
	x, y int
	u    float32
	_    float32 // padding
	A    *[]int
	F    func()
}

// 结构体成员方法
func (u DemoStruct) String() string {
	return strconv.FormatFloat(float64(u.x)+float64(u.y)+float64(u.u), 'f', 10, 64)
}

func TestTypingStructDefine(t *testing.T) {
	// 1. struct的定义
	println(`
		StructType    = "struct" "{" { FieldDecl ";" } "}" .
		FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
		EmbeddedField = [ "*" ] TypeName .
		Tag           = string_lit .

		IdentifierList = identifier { "," identifier } .
	`)

	// 2. struct 定义一个空结构体
	type x struct{}
	struct_empty := x{}
	println("%d", &struct_empty)

	limix := DemoStruct{
		x: 0.0,
		y: 1.0,
		u: 3.1415926,
	}
	fmt.Printf("%s\n", limix.String())
}

type User struct {
	email string
}

// 传递的是值,复制整个结构体,返回一个新指针
func (user User) ChangeEmail() User {
	user.email = "value@gmail.com"
	return user
}

// 传递的是值,复制一个指针,返回一个指针引用
func (user *User) ChangeEmailPoint() User {
	user.email = "point@gmail.com"
	return *user
}

// 结构体传值,结构体传指针的差异
func TestTypingStructParams(t *testing.T) {
	user := User{email: "coollimix@gmail.com"}
	println(&user)
	user1 := user.ChangeEmail()
	fmt.Printf("%p,%p\n", &user, &user1)
	println(user.email)
	user2 := user.ChangeEmailPoint()
	fmt.Printf("%p,%p,%p\n", &user, &user1, &user2)
	if user == user1 {
		panic("some wrong")
	} else {
		println("传递的是值,复制整个结构体,返回一个新指针")
		println("user != user1")
	}
	if user == user2 {
		println("传递的是值,复制一个指针,返回一个指针引用")
		println("user == user2")
	} else {
		println("user != user2")
	}
}
