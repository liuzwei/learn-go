package algorithm

func StrStr(haystack string, needle string) int {
	position := -1
	h := 0
	n := 0
	for h < len(haystack) && n < len(needle) {
		if haystack[h] == needle[n] {
			if position == -1 {
				position = h
			}
			n++
		} else if position != -1 {
			h = h - n
			n = 0
			position = -1
		}
		h++
	}
	if n != len(needle) {
		position = -1
	}

	return position
}
