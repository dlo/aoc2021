package utils

type PointWithRisk struct {
	point Point
	score int
}
type PointHeap []PointWithRisk

func (h PointHeap) Len() int {
	return len(h)
}

func (h PointHeap) Less(i, j int) bool {
	return h[i].score < h[j].score
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(PointWithRisk))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
