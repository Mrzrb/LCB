package main

import (
	"fmt"
	"sort"
)

func main() {
	Perm([]rune("AAC"), func(a []rune) {
		fmt.Println(string(a))
	})
}

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func getSortedRune(a []rune) []rune {
	tmp := []int{}
	for _, v := range a {
		tmp = append(tmp, int(v))
	}

	sort.Ints(tmp)
	result := []rune{}

	for _, v := range tmp {
		result = append(result, rune(v))
	}

	return result
}

func perm(a []rune, f func([]rune), startIndex int) {
	length := len(a)
	if length == 0 {
		return
	}

	a = getSortedRune(a)

	path := []rune{}
	used := make([]bool, length)

	var backtrack func(track []rune, index int)
	backtrack = func(track []rune, index int) {
		if len(track) == length {
			f(track)
		}

		for i := 0; i < length; i++ {
			if !used[i] {
				if i > 0 && a[i] == a[i-1] && !used[i-1] {
					continue
				}
				track = append(track, a[i])
				used[i] = true
				backtrack(track, i+1)
				track = track[:len(track)-1]
				used[i] = false
			}
		}
	}

	backtrack(path, startIndex)
}
