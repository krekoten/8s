package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// func init() {
// var code string
// l := lexer.New(code)
// tokens := l.Tokenize()

// p := parser.New(tokens)
// ast := p.Parse()

// c := compiler.New(ast)
// byteCode := c.Compile()

// v := vm.New(byteCode)
// v.Run()
// }
var rootCmd = &cobra.Command{
	Use:   "8s",
	Short: "8s is a brainfuck compiler & VM",
	Long: `A very overengineered and badly implemented
                compiler and VM for brainfuck in Go.
                It's just a small project to learn Go`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
