package parser

import (
	"fmt"
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"go-interpreter/token"
)

// 分析程序 结构体
type Parser struct {
	// 指向词法分析器实例的指针
	l *lexer.Lexer

	// 当前词法单元
	curToken token.Token

	// 下一个词法单元
	peekToken token.Token

	// 解析错误信息
	errors []string
}

// 构造函数
func New(l *lexer.Lexer) *Parser {
	// 创建实例p
	p := &Parser{l: l,
		errors: []string{},
	}

	// 读取两个词法单元 以设置curToken和peekToken
	p.nextToken()

	p.nextToken()

	return p
}

// 获取解析错误信息
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to tb %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// 指针后移
func (p *Parser) nextToken() {
	// 当前指针改成下一个指针
	p.curToken = p.peekToken
	// 下一个指针 改成 l的后移一个指针
	// 初始化的时候调用两次nextToken
	// 这两个指针都指向第一第二个位置了
	p.peekToken = p.l.NextToken()
}

// 解析程序
func (p *Parser) ParseProgram() *ast.Program {
	// 构建一个根节点
	program := &ast.Program{}
	// 初始化语句切片
	program.Statements = []ast.Statement{}

	// 循环解析程序 直到文件结束
	for !p.curTokenIs(token.EOF) {
		// 解析当前语句的语法
		stmt := p.parseStatement()
		// 如果语法不为nil
		if stmt != nil {
			// 添加语法到切片
			program.Statements = append(program.Statements, stmt)
		}
		// 指针后移
		p.nextToken()
	}
	// 返回结果
	return program
}

// 解析语法
func (p *Parser) parseStatement() ast.Statement {
	// 根据当前类型做不同的处理
	switch p.curToken.Type {
	case token.LET:
		// 处理let类型语法
		return p.parseLetStatement()
	default:
		// 其他类型暂不处理
		return nil
	}
}

// let类型语法处理
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// 创建let语法实体
	stmt := &ast.LetStatement{Token: p.curToken}

	// 判断下一个词法单元如果不是 IDENT类型 就返回nil
	// 是语法错误了
	// 这里做了返回nil处理
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// 创建了一个 表达式类型 语法实体
	// token 设置为当前词法单元（在expectPeek 判断里面，已经后移了一步， curToken已经 是x了）
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 判断下一个词法单元如果不是 ASSIGN类型 就返回nil
	// 这里又右移了一步
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// 此时当前的 语法单元是 =
	// TODO: 跳过对表达式的处理，直到遇见分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// 判断当前词法单元是否是期望的类型
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// 判断下一个词法单元是否是期望的类型
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// 如果下一个词法单元是期望的类型 就指针后移并返回true 否则返回false
func (p *Parser) expectPeek(t token.TokenType) bool {
	// 如果下一个词法单元的类型是 t
	if p.peekTokenIs(t) {
		// 双指针都后移一步
		p.nextToken()
		// 返回true
		return true
	} else {
		// 返回false
		p.peekError(t)
		return false
	}
}
