package main

import (
	"math"
	"math/rand"
)

func abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func getChance(percent float64) bool {
	if percent == 0 {
		return false
	}
	rate := percent / 100
	rate = 1 / rate
	a := int(math.Round(rate))
	if rand.Intn(a) == 0 {
		return true
	}
	return false
}

func getChanceTest() {
	for j := 0; j < 10; j++ {
		t_cnt := 0
		f_cnt := 0

		for i := 0; i < 100; i++ {

			if getChance(146) {
				t_cnt++
			} else {
				f_cnt++
			}
		}
		println(t_cnt)
		println(f_cnt)
	}
}
