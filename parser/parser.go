package parser

import (
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/mburbidg/grmtest/parser/gen"
)

func Parse(src string) error {
	input := antlr.NewInputStream(src)
	lexer := gen.NewGQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := gen.NewGQLParser(stream)
	errListener := newErrorListener()
	p.AddErrorListener(errListener)
	p.BuildParseTrees = true
	p.GqlProgram()
	if len(errListener.errors) > 0 {
		return errors.New(errListener.errors[0])
	}
	return nil
}
