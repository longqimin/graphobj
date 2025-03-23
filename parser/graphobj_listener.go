// Code generated from GraphObj.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GraphObj

import "github.com/antlr4-go/antlr/v4"

// GraphObjListener is a complete listener for a parse tree produced by GraphObjParser.
type GraphObjListener interface {
	antlr.ParseTreeListener

	// EnterGraph is called when entering the graph production.
	EnterGraph(c *GraphContext)

	// EnterGraphs is called when entering the graphs production.
	EnterGraphs(c *GraphsContext)

	// ExitGraph is called when exiting the graph production.
	ExitGraph(c *GraphContext)

	// ExitGraphs is called when exiting the graphs production.
	ExitGraphs(c *GraphsContext)
}
