package main

import (
	"fmt"
	"os"
	"os/exec"
)

type edge struct {
	left  string
	right string
}

func (g *edge) String() string {
	return fmt.Sprintf("\"%s\" -> \"%s\";\n", g.left, g.right)
}

type node struct {
	name string
	attr string
}

func (g *node) String() string {
	return fmt.Sprintf("\"%s\" %s;\n", g.name, g.attr)
}

type graph struct {
	name  string
	nodes []node
	edges []edge
}

func (g *graph) addEdge(left, right string) {
	g.edges = append(g.edges, edge{left, right})
}
func (g *graph) addNode(name, attr string) {
	g.nodes = append(g.nodes, node{name, attr})
}
func (g *graph) String() string {
	content := "subgraph " + g.name + "{\n"
	for _, n := range g.nodes {
		content += n.String()
	}
	for _, e := range g.edges {
		content += e.String()
	}
	content += "}\n"
	return content
}

type Dot struct {
	graphs []*graph
}

func (d *Dot) addGraph(g *graph) {
	d.graphs = append(d.graphs, g)
}

func (d *Dot) String() string {
	content := "digraph {\n"
	for _, g := range d.graphs {
		content += g.String()
	}
	content += "}\n"
	return content
}

func (d *Dot) render(filename string, verbose bool) {
	// 生成 DOT 文件
	dotfile := filename + ".dot"
	if verbose {
		fmt.Printf("generate .dot: %s\n", dotfile)
	}
	if err := os.WriteFile(dotfile, []byte(d.String()), 0644); err != nil {
		panic(err)
	}

	// 调用 Graphviz 渲染图像
	if verbose {
		fmt.Printf("exec: dot -Tpng %s -o %s\n", dotfile, filename)
	}
	cmd := exec.Command("dot", "-Tpng", dotfile, "-o", filename)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
