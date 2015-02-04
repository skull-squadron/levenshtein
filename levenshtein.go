package levenshtein

import "strings"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Distance(s, t string) int {
	if s == t {
		return 0
	}
	slen, tlen := len(s), len(t)
	if slen == 0 {
		return tlen
	}
	if tlen == 0 {
		return slen
	}
	if slen < tlen {
		slen, tlen = tlen, slen
		s, t = t, s
	}
	var cost, i, j int
	v0 := make([]int, tlen+1)
	v1 := make([]int, tlen+1)
	for i := range v0 {
		v0[i] = i
	}

	for i = 0; i < slen; i++ {
		v1[0] = i + 1
		for j = 0; j < tlen; j++ {
			if s[i] == t[j] {
				cost = 0
			} else {
				cost = 1
			}
			v1[j+1] = min(min(v1[j], v0[j+1])+1, v0[j]+cost)
		}
		copy(v0, v1)
	}
	return v1[tlen]
}

func DistanceCaseInsensitive(s, t string) int {
	return Distance(strings.ToLower(s), strings.ToLower(t))
}
