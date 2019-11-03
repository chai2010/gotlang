// 狗头语 版权 @2019 柴树杉。

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chai2010/gotlang"
)

const usage = `
Got is a tool for run got(Go template language) program.
Usage: got progran.gotmpl [args...]
       got -h
`

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}

	prog := loadProgram(os.Args[1])
	args := os.Args[2:]
	gotApp := gotlang.NewGotApp(prog, args...)

	if err := gotApp.Run(); err != nil {
		log.Fatal(err)
	}
}

func loadProgram(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
