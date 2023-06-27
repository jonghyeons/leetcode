type Pair struct {
	Sum   int
	Nums1 int
	Nums2 int
}

type PairHeap []Pair

func (h PairHeap) Len() int            { return len(h) }
func (h PairHeap) Less(i, j int) bool  { return h[i].Sum < h[j].Sum }
func (h PairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	h := &PairHeap{}
	heap.Init(h)

	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Pair{nums1[i] + nums2[0], i, 0})
	}

	result := [][]int{}
	for k > 0 && h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		result = append(result, []int{nums1[p.Nums1], nums2[p.Nums2]})

		if p.Nums2+1 < len(nums2) {
			heap.Push(h, Pair{nums1[p.Nums1] + nums2[p.Nums2+1], p.Nums1, p.Nums2 + 1})
		}

		k--
	}

	return result
}