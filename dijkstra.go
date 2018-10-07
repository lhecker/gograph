package gograph

import (
	"fmt"
	"math"
	"strconv"
)

var (
	infinity = math.Inf(1)
)

type dijkstraState struct {
	ID        ID
	Previous  *dijkstraState
	Distance  float64
	HeapIndex int
	Visited   bool
}

type DijkstraUnreachableError struct {
	SourceID ID
	TargetID ID
}

func (s *DijkstraUnreachableError) Error() string {
	return fmt.Sprintf("no path found for %v -> %v", s.SourceID, s.TargetID)
}

type DijkstraUnknownNodeError struct {
	NodeID ID
}

func (s *DijkstraUnknownNodeError) Error() string {
	return fmt.Sprintf("unknown node %v", s.NodeID)
}

type DijkstraNegativeWeightError struct {
	SourceID ID
	TargetID ID
	Weight   float64
}

func (s *DijkstraNegativeWeightError) Error() string {
	return fmt.Sprintf("negative weight %s on arc %v -> %v",
		strconv.FormatFloat(s.Weight, 'g', -1, 64),
		s.SourceID,
		s.TargetID,
	)
}

func Dijkstra(graph DirectedGraph, source ID, target ID) ([]ID, float64, error) {
	nodes := graph.GetNodes()
	arcs := graph.GetArcs()

	if _, ok := nodes[source]; !ok {
		return nil, infinity, &DijkstraUnknownNodeError{
			NodeID: source,
		}
	}
	if _, ok := nodes[target]; !ok {
		return nil, infinity, &DijkstraUnknownNodeError{
			NodeID: target,
		}
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
			Distance: infinity,
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
			if vdata == nil {
				return nil, infinity, &DijkstraUnknownNodeError{
					NodeID: v,
				}
			}
			if vdata.Visited {
				continue
			}

			weight := arc.Weight()
			if weight < 0 {
				return nil, infinity, &DijkstraNegativeWeightError{
					SourceID: udata.ID,
					TargetID: v,
					Weight:   weight,
				}
			}

			alt := distance + weight

			if alt < vdata.Distance {
				vdata.Distance = alt
				vdata.Previous = udata
				Q.Fix(vdata)
			}
		}
	}

	root := state[target]
	if math.IsInf(root.Distance, 1) {
		return nil, infinity, &DijkstraUnreachableError{
			SourceID: source,
			TargetID: target,
		}
	}
	if root.Previous == nil {
		return nil, 0, nil
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

	return path, root.Distance, nil
}
