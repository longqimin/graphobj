// Code generated from GraphObj.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GraphObjLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GraphObjLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func graphobjlexerLexerInit() {
	staticData := &GraphObjLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "Name", "WS", "Attr",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "Name", "WS", "Attr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 37, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 4, 2, 17, 8, 2, 11, 2, 12, 2, 18,
		1, 3, 5, 3, 22, 8, 3, 10, 3, 12, 3, 25, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 5,
		4, 31, 8, 4, 10, 4, 12, 4, 34, 9, 4, 1, 4, 1, 4, 0, 0, 5, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 1, 0, 2, 6, 0, 9, 10, 13, 13, 32, 32, 44, 44, 123, 123,
		125, 125, 4, 0, 9, 10, 13, 13, 32, 32, 44, 44, 39, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1,
		11, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 16, 1, 0, 0, 0, 7, 23, 1, 0, 0, 0,
		9, 28, 1, 0, 0, 0, 11, 12, 5, 123, 0, 0, 12, 2, 1, 0, 0, 0, 13, 14, 5,
		125, 0, 0, 14, 4, 1, 0, 0, 0, 15, 17, 8, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17,
		18, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 6, 1, 0, 0,
		0, 20, 22, 7, 1, 0, 0, 21, 20, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21,
		1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 26, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0,
		26, 27, 6, 3, 0, 0, 27, 8, 1, 0, 0, 0, 28, 32, 5, 91, 0, 0, 29, 31, 9,
		0, 0, 0, 30, 29, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32,
		33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 36, 5, 93,
		0, 0, 36, 10, 1, 0, 0, 0, 4, 0, 18, 23, 32, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GraphObjLexerInit initializes any static state used to implement GraphObjLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGraphObjLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GraphObjLexerInit() {
	staticData := &GraphObjLexerLexerStaticData
	staticData.once.Do(graphobjlexerLexerInit)
}

// NewGraphObjLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGraphObjLexer(input antlr.CharStream) *GraphObjLexer {
	GraphObjLexerInit()
	l := new(GraphObjLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GraphObjLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GraphObj.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GraphObjLexer tokens.
const (
	GraphObjLexerT__0 = 1
	GraphObjLexerT__1 = 2
	GraphObjLexerName = 3
	GraphObjLexerWS   = 4
	GraphObjLexerAttr = 5
)
