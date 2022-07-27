package main

import "testing"

func TestRand13rand5(t *testing.T) {
	v := rand13rand5()
	if v > 5 {
		t.Error("error, rand can not great than 5")
	}
}
