package token

const (
	// 非法的
	ILLEGAL = "ILLEGAL"
	// 文件结束
	EOF = "EOF"

	// ILLEGAL表示 遇到未知的词法单元或字符，EOF则表示文件结尾（End Of File），用于通知后续章节会介绍的语法分析器停机。

	// 标识符 + 字面量
	IDENT = "IDENT" // add, foobar, x, y ...
	INT   = "INT"   // 1343456

	// 运算符
	ASSIGN = "="
	PLUS   = "+"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
