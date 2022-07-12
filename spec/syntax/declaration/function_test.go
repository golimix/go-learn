package declaration

import "testing"

func TestSyntaxDeclarationFunctionSpec(t *testing.T) {
	println(`
		FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
		FunctionName = identifier .
		FunctionBody = Block .
	`)
}
