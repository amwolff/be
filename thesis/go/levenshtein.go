package levenshtein

import (
	"unicode/utf8"
)

func min(a, b, c int) int {
	ret := a
	if b < ret {
		ret = b
	}
	if c < ret {
		ret = c
	}
	return ret
}

func WagnerFischer(s, t string) int {
	m, n := utf8.RuneCountInString(s), utf8.RuneCountInString(t)
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		d[i][0] = i
	}
	for j := 1; j <= n; j++ {
		d[0][j] = j
	}
	j := 1
	for _, r2 := range t {
		i := 1
		for _, r1 := range s {
			var cost int
			if r1 != r2 {
				cost = 1
			}
			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+cost)
			i++
		}
		j++
	}
	return d[m][n]
}

func LinearSpace(s, t string) int {
	m, n := utf8.RuneCountInString(s), utf8.RuneCountInString(t)
	if m > n {
		s, t = t, s
		m, n = n, m
	}
	d := make([]int, m+1)
	for i := 1; i <= m; i++ {
		d[i] = i
	}
	j := 1
	for _, r2 := range t {
		prev := d[0]
		d[0] = j
		i := 1
		for _, r1 := range s {
			var cost int
			if r1 != r2 {
				cost = 1
			}
			curr := d[i]
			d[i] = min(curr+1, d[i-1]+1, prev+cost)
			prev = curr
			i++
		}
		j++
	}
	return d[m]
}
