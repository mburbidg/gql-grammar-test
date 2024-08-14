package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mburbidg/mygql/gql/parser/ast"
	"github.com/mburbidg/mygql/gql/parser/gen"
)

func Parse(src string) (*ast.GQLProgram, error) {
	input := antlr.NewInputStream(src)
	lexer := gen.NewGQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := gen.NewGQLParser(stream)
	errListener := newErrorListener()
	p.AddErrorListener(errListener)
	p.BuildParseTrees = true
	tree := p.GqlProgram()
	listener := NewListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	if len(errListener.errors) > 0 {
		return nil, errListener.errors[0]
	}
	return listener.GQLProgram, nil
}
