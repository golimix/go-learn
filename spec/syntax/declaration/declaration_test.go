package declaration

import (
	"fmt"
	"testing"
)

func TestSyntaxDeclarationSpec(t *testing.T) {
	fmt.Println(`
		总共就五种顶级要素声明
		声明: 常量、变量、类型、函数、方法绑定标识符的过程
		声明常量
		声明变量
		声明类型
		声明函数
		声明方法
		Declaration   = ConstDecl | TypeDecl | VarDecl .
		TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
	`)
	TestSyntaxDeclarationTypeSpec(t)
	TestSyntaxDeclarationConstSpec(t)
	TestSyntaxDeclarationVarSpec(t)
	TestSyntaxDeclarationFunctionSpec(t)
	TestSyntaxDeclarationMethodSpec(t)
}
