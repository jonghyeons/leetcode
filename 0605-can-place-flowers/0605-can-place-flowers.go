func canPlaceFlowers(flowerbed []int, n int) bool {
	max := 0
	for i := 0; max < n && i < len(flowerbed); i++ {
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}
		if flowerbed[i] != 0 {
			continue
		}

		flowerbed[i] = 1
		max++
		if i < len(flowerbed)-1 {
			flowerbed[i+1] = -1
		}
	}
	return max >= n
}