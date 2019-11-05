// 狗头语 版权 @2019 柴树杉。

package gotlang

import (
	"fmt"

	template "github.com/chai2010/gotlang/compiler"
)

func (p *GotApp) builtinFuncMap() template.FuncMap {
	return template.FuncMap{
		// 打印到标准输出
		"print":   p.fn_print,
		"printf":  p.fn_printf,
		"println": p.fn_println,

		// 打印到错误输出
		"eprint":   p.fn_eprint,
		"eprintf":  p.fn_eprintf,
		"eprintln": p.fn_eprintln,

		// 打印到字符串
		"sprint":   p.fn_sprint,
		"sprintf":  p.fn_sprintf,
		"sprintln": p.fn_sprintln,

		// 输入函数
		"read":       p.fn_read,
		"readrune":   p.fn_readrune,
		"readint":    p.fn_readint,
		"readstring": p.fn_readstring,
		"readline":   p.fn_readline,

		// 算术函数
		"add": p.fn_add,
		"sub": p.fn_sub,
		"mul": p.fn_mul,
		"div": p.fn_div,
		"mod": p.fn_mod,

		// 切片类型
		"mkslice": p.fn_mkslice,
		"append":  p.fn_append,

		// map类型
		"mkmap":  p.fn_mkmap,
		"mapset": p.fn_mapset,
		"mapdel": p.fn_mapdel,

		// xrange(stop)
		// xrange(start, end)
		// xrange(start, end, step)
		"xrange": p.fn_xrange,

		// 调用模板函数
		"template_call": p.fn_template_call,
		"template_ret":  p.fn_template_ret,
	}
}

func (p *GotApp) fn_print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(p.Stdout, a...)
}
func (p *GotApp) fn_printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(p.Stdout, format, a...)
}
func (p *GotApp) fn_println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(p.Stdout, a...)
}

func (p *GotApp) fn_eprint(a ...interface{}) (n int, err error) {
	return fmt.Fprint(p.Stderr, a...)
}
func (p *GotApp) fn_eprintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(p.Stderr, format, a...)
}
func (p *GotApp) fn_eprintln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(p.Stderr, a...)
}

func (p *GotApp) fn_sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}
func (p *GotApp) fn_sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
func (p *GotApp) fn_sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func (p *GotApp) fn_read(n int) []byte {
	data := make([]byte, n)
	n, _ = p.Stdin.Read(data)
	return data[:n]
}
func (p *GotApp) fn_readrune() int {
	r, _, _ := p.Stdin.ReadRune()
	return int(r)
}
func (p *GotApp) fn_readint() int {
	var i int
	fmt.Scanf("%d", &i)
	return i
}
func (p *GotApp) fn_readstring() string {
	var s string
	fmt.Scanf("%s", &s)
	return s
}
func (p *GotApp) fn_readline() string {
	line, _, _ := p.Stdin.ReadLine()
	return string(line)
}

func (p *GotApp) fn_add(a ...interface{}) int {
	var sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i].(int)
	}
	return sum
}
func (p *GotApp) fn_mul(a ...interface{}) int {
	var sum = 0
	for i := 0; i < len(a); i++ {
		sum *= a[i].(int)
	}
	return sum
}

func (p *GotApp) fn_sub(a ...interface{}) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0].(int)
	}

	var sum = a[0].(int)
	for i := 1; i < len(a); i++ {
		sum -= a[i].(int)
	}
	return sum
}

func (p *GotApp) fn_div(a ...interface{}) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0].(int)
	}

	var sum = a[0].(int)
	for i := 1; i < len(a); i++ {
		sum /= a[i].(int)
	}
	return sum
}

func (p *GotApp) fn_mod(a, b int) int {
	return a % b
}

func (p *GotApp) fn_mkslice(a ...interface{}) []interface{} {
	return a
}
func (p *GotApp) fn_append(a []interface{}, v ...interface{}) []interface{} {
	return append(a, v...)
}

func (p *GotApp) fn_mkmap() map[string]interface{} {
	return make(map[string]interface{})
}
func (p *GotApp) fn_mapset(m map[string]interface{}, k string, v interface{}) error {
	m[k] = v
	return nil
}
func (p *GotApp) fn_mapdel(m map[string]interface{}, k string) error {
	delete(m, k)
	return nil
}

func (p *GotApp) fn_xrange(stop interface{}, or_start_end_step ...interface{}) []int {
	if len(or_start_end_step) <= 0 {
		v := make([]int, stop.(int))
		for i := 0; i < len(v); i++ {
			v[i] = i
		}
		return v
	}

	var (
		start = stop.(int)
		end   = or_start_end_step[0].(int)
		step  = 1
	)

	if len(or_start_end_step) > 1 {
		step = or_start_end_step[1].(int)
	}

	if start == end || step <= 0 {
		return []int{}
	}

	if start < end {
		v := make([]int, (end-start)/step)
		for i, vi := 0, start; i < len(v); i, vi = i+1, vi+step {
			v[i] = vi
		}
		return v
	} else {
		v := make([]int, (start-end)/step)
		for i, vi := 0, start; i < len(v); i, vi = i+1, vi-step {
			v[i] = vi
		}
		return v
	}
}

func (p *GotApp) fn_template_call(name string, data interface{}) (ret interface{}, err error) {
	// 为返回值准备空间
	p.tmplRetStack = append(p.tmplRetStack, nil)

	// 从栈弹出最新的返回值
	defer func() {
		ret = p.tmplRetStack[len(p.tmplRetStack)-1]
		p.tmplRetStack = p.tmplRetStack[:len(p.tmplRetStack)-1]
	}()

	// 执行模板
	if err := p.tmpl.ExecuteTemplate(&p.Tout, name, data); err != nil {
		return nil, err
	}

	// 返回值在defer中处理
	return
}

func (p *GotApp) fn_template_ret(v interface{}) error {
	// 通一个 fn_template_call 内部多次调用
	// 将覆盖之前的值
	p.tmplRetStack[len(p.tmplRetStack)-1] = v
	return nil
}
