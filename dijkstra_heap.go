// Copyright 2009 The Go Authors, 2018 Leonard Hecker. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found at:
// https://github.com/golang/go/blob/master/LICENSE.

// The contents of this file are an inlined version of the code found in Go's "container/heap" package.
// Due to this optimization Dijkstra()'s performance was improved by nearly 50%.
package gograph

type dijkstraHeap []*dijkstraState

func (h *dijkstraHeap) PushMaximum(tail *dijkstraState) {
	tail.HeapIndex = len(*h)
	*h = append(*h, tail)
}

func (h *dijkstraHeap) Pop() *dijkstraState {
	heap := *h

	idx := len(heap) - 1
	h.swap(0, idx)
	h.down(0, idx)

	tail := heap[idx]
	heap[idx] = nil

	*h = heap[:idx]
	return tail
}

func (h dijkstraHeap) Fix(item *dijkstraState) {
	idx := item.HeapIndex

	if !h.down(idx, len(h)) {
		h.up(idx)
	}
}

func (h dijkstraHeap) less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h dijkstraHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].HeapIndex = i
	h[j].HeapIndex = j
}

func (h dijkstraHeap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h dijkstraHeap) down(i0, n int) bool {
	i := i0

	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2 // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}

	return i > i0
}
