package gograph

const (
	maxInt = int(^uint(0) >> 1)
)

type dijkstraState struct {
	ID        ID
	Previous  *dijkstraState
	Distance  int
	HeapIndex int
	Visited   bool
}

func Dijkstra(graph *Graph, source ID, target ID) []ID {
	if _, ok := graph.vertices[source]; !ok {
		return nil
	}
	if _, ok := graph.vertices[target]; !ok {
		return nil
	}

	vertexData := make(map[ID]*dijkstraState, len(graph.vertices))
	Q := make(dijkstraHeap, 0, len(graph.vertices))

	sourceData := &dijkstraState{
		ID:       source,
		Distance: 0,
	}
	vertexData[source] = sourceData
	Q.PushMaximum(sourceData)

	for id := range graph.vertices {
		if id == source {
			continue
		}

		data := &dijkstraState{
			ID:       id,
			Distance: maxInt,
		}
		vertexData[id] = data
		Q.PushMaximum(data)
	}

	for len(Q) != 0 {
		udata := Q.Pop()
		if udata.ID == target {
			break
		}

		distance := udata.Distance
		udata.Visited = true

		for v, arc := range graph.arcs[udata.ID] {
			vdata := vertexData[v]
			if vdata.Visited {
				continue
			}

			alt := distance + arc.Weight()

			if alt < vdata.Distance {
				vdata.Distance = alt
				vdata.Previous = udata
				Q.Fix(vdata)
			}
		}
	}

	root := vertexData[target]
	if root.Distance == maxInt {
		return nil
	}

	pathLength := 0
	for data := root; data != nil; data = data.Previous {
		pathLength++
	}

	path := make([]ID, pathLength)
	idx := pathLength - 1

	for data := root; data != nil; data = data.Previous {
		path[idx] = data.ID
		idx--
	}

	return path
}
