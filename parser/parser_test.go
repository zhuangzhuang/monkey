package parser

import (
	"monkey/ast"
	"testing"
)

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

// func testLiteralExpression(
// 	t *testing.T,
// 	exp ast.Expression,
// 	expected interface{},
// ) bool {
// 	switch v := expected.(type) {
// 	case int:
// 		return testIntegerLiteral(t, exp, int64(v))
// 	case int64:
// 		return testIntegerLiteral(t, exp, v)
// 	case string:
// 		return testIdentifier(t, exp, v)
// 	}
// 	t.Errorf("type of exp not handled, got=%T", exp)
// 	return false
// }
