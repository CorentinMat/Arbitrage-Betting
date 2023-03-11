package data

import (
	PMU "Arbitrage_Betting/pmu_sport"
	"Arbitrage_Betting/winamax"
	"fmt"
)

func StringMatching(s1 string, s2 string) bool {
	var km int
	str1 := []rune(s1)
	str2 := []rune(s2)
	// Gestalt pattern matching
	if len(str1) > len(str2) {
		for i := range str2 {
			if str1[i] == str2[i] {
				km += 1
				str2[i] = ' '
				continue
			}

		}
	} else {
		for i := range str1 {
			if str1[i] == str2[i] {
				km += 1
				str2[i] = ' '
				continue
			}

		}
	}

	result := float32(2*km) / float32(len(s1)+len(s2))

	if result > 0.365 {
		fmt.Println("coef => ", result)
		return true
	} else {
		return false

	}
}

func CleanAllData() string {
	winamaxData := winamax.Winamax()
	pmuData := PMU.Pmu()
	// Analyse de toutes les données pour voir ceux en commun
	for i := range winamaxData {
		for k := range pmuData {
			s1 := winamaxData[i].Title
			s2 := pmuData[k].Title
			if StringMatching(s1, s2) {
				var bet1 []float64
				bet1 = append(bet1, winamaxData[i].Bet1.Odd)
				bet1 = append(bet1, pmuData[k].Bet1.Odd)
				var bet2 []float64
				bet2 = append(bet2, winamaxData[i].Bet2.Odd)
				bet2 = append(bet2, pmuData[k].Bet2.Odd)
				var draw []float64
				draw = append(draw, winamaxData[i].Draw.Odd)
				draw = append(draw, pmuData[k].Draw.Odd)
				if ArbitrageCheck(bet1, bet2, draw) {

					fmt.Println("OMGGGGG YA UN TRUC LA => ", winamaxData[i].Title, "////", ArbitrageCheck(bet1, bet2, draw))

				}

			}

			// if winamaxData[i].Title == pmuData[k].Title {
			// 	fmt.Println(pmuData[k])
			// }

		}
	}

	return "y'a r à voir pour le moment"
}
