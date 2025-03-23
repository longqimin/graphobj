// Code generated from GraphObj.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GraphObj

import "github.com/antlr4-go/antlr/v4"

// BaseGraphObjListener is a complete listener for a parse tree produced by GraphObjParser.
type BaseGraphObjListener struct{}

var _ GraphObjListener = &BaseGraphObjListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphObjListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphObjListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphObjListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphObjListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGraph is called when production graph is entered.
func (s *BaseGraphObjListener) EnterGraph(ctx *GraphContext) {}

// ExitGraph is called when production graph is exited.
func (s *BaseGraphObjListener) ExitGraph(ctx *GraphContext) {}

// EnterGraphs is called when production graphs is entered.
func (s *BaseGraphObjListener) EnterGraphs(ctx *GraphsContext) {}

// ExitGraphs is called when production graphs is exited.
func (s *BaseGraphObjListener) ExitGraphs(ctx *GraphsContext) {}
