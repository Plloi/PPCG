package dailycodingproblem

func find(k int, s string) string {
	charmap := make(map[rune]int)
	ret, tmp := "", ""

	for _, c := range s {
		charmap[c]++
		tmp += string(c)

		for len(charmap) > k {
			c := rune(tmp[0])
			charmap[c]--
			if charmap[c] == 0 {
				delete(charmap, c)
			}
			tmp = tmp[1:len(tmp)]
		}

		if len(tmp) > len(ret) {
			ret = tmp
		}
	}
	return ret
}
