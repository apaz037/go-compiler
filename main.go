package main

import (
	"github.com/apaz037/go-compiler/lexer"
	"github.com/apaz037/go-compiler/parser"
)

func must (err error) {
	if err != nil {
		panic(err)
	}
}

func Parse(input string) {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()

}