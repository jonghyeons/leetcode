func minSpeedOnTime(dist []int, hour float64) int {
	minSpeed := 1
	maxSpeed := int(math.Pow10(7)) // 10^7

	for minSpeed <= maxSpeed {
		midSpeed := minSpeed + (maxSpeed-minSpeed)/2

		if canReach(dist, hour, float64(midSpeed)) {
			maxSpeed = midSpeed - 1
		} else {
			minSpeed = midSpeed + 1
		}
	}

	if minSpeed > int(math.Pow10(7)) {
		return -1
	}

	return minSpeed
}

func canReach(dist []int, hour float64, speed float64) bool {
	totalTime := 0.0
	for i := 0; i < len(dist)-1; i++ {
		totalTime += math.Ceil(float64(dist[i]) / speed)
	}

	totalTime += float64(dist[len(dist)-1]) / speed

	return totalTime <= hour
}