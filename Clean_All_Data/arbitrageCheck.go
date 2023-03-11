package data

import "fmt"

func ArbitrageCheck(bet1 []float64, bet2 []float64, draw []float64) bool {
	for _, bet1 := range bet1 {
		for _, bet2 := range bet2 {
			for _, draw := range draw {
				arbitragePrct := ((1 / bet1) * 100) + ((1 / bet2) * 100) + ((1 / draw) * 100)
				if arbitragePrct < 99 {
					fmt.Println("pourcentage :", 100-arbitragePrct, " %")
					fmt.Println(" bet1 =>", bet1)
					fmt.Println(" bet2 =>", bet2)
					fmt.Println(" Draw =>", draw)

					return true
				}
			}
		}

	}

	return false
}
