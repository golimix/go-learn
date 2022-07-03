package typing

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

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
		a	%s	%s
		b	%s	%s
		m	%s	%s
		n	%s	%s
		p	%s	%s
		s	%s	%s
		t	%s	%s
		u	%s	%s
		`,
		reflect.TypeOf(a).String(), reflect.TypeOf(a).Kind(),
		reflect.TypeOf(b).String(), reflect.TypeOf(b).Kind(),
		reflect.TypeOf(m).String()+"\t"+reflect.TypeOf(m).Elem().String(), reflect.TypeOf(m).Kind(),
		reflect.TypeOf(n).String()+"\t"+reflect.TypeOf(n).Elem().String(), reflect.TypeOf(n).Kind(),
		reflect.TypeOf(p).String(), reflect.TypeOf(p).Kind(),
		reflect.TypeOf(s).String(), reflect.TypeOf(s).Kind(),
		reflect.TypeOf(t).String(), reflect.TypeOf(t).Kind(),
		reflect.TypeOf(u).String(), reflect.TypeOf(u).Kind(),
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
