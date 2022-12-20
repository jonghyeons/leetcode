func canVisitAllRooms(rooms [][]int) bool {
	isVisited := make([]bool, len(rooms))
	queue := []int{0}
	for len(queue) > 0 {
		roomNumber := queue[0]
		queue = queue[1:]

		if isVisited[roomNumber] {
			continue
		}
		isVisited[roomNumber] = true

		for _, key := range rooms[roomNumber] {
			if !isVisited[key] {
				queue = append(queue, key)
			}
		}
	}

	for i := range isVisited {
		if !isVisited[i] {
			return false
		}
	}
	return true
}