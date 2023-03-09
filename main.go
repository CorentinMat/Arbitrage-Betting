package main

import (
	PMU "Arbitrage_Betting/pmu_sport"
	"Arbitrage_Betting/winamax"
	"fmt"
)

func main() {
	winamaxData := winamax.Winamax()
	for k := range winamaxData {
		fmt.Println("test => ", *winamaxData[k])
	}
	pmuData := PMU.Pmu()
	for k := range pmuData {
		fmt.Println(pmuData[k])
	}

}
