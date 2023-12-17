func isAnagram(s string, t string) bool {
	s1 := strings.Split(s, "")
	s2 := strings.Split(t, "")

	sort.Strings(s1)
	sort.Strings(s2)

	if strings.Join(s1, "") == strings.Join(s2, "") {
		return true
	}
	return false
}