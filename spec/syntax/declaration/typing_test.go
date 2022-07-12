package declaration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSyntaxDeclarationTypeSpec(t *testing.T) {
	println(`
		类型声明
		TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
		TypeSpec = AliasDecl | TypeDef .

		# 1. 类型别名
		AliasDecl = identifier "=" Type .

		# 2. 新类型
		TypeDef = identifier [ TypeParameters ] Type .
	`)
}

type (
	Point  struct{ x, y float64 } // Point and struct{ x, y float64 } are different types
	polar  Point                  // polar and Point denote different types
	Point1 = Point
)

func TestDeclarationTypeAlias(t *testing.T) {
	println(`
		别名类型是相同的类型
		type Point1 = Point
	`)
	a := Point{1, 1}
	fmt.Printf("reflect.TypeOf(a): %v\n", reflect.TypeOf(a))
	a = Point1{1, 1}
	b := Point1{1, 1}
	fmt.Printf("reflect.TypeOf(a): %v\n", reflect.TypeOf(b))
	fmt.Printf("a: %v\n", a)
}

func TestDeclarationTypeDef(t *testing.T) {
	println(`
		非别名类型定义的是新的类型,但是可以通过强制转换进行类型转换
		polar  Point
	`)
	a := polar{1, 1}
	b := Point1{1, 1}
	fmt.Printf("reflect.TypeOf(a).Kind().String(): %v\n", reflect.TypeOf(a).Kind().String())
	fmt.Printf("reflect.TypeOf(b).Kind().String(): %v\n", reflect.TypeOf(b).Kind().String())
	a = polar(b)
	fmt.Printf("reflect.TypeOf(a): %v\n", reflect.TypeOf(a))
	fmt.Printf("a: %v\n", a)
}
