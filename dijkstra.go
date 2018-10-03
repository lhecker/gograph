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

func Dijkstra(graph DirectedGraph, source ID, target ID) []ID {
	nodes := graph.GetNodes()
	arcs := graph.GetArcs()

	if _, ok := nodes[source]; !ok {
		return nil
	}
	if _, ok := nodes[target]; !ok {
		return nil
	}

	state := make(map[ID]*dijkstraState, len(nodes))
	Q := make(dijkstraHeap, 0, len(nodes))

	sourceData := &dijkstraState{
		ID:       source,
		Distance: 0,
	}
	state[source] = sourceData
	Q.PushMaximum(sourceData)

	for id := range nodes {
		if id == source {
			continue
		}

		data := &dijkstraState{
			ID:       id,
			Distance: maxInt,
		}
		state[id] = data
		Q.PushMaximum(data)
	}

	for len(Q) != 0 {
		udata := Q.Pop()
		if udata.ID == target {
			break
		}

		distance := udata.Distance
		udata.Visited = true

		for v, arc := range arcs[udata.ID] {
			vdata := state[v]
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

	root := state[target]
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
