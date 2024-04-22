func validPath(n int, edges [][]int, source int, destination int) bool {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	queue := list.New()
	queue.PushBack(source)

	visited := make(map[int]bool)
	visited[source] = true

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(int)
		queue.Remove(element)

		if node == destination {
			return true
		}

		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false
}