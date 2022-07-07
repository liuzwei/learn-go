package swordoffer

func ReverseLeftWords(s string, n int) string {
	for n > 0 {
		temp := s[0]
		s = s[1:] + string(temp)
		n--
	}
	return s
}
