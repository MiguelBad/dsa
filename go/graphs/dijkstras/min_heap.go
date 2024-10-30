package main

type Distances struct {
	Vertex   *Vertex
	Distance int
}

type MinHeap struct {
	elements []*Distances
}

func distanceNode(item *Vertex, distance int) *Distances {
	return &Distances{
		Vertex:   item,
		Distance: distance,
	}
}

func (h *MinHeap) insert(item *Vertex, distance int) {
	newDistance := distanceNode(item, distance)
	h.elements = append(h.elements, newDistance)
	h.heapifyUp(len(h.elements) - 1)
}

func (h *MinHeap) delete() *Vertex {
	if len(h.elements) < 1 {
		return nil
	}

	item := h.elements[0].Vertex
	if len(h.elements) == 1 {
		h.elements = h.elements[:0]
	} else {
		h.elements[0] = h.elements[len(h.elements)-1]
		h.elements = h.elements[:len(h.elements)-1]
		h.heapifyDown(0)
	}
	return item
}

func (h *MinHeap) heapifyDown(currIdx int) {
	left := h.leftIdx(currIdx)
	right := h.rightIdx(currIdx)
	smallest := currIdx

	if left < len(h.elements) && h.elements[left].Distance < h.elements[currIdx].Distance {
		smallest = left
	}

	if right < len(h.elements) && h.elements[right].Distance < h.elements[currIdx].Distance {
		smallest = right
	}

	if smallest != currIdx {
		h.elements[smallest], h.elements[currIdx] = h.elements[currIdx], h.elements[smallest]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) heapifyUp(currIdx int) {
	parent := h.parentIdx(currIdx)

	if parent < 0 {
		return
	}

	if h.elements[parent].Distance > h.elements[currIdx].Distance {
		h.elements[currIdx], h.elements[parent] = h.elements[parent], h.elements[currIdx]
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) leftIdx(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) rightIdx(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) parentIdx(idx int) int {
	return (idx - 1) / 2
}
