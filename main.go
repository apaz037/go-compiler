package main

import (
	"github.com/apaz037/go-compiler/lexer"
	"github.com/apaz037/go-compiler/parser"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		panic("no valid file name or path provided for file!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	contents, err := ioutil.ReadFile(absPath)
	check(err)
	Parse(string(contents))
}


func check (err error) {
	if err != nil {
		panic(err)
	}
}

func Parse(input string) {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()

	_, err := p.Parse(l)
	check(err)
}
