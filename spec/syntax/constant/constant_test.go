package syntax

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstantPredefined(t *testing.T) {
	println(`
		// Go语言默认预定义了三种常量 true, false, iota
		true false iota
	`)
}

func TestConstantIota(t *testing.T) {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Partyday
		numberOfDays // this constant is not exported
	)
	println(Sunday)
	println(Monday)
}

func TestConstantDemo(t *testing.T) {
	println(`
		const Pi float64 = 3.14159265358979323846
		const zero = 0.0         // untyped floating-point constant
		const (
			size int64 = 1024
			eof        = -1  // untyped integer constant
		)
		const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
		const u, v float32 = 0, 3    // u = 0.0, v = 3.0
	`)
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // untyped floating-point constant
	fmt.Println(Pi)
	fmt.Println(zero)
}

func TestConstantTyped(t *testing.T) {

}

func TestConstantUnTyped(t *testing.T) {
	// The default type of an untyped constant is bool, rune, int, float64, complex128
	fmt.Println("无类型常量,具有一个默认的类型,默认是bool, rune, int, float64, complex128")
	a := true
	Print(a)
	b := 'b'
	Print(b)
	c := 1024
	Print(c)
	d := 3.141592611123123123213131
	Print(d)
	e := 1 + 3.1415926i
	Print(e)
}

func Print(a any) {
	fmt.Printf("%s\t%s\n", reflect.TypeOf(a).Kind(), reflect.TypeOf(a).String())
}
