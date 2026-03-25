package lexer

import "go-interpreter/token"

// 词法分析器结构体定义
type Lexer struct {
	input        string // 输入的字符串 也就是源代码
	position     int    // 所输入字符串中的当前位置（指向当前字符）
	readPosition int    //  所输入字符串中的当前读取位置（指向当前字符之后的一个字符）
	ch           byte   //当前正在检查的字符
}

// 构造函数
func New(input string) *Lexer {
	// 赋值输入程序
	l := &Lexer{input: input}
	// 读取第一个字符
	l.readChar()

	// 返回词法分析器实例
	return l
}

// 读词方法
func (l *Lexer) readChar() {
	// 判断是否已经读取到输入字符串的末尾
	// 读取的位置 >= 输入字符串的长度
	if l.readPosition >= len(l.input) {
		// 读取到末尾了，设置当前字符为0
		l.ch = 0
	} else {
		// 还没有读取到末尾，设置当前字符为输入字符串中读取位置的字符
		l.ch = l.input[l.readPosition]
	}
	// 把当前位置更新到读取位置
	l.position = l.readPosition
	// 把读取位置向前移动一位
	l.readPosition += 1
}

// 获取下一个Token的方法
func (l *Lexer) NextToken() token.Token {
	// 初始化Token
	var tok token.Token

	// 根据l.ch 做判断
	// 在词法分析器初始化的时候 需要调用一下readChar方法 读取第一个字符
	// new 构造方法里面调用了readChar方法 读取了第一个字符

	// 根据ch 字符的不同创建不同类型的词（类型、值）
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// 这里比较特殊，如果值是0, 把token的类型设置为EOF, 不设置值，根据GO语法设置零值，应该是空字符串
		tok.Literal = ""
		tok.Type = token.EOF
	}
	// 取到token之后，调用readChar方法 读取下一个字符
	l.readChar()
	// 返回tgoken
	return tok
}

// Token的构造函数， 设置token的类型和token的值
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

