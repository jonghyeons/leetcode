import (
	"container/heap"
)

type Edge struct {
	node     int
	prob     float64
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].prob > pq[j].prob
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]Edge, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		p := succProb[i]
		graph[a] = append(graph[a], Edge{b, p})
		graph[b] = append(graph[b], Edge{a, p})
	}

	dist := make([]float64, n)
	dist[start] = 1.0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Edge{start, 1.0})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Edge)
		node := cur.node
		prob := cur.prob

		if node == end {
			return prob
		}

		if prob < dist[node] {
			continue
		}

		for _, neighbor := range graph[node] {
			nextNode := neighbor.node
			nextProb := prob * neighbor.prob

			if nextProb > dist[nextNode] {
				dist[nextNode] = nextProb
				heap.Push(&pq, &Edge{nextNode, nextProb})
			}
		}
	}

	return 0
}