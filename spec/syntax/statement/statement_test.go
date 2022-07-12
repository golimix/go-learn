// 语句
package statement

import "testing"

// TODO 规范-语法-语句-完善内容
func TestSyntaxStatementSpec(t *testing.T) {
	println(`
		Statement =
			Declaration | LabeledStmt | SimpleStmt |
			GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
			FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
			DeferStmt .

		SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
	`)
}
