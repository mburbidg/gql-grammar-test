package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mburbidg/mygql/gql/runtime/status"
	"log"
)

type errorListener struct {
	errors []status.GQLError
}

func newErrorListener() *errorListener {
	return &errorListener{errors: []status.GQLError{}}
}

func (listener *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	listener.errors = append(listener.errors, status.NewSyntaxError(msg, line, column))
}

func (listener *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	log.Printf("ReportAmbiguity\n")
}

func (listener *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	log.Printf("ReportAttemptingFullContext\n")
}

func (listener *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	log.Printf("ReportContextSensitivity\n")
}

func (listener *errorListener) SemanticAnalysisError(line, column int, msg string) {
	listener.errors = append(listener.errors, status.NewSyntaxError(msg, line, column))
}
