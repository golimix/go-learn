package typing

import "testing"

func TestTypingFunctionSpec(t *testing.T) {
	println("the spec 定义如下")
	println(`
		FunctionType   = "func" Signature .
		Signature      = Parameters [ Result ] .
		Result         = Parameters | Type .
		Parameters     = "(" [ ParameterList [ "," ] ] ")" .
		ParameterList  = ParameterDecl { "," ParameterDecl } .
		ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
	`)

	println("函数定义样例如下")
	println(`
		func() // 无参数和返回值函数
		func(x int) int
		func(a, _ int, z float32) bool  // 不具名参数,为什么要不具名参数
		func(a, b int, z float32) (bool)
		func(prefix string, values ...int) // 可变参数
		func(a, b int, z float64, opt ...interface{}) (success bool)
		func(int, int, float64) (float64, *[]int)
		func(n int) func(p *T)
	`)
}
