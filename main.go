package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/longqimin/graphobj/parser"

	"github.com/antlr4-go/antlr/v4"
)

var (
	outputfile string
	verbose    bool
)

type GraphListener struct {
	*parser.BaseGraphObjListener

	path []string
	g    *graph
	d    Dot
}

func (l *GraphListener) EnterGraph(ctx *parser.GraphContext) {
	current_node := ctx.Name().GetSymbol().GetText()
	attr := ""
	if ctx.Attr() != nil {
		attr = ctx.Attr().GetSymbol().GetText()
	}

	// fmt.Println("Entering object:", current_node, attr)

	if len(l.path) == 0 {
		l.g = new(graph)
		l.g.name = current_node
	} else {
		l.g.addEdge(l.path[len(l.path)-1], current_node)
	}
	l.g.addNode(current_node, attr)
	l.path = append(l.path, current_node)
}

func (l *GraphListener) ExitGraph(ctx *parser.GraphContext) {
	// fmt.Println("ExitObject object:", ctx.Name().GetSymbol().GetText())

	l.path = l.path[0 : len(l.path)-1]
	if len(l.path) == 0 {
		l.d.addGraph(l.g)
		l.g = nil
	}
}

func main() {
	pflag.StringVarP(&outputfile, "output", "o", "", "output file")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s input [flags]:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "-o, --output string	output filename\n")
	}

	// 解析命令行标志
	pflag.Parse()

	// 获取位置参数
	positionArgs := pflag.Args()
	if len(positionArgs) == 0 || len(outputfile) == 0 {
		pflag.Usage()
		return
	}
	inputfile := positionArgs[0]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}

	// 创建输入流
	input := antlr.NewInputStream(string(data))

	// 创建词法分析器
	lexer := parser.NewGraphObjLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建语法分析器
	parser := parser.NewGraphObjParser(stream)

	// 开始解析
	tree := parser.Graphs()

	// 可以使用访问者模式遍历语法树进行处理
	listener := &GraphListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if verbose {
		fmt.Printf("parse GraphObject:\n%s", listener.d.String())
	}
	listener.d.render(outputfile, verbose)

	// fmt.Println(tree.ToStringTree(nil, parser))
}
