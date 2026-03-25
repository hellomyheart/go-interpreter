package lexer

import (
	"testing"

	"go-interpreter/token"
)

// 测试词法分析器的NextToken方法
func TestNextToken(t *testing.T) {

	// 设置一个字符串，充当源代码
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	`

	// 定义一个测试用例的切片，包含了期望的Token类型和Token值
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// 设置期望的Token类型和Token值
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	// 构造一个词法分析器实例l, 把源代码赋值到词法分析器的input字段
	l := New(input)

	// 遍历预期结果tests
	for i, tt := range tests {
		// 获取Token
		tok := l.NextToken()
		// 预期结果和词法分析器获取到的结果比较
		if tok.Type != tt.expectedType {
			// 类型不相等，测试失败，输出错误信息
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		// 预期结果和词法分析器获取到值的结果比较
		if tok.Literal != tt.expectedLiteral {

			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
