package medianfinder

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

type MedianFinder struct {
	low  maxHeap
	high minHeap
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {

	if mf.high.Len() == 0 {
		heap.Push(&mf.high, num)
		return
	}

	med := mf.FindMedian()

	if float64(num) >= med {
		heap.Push(&mf.high, num)
	} else {
		heap.Push(&mf.low, num)
	}

	if mf.low.Len() > mf.high.Len() {
		heap.Push(&mf.high, heap.Pop(&mf.low))
	}

	if mf.high.Len() > mf.low.Len()+1 {
		heap.Push(&mf.low, heap.Pop(&mf.high))
	}

}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() == mf.high.Len() {
		return float64((mf.low[0] + mf.high[0])) / 2
	}
	return float64(mf.high[0])
}
