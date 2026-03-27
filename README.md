这个是用GO语言制作解释器的代码

1.初始化项目

```shell
go mod init go-interpreter
```

# 1.词法分析

## 1.1 词法分析

解释器：

先使用词法分析器将源代码转换为词法单元，

再通过语法分析器将词法单元转换成抽象语法树。

举例：

代码：

```go
"let x = 5 + 5;"
```

转换为词法单元是：

```shell
[
LET,
IDENTIFIER("x"),
EQUAL_SIGN,
INTEGER(5),
PLUS_SIGN,
INTEGER(5),
SEMICOLON
]
```

具有完整功能的词法分析器还可以将行号、列号和文件名附加到词法
单元中。这么做是为了在后面的语法分析阶段输出更有用的报错消息。

## 1.2 定义词法单元

示例代码

```shell
let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
```

源代码中的词法 （关键字、操作符、标识符、标点符号）

定义Token数据结构，用来保存词法基础信息

## 1.3 词法分析器

编写完代码后执行测试

```shell
go test .\lexer\
```

当前完成了最基础的词法分析器， 当然支持的特性有限。

## 1.4 扩展词法分析器


基本完成基础的词法分析
