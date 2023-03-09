package winamax

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func get_page() string {
	client := &http.Client{}
	url := "https://www.winamax.fr/paris-sportifs/sports/1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set(
		"User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
	)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	data := string(body)
	data = strings.Split(data, `var PRELOADED_STATE = `)[1]
	data = strings.Split(data, `;</script>`)[0]
	return data

}

func get_json() WinamaxData {
	// all json data on website Winamax
	var data Data
	AllJson := get_page()
	// fmt.Println(AllJson)

	err := json.Unmarshal([]byte(AllJson), &data)
	if err != nil {
		log.Fatalln("Can't convert to JSON : ", err)
	}
	// label + id  des cotes
	outcomes := data.Outcomes
	// bet info
	bets := data.Bets
	// toutes le cote
	odds := data.Odds
	// match info
	matches := data.Matches

	return CleanJson(odds, outcomes, bets, matches)

}
func CleanJson(odds map[string]float64, outcomes map[string]OutcomesId, bets map[string]BetId, matches map[string]MatchId) WinamaxData {
	// cleanOdds := make(map[string]float32) // get odds for players
	// betId := make(map[int]int)
	winamaxData := make(WinamaxData)

	// fmt.Println(cleanOdds["JessicaPegula"])
	// for k := range bets {
	// 	for i := range bets[k].Outcomes {
	// 		betId[(bets[k].Outcomes[i])] = bets[k].BetId
	// 	}
	// }
	for k := range matches {
		result := strconv.Itoa(matches[k].MainBetId)
		var oddsId1 int
		var oddsId2 int
		// je sais pas pk ya un 0 qui c'est mit dans les data  WINAMAX ?????
		if result != "0" {
			oddsId1 = bets[result].Outcomes[0]
			oddsId2 = bets[result].Outcomes[1]
		}

		// initialize check if there are no tennis matches in ....
		draw := 0
		if len(bets[result].Outcomes) == 3 {

			draw = bets[result].Outcomes[2]
		}

		winamaxData[k] = &CleanMatch{"id", CleanBet{"id1", 1.3}, CleanBet{"id2", 1.3}, CleanBet{"draw", 1.3}}

		winamaxData[k].Title = matches[k].Title
		// odds 1
		winamaxData[k].Bet1.Title = outcomes[strconv.Itoa(oddsId1)].Label
		winamaxData[k].Bet1.Odd = odds[strconv.Itoa(oddsId1)]
		// odds 2
		winamaxData[k].Bet2.Title = outcomes[strconv.Itoa(oddsId2)].Label
		winamaxData[k].Bet2.Odd = odds[strconv.Itoa(oddsId2)]
		// Draw
		winamaxData[k].Draw.Title = "Draw"
		winamaxData[k].Draw.Odd = odds[strconv.Itoa(draw)]
		// fmt.Println(*winamaxData[k])
	}
	return winamaxData

}

func Winamax() WinamaxData {
	winamaxData := get_json()

	return winamaxData
	// ArbitrageCheck(winamaxData)
}
