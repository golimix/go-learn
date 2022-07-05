package declaration

import (
	"fmt"
	"testing"
)

func TestSyntaxDeclarationSpec(t *testing.T) {
	fmt.Println(`
		总共就五种顶级要素,声明常量、声明变量、声明类型、声明函数、声明方法
		Declaration   = ConstDecl | TypeDecl | VarDecl .
		TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
	`)
}
