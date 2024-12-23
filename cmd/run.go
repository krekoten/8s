package cmd

import (
	"os"

	"github.com/krekoten/8s/compiler"
	"github.com/krekoten/8s/lexer"
	"github.com/krekoten/8s/parser"
	"github.com/krekoten/8s/vm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "run [filename]",
		Short: "Run brainfuck code in file",
		Args:  cobra.MinimumNArgs(1),
		RunE:  cmdRunHandler,
	})
}

func cmdRunHandler(cmd *cobra.Command, args []string) error {
	fileName := args[0]
	code, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	runCode(string(code))

	return nil
}

func runCode(code string) {
	l := lexer.New(code)
	tokens := l.Tokenize()

	p := parser.New(tokens)
	ast := p.Parse()

	c := compiler.New(ast)
	byteCode := c.Compile()

	v := vm.New(byteCode)
	v.Run()
}
