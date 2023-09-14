package main

import (
	"fmt"
	"os"

	"github.com/mdshack/brackets/pkg/parse/lexer"
	"github.com/mdshack/brackets/pkg/parse/token"
)

func main() {
	f, err := os.Open("./example.brk")
	if err != nil {
		panic(err)
	}

	l := lexer.NewFromReader(f)

	for {
		p, t, v := l.Lex()
		fmt.Println(p, token.Token(t), v)
		if t == token.EOF {
			break
		}
	}
}
