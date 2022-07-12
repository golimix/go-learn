package declaration

import "testing"

func TestSyntaxDeclarationMethodSpec(t *testing.T) {
	println(`
		方法声明相对于函数声明多了一个 Receiver, 建议使用指针类型的 Receiver
		MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
		Receiver   = Parameters .
		Parameters     = "(" [ ParameterList [ "," ] ] ")" .
		ParameterList  = ParameterDecl { "," ParameterDecl } .
		ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
	`)
}
