package gograph

type ID interface{}

type DirectedGraph interface {
	GetNodes() map[ID]Node
	GetNode(id ID) Node

	GetArcs() map[ID]map[ID]Arc
	GetArc(source ID, target ID) Arc
}

type Node interface {
	GraphID() ID
}

type Arc interface {
	GraphSourceID() ID
	GraphTargetID() ID
	Weight() int
}

type BuiltinGraph struct {
	nodes map[ID]Node
	arcs  map[ID]map[ID]Arc
}

func NewDirectedGraph() *BuiltinGraph {
	return &BuiltinGraph{
		nodes: map[ID]Node{},
		arcs:  map[ID]map[ID]Arc{},
	}
}

func (g *BuiltinGraph) GetNodes() map[ID]Node {
	return g.nodes
}

func (g *BuiltinGraph) GetNode(id ID) Node {
	return g.nodes[id]
}

func (g *BuiltinGraph) GetArcs() map[ID]map[ID]Arc {
	return g.arcs
}

func (g *BuiltinGraph) GetArc(source ID, target ID) Arc {
	m := g.arcs[source]
	if m != nil {
		return m[target]
	}
	return nil
}

func (g *BuiltinGraph) AddNode(node Node) {
	g.nodes[node.GraphID()] = node
}

func (g *BuiltinGraph) AddArc(arc Arc) {
	source := arc.GraphSourceID()
	target := arc.GraphTargetID()

	m := g.arcs[source]
	if m == nil {
		m = map[ID]Arc{}
		g.arcs[source] = m
	}

	m[target] = arc
}

type BuiltinNode struct {
	id ID
}

func NewNode(id ID) *BuiltinNode {
	return &BuiltinNode{
		id: id,
	}
}

func (v *BuiltinNode) GraphID() ID {
	return v.id
}

type BuiltinArc struct {
	source ID
	target ID
	weight int
}

func NewArc(source ID, target ID, weight int) *BuiltinArc {
	return &BuiltinArc{
		source: source,
		target: target,
		weight: weight,
	}
}

func (e *BuiltinArc) GraphSourceID() ID {
	return e.source
}

func (e *BuiltinArc) GraphTargetID() ID {
	return e.target
}

func (e *BuiltinArc) Weight() int {
	return e.weight
}
