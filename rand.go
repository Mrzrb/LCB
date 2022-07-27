package main

import (
	"math/rand"
	"time"
)

func rand13() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(12) + 1
}

func rand5() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5) + 1
}

func rand5rand13() int {
	for {
		v := 5*(rand5()-1) + rand5()

		if v > 13 {
			continue
		}

		return v

	}
}

func rand13rand5() int {
	for {
		v := rand13()

		if v > 10 {
			continue
		}

		return v / 2
	}
}
