package typing

import "testing"

// INFO 注意区分函数类型和函数声明,这里指的是一个类型而不是声明
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

func TestTypingFunctionsPredeclared(t *testing.T) {
	println(`
		Functions:
		append cap close complex copy delete imag len
		make new panic print println real recover	
	`)
}

// 测试变长函数-函数的参数长度可变
func TestTypingFunctionVariableParams(t *testing.T) {
	toPrint := []string{"a", "b", "c"}
	FunctionVariableParams(toPrint...)
}

func FunctionVariableParams(toPrints ...string) {
	println(toPrints)
	for _, v := range toPrints {
		print(v + "\t")
	}
	println()
}
