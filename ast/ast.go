package ast

import "go-interpreter/token"

// Node接口
type Node interface {
	// 返回结果对应的Token字面量 （字符串）
	TokenLiteral() string
}

// 语句接口
type Statement interface {
	// 嵌入了Node接口
	Node
	// 语句方法 标记方法，无返回值
	statementNode()
}

// 表达式接口
type Expression interface {
	// 嵌入了node接口
	Node
	// 表达式方法 标记方法，无返回值
	expressionNode()
}

// 语句结构体
type Program struct {
	Statements []Statement
}

// 返回第一个语句的 Token， 没有则返回空字符串
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Let 语法单元
// let 的 标识符其实是不会产生值的， 但是（Identifier）依然使用表达式类型，
// 是为了让代码保持简单， 所有的表达式语法节点，都可以使用这一个
type LetStatement struct {
	Token token.Token // token.LET词法单元
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

// 标记方法
func (ls *LetStatement) statementNode() {}

// 返回词的值
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// 表达式类型的结构体
type Identifier struct {
	Token token.Token // token.IDENT 词法单元
	Value string      // 值
}

// ┌─────────────────────────────────┐
// │      Identifier (结构体)         │
// │  ┌───────────────────────────┐  │
// │  │  Token: {Type: IDENT,     │  │
// │  │         Literal: "x"}     │  │
// │  ├───────────────────────────┤  │
// │  │  Value: "x"               │  │
// │  └───────────────────────────┘  │
// └─────────────────────────────────┘

// 标记接口
func (i *Identifier) expressionNode() {}

// 返回标识符的字符串值
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
