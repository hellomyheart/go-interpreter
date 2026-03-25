package token

// 把所有的词都以常量的形式声明
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

// 词的类型
type TokenType string

// 词结构体
// 当前是包含的词的类型 和词的值
type Token struct {
	Type    TokenType
	Literal string
}

// 定义关键字
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// 从map中获取，获取到则返回关键字
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// 找不到则返回标识符
	return IDENT
}
