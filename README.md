# 狗头语(Go Template Language)

![](gotlang-logo.png)

狗头语是“柴树杉(chai2010)”于2019年11月基于Go语言标准库`text/template`包定制的图灵完备的语言。

## 安装

```
$ go get github.com/chai2010/gotlang/got
```

## 例子1：输出命令行参数

创建`hello.gotmpl`程序，输出命令行参数：

```gotmpl
{{/* 狗头语 版权 @2019 柴树杉 */}}}

{{template "main" .}}

{{define "main"}}
	{{range $i, $v := . }}
		{{println $i $v}}
	{{end}}
{{end}}
```

运行脚本：

```
$ got hello.gotmpl aa bb cc
0 aa
1 bb
2 cc
```

## 例子2：打印斐波那契数列

创建`./examples/fib.gotmpl`程序，输出命令行参数：

```
{{/* 狗头语 版权 @2019 柴树杉 */}}}

{{template_call "main" .}}

{{define "main"}}
	{{range $_, $i := (xrange 10) }}
		{{println $i ":" (template_call "fib" $i)}}
	{{end}}
{{end}}}

{{define "fib"}}
	{{if (le . 0)}}
		{{template_ret 0}}
		{{/*return fib(0)*/}}
	{{else if (eq . 1)}}
		{{template_ret 1}}
		{{/*return fib(1)*/}}
	{{else}}
		{{template_ret (add (template_call "fib" (sub . 1)) (template_call "fib" (sub . 2)))}}
		{{/*return fib(n-1) + fib(n-2)*/}}
	{{end}}
{{end}}
```

运行脚本：

```
$ got ./examples/fib.gotmpl
0 : 0
1 : 1
2 : 1
3 : 2
4 : 3
5 : 5
6 : 8
7 : 13
8 : 21
9 : 34
```

## 版权

狗头语 版权 @2019 柴树杉。
