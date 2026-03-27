package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	// 获取当前操作系统的用户信息
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the go-interpreter programming language!\n", user.Username)
	fmt.Print("Feel free to type in commands\n")
	// 调用 repl.Start 函数启动解释器的交互循环（Read-Eval-Print Loop）
	// 传入 os.Stdin 作为标准输入流，os.Stdout 作为标准输出流。
	repl.Start(os.Stdin, os.Stdout)
}
