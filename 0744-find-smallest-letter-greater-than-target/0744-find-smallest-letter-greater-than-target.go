func nextGreatestLetter(letters []byte, target byte) byte {
	idx := sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})

	if idx < len(letters) {
		return letters[idx]
	}

	return letters[0]
}