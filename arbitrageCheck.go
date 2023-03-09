package main

import "fmt"

func ArbitrageCheck(bet1 float64, bet2 float64, draw float64) bool {
	arbitragePrct := ((1 / bet1) * 100) + ((1 / bet2) * 100) + ((1 / draw) * 100)
	if arbitragePrct < 99 {
		fmt.Println("pourcentage :", 100-arbitragePrct, " %")
		return true
	}
	return false
}
