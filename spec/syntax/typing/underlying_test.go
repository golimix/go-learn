package typing

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// 类型的基础类型
func TestUnderlying(test *testing.T) {

	type A string
	type B A
	type M map[string]int
	type N M
	type P *N
	type S string
	type T map[S]int
	type U T

	var a A
	var b B
	var m M
	var n N
	var p P
	var s S
	var t T
	var u U
	fmt.Printf(`
		a	%20s	%10s
		b	%20s	%10s
		m	%20s	%10s
		n	%20s	%10s
		p	%20s	%10s
		s	%20s	%10s
		t	%20s	%10s
		u	%20s	%10s
		`,
		reflect.TypeOf(a).String(), reflect.TypeOf(a).Kind(),
		reflect.TypeOf(b).String(), reflect.TypeOf(b).Kind(),
		reflect.TypeOf(m).String(), reflect.TypeOf(m).Elem().String()+reflect.TypeOf(m).Kind().String(),
		reflect.TypeOf(n).String(), reflect.TypeOf(n).Elem().String()+reflect.TypeOf(n).Kind().String(),
		reflect.TypeOf(p).String(), reflect.TypeOf(p).Kind(),
		reflect.TypeOf(s).String(), reflect.TypeOf(s).Kind(),
		reflect.TypeOf(t).String(), reflect.TypeOf(t).Elem().String()+reflect.TypeOf(t).Kind().String(),
		reflect.TypeOf(u).String(), reflect.TypeOf(u).Elem().String()+reflect.TypeOf(u).Kind().String(),
	)

}

type M struct {
	Value int64
}

type N struct {
	Value int64
}

type X struct {
	Value int32
}

func (m M) String() string {
	return strconv.FormatInt(m.Value, 10)
}

func (m N) String() string {
	return strconv.FormatInt(int64(m.Value), 10)
}

func (m X) String() string {
	return strconv.FormatInt(int64(m.Value), 10)
}

func TestConvention(t *testing.T) {
	m := M{1024}
	n := N{2048}
	x := X{3096}
	fmt.Printf("m=%s\nn=%s\nx=%s\n", m.String(), n.String(), x.String())
	fmt.Printf("%s=%s\t\n", reflect.TypeOf(m).Kind(), reflect.TypeOf(m).String())
	fmt.Printf("%s=%s\t\n", reflect.TypeOf(n).Kind(), reflect.TypeOf(n).String())
	fmt.Printf("%s=%s\t\n", reflect.TypeOf(x).Kind(), reflect.TypeOf(x).String())
	y := N(m)
	// x := X(m)
	println(y.String())
}
