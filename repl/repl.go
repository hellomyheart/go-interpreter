package repl

import (
	"bufio"
	"fmt"
	"go-interpreter/lexer"
	"go-interpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// 创建一个 bufio.Scanner 对象 scanner，用于从输入流 in 中高效地读取文本行。
	scanner := bufio.NewScanner(in)

	// 开始一个无限循环，确保持续监听用户输入，直到程序退出。
	for {
		// 向输出流 out 打印提示符 [PROMPT]
		fmt.Fprint(out, PROMPT)
		// 调用 scanner.Scan 方法读取下一行输入。
		scanned := scanner.Scan()
		// 返回布尔值 scanned，表示是否成功读取到数据。
		if !scanned {
			// 检查 scanned 是否为 false，通常意味着遇到了输入结束（EOF）或发生错误
			// 如果读取失败，直接退出
			return
		}

		// 获取 scanner 当前读取到的完整文本行，赋值给变量 line。
		line := scanner.Text()
		// 调用 lexer.New构造函数，传入 line 创建一个新的词法分析器实例 l。
		l := lexer.New(line)
		// 开始内部循环，不断调用 l.NextToken 获取下一个令牌（Token）。
		// 初始化 tok 为第一个令牌。
		// 循环条件是当前令牌的类型 tok.Type 不等于 [token.EOF]
		// 每次迭代后更新 tok 为下一个令牌。
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// 将当前令牌 tok 的详细结构体内容格式化打印到输出流 out。
			// %+v 格式符会显示结构体的字段名和值。
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}
}
