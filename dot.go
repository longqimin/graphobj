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
	content := ""
	for _, e := range g.edges {
		content += e.String()
	}
	return content
}

type Dot struct {
	graphs []*graph
}

func (d *Dot) addGraph(g *graph) {
	d.graphs = append(d.graphs, g)
}

func (d *Dot) String() string {
	content := ""
	for _, g := range d.graphs {
		content += g.String()
	}
	return content
}

func (d *Dot) render(theme, layout, filename string, verbose bool) {
	// 生成 D2 文件
	d2File := filename + ".d2"
	if verbose {
		fmt.Printf("generate .d2: %s\n", d2File)
	}
	if err := os.WriteFile(d2File, []byte(d.String()), 0644); err != nil {
		panic(err)
	}

	// D2渲染图像
	if verbose {
		fmt.Printf("exec: d2 --theme %s --layout %s %s %s\n", theme, layout, d2File, filename)
	}
	cmd := exec.Command("d2", "--theme", theme, "--layout", layout, d2File, filename)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
