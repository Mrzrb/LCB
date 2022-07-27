package main

import "testing"

func checkEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Tmp := make([]string, len(str1))
	copy(str1Tmp, str1)
	str2Tmp := make([]string, len(str2))
	copy(str2Tmp, str2)

	for i := 0; i < len(str1); i++ {
		if str1Tmp[i] != str2Tmp[i] {
			return false
		}
	}

	return true
}

func TestEmptyRune(t *testing.T) {
	excepted := []string{}
	actual := []string{}

	perm([]rune(""), func(a []rune) {
		actual = append(actual, string(a))
	}, 0)

	if !checkEqual(actual, excepted) {
		t.Errorf("Wrong, empty rune do not have result")
	}

}

func TestDuplicateRune(t *testing.T) {
	excepted := []string{"AB", "BA"}
	actual := []string{}
	perm([]rune("AB"), func(a []rune) {
		actual = append(actual, string(a))
	}, 0)
	if !checkEqual(actual, excepted) {
		t.Errorf("Wrong, empty rune do not have result")
	}
}

func FuzzRune(f *testing.F) {
	f.Add("teasdasd")

	f.Fuzz(func(t *testing.T, a string) {
		result := []string{}

		perm([]rune(a), func(r []rune) {
			result = append(result, string(r))
		}, 0)
	})
}
