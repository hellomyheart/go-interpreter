package parser

import (
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {

	// 测试语句
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;`

	// 创建词法分析器
	l := lexer.New(input)
	// 使用词法分析器创建语法分析器
	p := New(l)

	// 调用语法分析方法
	program := p.ParseProgram()

	// 检查语法异常
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	// Statements 长度是3
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	// 测试结果结构体
	tests := []struct {
		expectedIdentifier string
	}{
		// 期望测试结果
		{"x"},
		{"y"},
		{"foobar"},
	}

	// 遍历预期结果
	for i, tt := range tests {
		// 遍历语法分析结果
		stmt := program.Statements[i]
		// 测试逻辑 只测试let
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// 检查语法一次
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	// 语法异常直接失败
	// 立即标记测试为失败，并停止执行测试
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// 判断字符串是否是let
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q, s.TokenLiteral()", s.TokenLiteral())
		return false
	}

	// 类型转换，把接口类型转换为 LetStatement
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	// 标识符判断
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	// 标识符的字符串值
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
