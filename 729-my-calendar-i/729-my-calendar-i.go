type MyCalendar struct {
	Schedule [][]int
}

func Constructor() MyCalendar {
	return *&MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	isBookingPossible := true
	for i := range mc.Schedule {
		min := mc.Schedule[i][0]
		max := mc.Schedule[i][1]

		if start > min {
			min = start
		}

		if end < max {
			max = end
		}

		if min < max {
			isBookingPossible = false
			break
		}
	}

	if isBookingPossible {
		mc.Schedule = append(mc.Schedule, []int{start, end})
		return true
	}
	return false
}