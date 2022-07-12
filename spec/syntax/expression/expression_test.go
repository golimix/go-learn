package expression

import "testing"

// TODO 语法-表达式-完善整体体系
// 2022-07-12
func TestExpressionSpec(*testing.T) {
	println(`
		是通过运算符或者函数作用于操作数从而用于计算出一个值

		Operand     = Literal | OperandName [ TypeArgs ] | "(" Expression ")" .
		Literal     = BasicLit | CompositeLit | FunctionLit .
		BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
		OperandName = identifier | QualifiedIdent .


		操作数 operand
			Literal 字面量
		操作符 operator
			算数
			比较
			逻辑
			地址
			接受

		表达式分类
			基础表达式
			selector表达式
			method表达式
			index表达式
			切片表达式
			类型断言表达式
			函数调用表达式
			变长参数表达式
			泛型实例化表达式
			类型推断表达式
	`)
}
