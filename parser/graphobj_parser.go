// Code generated from GraphObj.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GraphObj

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GraphObjParser struct {
	*antlr.BaseParser
}

var GraphObjParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func graphobjParserInit() {
	staticData := &GraphObjParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "Name", "WS", "Attr",
	}
	staticData.RuleNames = []string{
		"graph", "graphs",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 29, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 3, 0, 6, 8, 0, 1, 0, 1, 0, 3,
		0, 10, 8, 0, 1, 0, 3, 0, 13, 8, 0, 1, 0, 1, 0, 5, 0, 17, 8, 0, 10, 0, 12,
		0, 20, 9, 0, 1, 0, 1, 0, 1, 1, 4, 1, 25, 8, 1, 11, 1, 12, 1, 26, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 31, 0, 5, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 6, 5,
		4, 0, 0, 5, 4, 1, 0, 0, 0, 5, 6, 1, 0, 0, 0, 6, 7, 1, 0, 0, 0, 7, 9, 5,
		3, 0, 0, 8, 10, 5, 5, 0, 0, 9, 8, 1, 0, 0, 0, 9, 10, 1, 0, 0, 0, 10, 12,
		1, 0, 0, 0, 11, 13, 5, 4, 0, 0, 12, 11, 1, 0, 0, 0, 12, 13, 1, 0, 0, 0,
		13, 14, 1, 0, 0, 0, 14, 18, 5, 1, 0, 0, 15, 17, 3, 0, 0, 0, 16, 15, 1,
		0, 0, 0, 17, 20, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19,
		21, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 21, 22, 5, 2, 0, 0, 22, 1, 1, 0, 0,
		0, 23, 25, 3, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24,
		1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 3, 1, 0, 0, 0, 5, 5, 9, 12, 18, 26,
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

// GraphObjParserInit initializes any static state used to implement GraphObjParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGraphObjParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GraphObjParserInit() {
	staticData := &GraphObjParserStaticData
	staticData.once.Do(graphobjParserInit)
}

// NewGraphObjParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGraphObjParser(input antlr.TokenStream) *GraphObjParser {
	GraphObjParserInit()
	this := new(GraphObjParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GraphObjParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GraphObj.g4"

	return this
}

// GraphObjParser tokens.
const (
	GraphObjParserEOF  = antlr.TokenEOF
	GraphObjParserT__0 = 1
	GraphObjParserT__1 = 2
	GraphObjParserName = 3
	GraphObjParserWS   = 4
	GraphObjParserAttr = 5
)

// GraphObjParser rules.
const (
	GraphObjParserRULE_graph  = 0
	GraphObjParserRULE_graphs = 1
)

// IGraphContext is an interface to support dynamic dispatch.
type IGraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	Attr() antlr.TerminalNode
	AllGraph() []IGraphContext
	Graph(i int) IGraphContext

	// IsGraphContext differentiates from other interfaces.
	IsGraphContext()
}

type GraphContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphContext() *GraphContext {
	var p = new(GraphContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphObjParserRULE_graph
	return p
}

func InitEmptyGraphContext(p *GraphContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphObjParserRULE_graph
}

func (*GraphContext) IsGraphContext() {}

func NewGraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphContext {
	var p = new(GraphContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphObjParserRULE_graph

	return p
}

func (s *GraphContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphContext) Name() antlr.TerminalNode {
	return s.GetToken(GraphObjParserName, 0)
}

func (s *GraphContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GraphObjParserWS)
}

func (s *GraphContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GraphObjParserWS, i)
}

func (s *GraphContext) Attr() antlr.TerminalNode {
	return s.GetToken(GraphObjParserAttr, 0)
}

func (s *GraphContext) AllGraph() []IGraphContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGraphContext); ok {
			len++
		}
	}

	tst := make([]IGraphContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGraphContext); ok {
			tst[i] = t.(IGraphContext)
			i++
		}
	}

	return tst
}

func (s *GraphContext) Graph(i int) IGraphContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphContext)
}

func (s *GraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphObjListener); ok {
		listenerT.EnterGraph(s)
	}
}

func (s *GraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphObjListener); ok {
		listenerT.ExitGraph(s)
	}
}

func (p *GraphObjParser) Graph() (localctx IGraphContext) {
	localctx = NewGraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphObjParserRULE_graph)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(5)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphObjParserWS {
		{
			p.SetState(4)
			p.Match(GraphObjParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(7)
		p.Match(GraphObjParserName)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphObjParserAttr {
		{
			p.SetState(8)
			p.Match(GraphObjParserAttr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphObjParserWS {
		{
			p.SetState(11)
			p.Match(GraphObjParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(14)
		p.Match(GraphObjParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphObjParserName || _la == GraphObjParserWS {
		{
			p.SetState(15)
			p.Graph()
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(GraphObjParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphsContext is an interface to support dynamic dispatch.
type IGraphsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGraph() []IGraphContext
	Graph(i int) IGraphContext

	// IsGraphsContext differentiates from other interfaces.
	IsGraphsContext()
}

type GraphsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphsContext() *GraphsContext {
	var p = new(GraphsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphObjParserRULE_graphs
	return p
}

func InitEmptyGraphsContext(p *GraphsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphObjParserRULE_graphs
}

func (*GraphsContext) IsGraphsContext() {}

func NewGraphsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphsContext {
	var p = new(GraphsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphObjParserRULE_graphs

	return p
}

func (s *GraphsContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphsContext) AllGraph() []IGraphContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGraphContext); ok {
			len++
		}
	}

	tst := make([]IGraphContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGraphContext); ok {
			tst[i] = t.(IGraphContext)
			i++
		}
	}

	return tst
}

func (s *GraphsContext) Graph(i int) IGraphContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphContext)
}

func (s *GraphsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphObjListener); ok {
		listenerT.EnterGraphs(s)
	}
}

func (s *GraphsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphObjListener); ok {
		listenerT.ExitGraphs(s)
	}
}

func (p *GraphObjParser) Graphs() (localctx IGraphsContext) {
	localctx = NewGraphsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphObjParserRULE_graphs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphObjParserName || _la == GraphObjParserWS {
		{
			p.SetState(23)
			p.Graph()
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
