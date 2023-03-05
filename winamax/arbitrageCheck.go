package main

import "fmt"

func ArbitrageCheck(winamaxData WinamaxData) {

	for k := range winamaxData {
		arbitragePrct := ((1 / winamaxData[k].Bet1.Odd) * 100) + ((1 / winamaxData[k].Bet2.Odd) * 100) + ((1 / winamaxData[k].Draw.Odd) * 100)
		if arbitragePrct < 105 {
			fmt.Println("pourcentage :", arbitragePrct)
			fmt.Println(*winamaxData[k])
			fmt.Println("ID :", k)
		}
	}
}
