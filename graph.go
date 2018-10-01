package gograph

type ID interface{}

type Graph struct {
	vertices map[ID]Vertex
	arcs     map[ID]map[ID]Arc
}

func NewGraph() *Graph {
	return &Graph{
		vertices: map[ID]Vertex{},
		arcs:     map[ID]map[ID]Arc{},
	}
}

func (g *Graph) AddVertex(vertex Vertex) {
	g.vertices[vertex.GraphID()] = vertex
}

func (g *Graph) GetVertex(id ID) Vertex {
	return g.vertices[id]
}

func (g *Graph) AddArc(arc Arc) {
	source := arc.GraphSourceID()
	target := arc.GraphTargetID()

	m := g.arcs[source]
	if m == nil {
		m = map[ID]Arc{}
		g.arcs[source] = m
	}

	m[target] = arc
}

func (g *Graph) GetArc(source ID, target ID) Arc {
	m := g.arcs[source]
	if m != nil {
		return m[target]
	}
	return nil
}

type Vertex interface {
	GraphID() ID
}

type simpleVertex struct {
	id ID
}

func NewSimpleVertex(id ID) Vertex {
	return &simpleVertex{
		id: id,
	}
}

func (v *simpleVertex) GraphID() ID {
	return v.id
}

type Arc interface {
	GraphSourceID() ID
	GraphTargetID() ID
	Weight() int
}

type simpleArc struct {
	source ID
	target ID
	weight int
}

func NewSimpleArc(source ID, target ID, weight int) Arc {
	return &simpleArc{
		source: source,
		target: target,
		weight: weight,
	}
}

func (e *simpleArc) GraphSourceID() ID {
	return e.source
}

func (e *simpleArc) GraphTargetID() ID {
	return e.target
}

func (e *simpleArc) Weight() int {
	return e.weight
}
