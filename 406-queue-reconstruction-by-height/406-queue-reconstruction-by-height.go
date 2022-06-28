func reconstructQueue(people [][]int) [][]int {
    var res [][]int
	sort.Sort(person(people))

	insert := func(idx int, person []int) {
		res = append(res, person)
		if len(res)-1 == idx {
			return
		}
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, p := range people {
		insert(p[1], p)
	}

	return res
}

type person [][]int

func (p person) Len() int {
	return len(p)
}

func (p person) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	} else {
		return p[i][0] > p[j][0]
	}
}

func (p person) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
