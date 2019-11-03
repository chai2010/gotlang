// 狗头语 版权 @2019 柴树杉。

package gotlang

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"text/template"
)

// 狗头语程序对象
type GotApp struct {
	Program string   // 脚本程序
	Args    []string // 命令行参数
	Dir     string   // 工作目录

	Tout   bytes.Buffer  // 模板输出
	Stdin  *bufio.Reader // 标准输入
	Stdout io.Writer     // 标准输出
	Stderr io.Writer     // 错误输出

	LeftDelimiter  string // 左分隔符，默认“{{”
	RightDelimiter string // 右分隔符，默认“}}”

	tmpl         *template.Template // 模板引擎
	tmplRetStack []interface{}      // 模板的返回值栈
}

func NewGotApp(prog string, args ...string) *GotApp {
	wd, _ := os.Getwd()
	return &GotApp{
		Program: prog,
		Args:    append([]string{}, args...),
		Dir:     wd,

		Stdin:  bufio.NewReader(os.Stdin),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (p *GotApp) Run() error {
	if p.tmpl == nil {
		if err := p.buildTemplate(); err != nil {
			return err
		}
	}

	// 执行程序
	p.Tout.Reset()
	if err := p.tmpl.Execute(&p.Tout, p.Args); err != nil {
		return err
	}

	return nil
}

func (p *GotApp) buildTemplate() error {
	if p.tmpl != nil {
		return nil
	}

	t := template.New("")

	// 指定分隔符
	left, right := "{{", "}}"
	if p.LeftDelimiter != "" {
		left = p.LeftDelimiter
	}
	if p.RightDelimiter != "" {
		right = p.RightDelimiter
	}
	t = t.Delims(left, right)

	// 注册内置到模板函数
	t = t.Funcs(p.builtinFuncMap())

	// 解析程序
	t, err := t.Parse(p.Program)
	if err != nil {
		return err
	}

	// OK
	p.tmpl = t
	return nil
}
