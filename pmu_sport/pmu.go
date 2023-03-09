package PMU

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func cleanOdds(bet string) float64 {

	odds := strings.ReplaceAll(bet, ",", ".")
	oddsF, err := strconv.ParseFloat(odds, 32)
	if err != nil {
		log.Fatalln("amam", err)
	}
	oddsF = float64(int(oddsF*100)) / 100
	return oddsF
}

// SCRAPPING À REFAIRE CAR SCRAPP PAS TOUS LES MATCH QUE CEUX QUI S'AFFICHE PAS CEUX AVEC VOIR PLUS D'ÉLÉMENTS ................
func get_page() PMUData {
	var PmuData PMUData

	var data []string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"),
	)
	i := 0
	// Find and visit all links
	c.OnHTML("div.pmu-event-list-grid-highlights-formatter-row", func(e *colly.HTMLElement) {
		var cleanMatch CleanMatch

		if e.Text != "" && e.Text != "\n" {
			data = append(data, strings.ReplaceAll(e.Text, " ", ""))

		}
		temp := e.ChildText(".trow--event ")
		temp = strings.ReplaceAll(temp, " ", "")
		temp = strings.ReplaceAll(temp, "\n", "")
		cleanMatch.Title = temp
		temp = strings.Split(e.ChildText(".trow--event "), `//`)[0]
		temp = strings.ReplaceAll(temp, " ", "")
		temp = strings.ReplaceAll(temp, "\n", "")
		cleanMatch.Bet1.Title = strings.ReplaceAll(temp, " ", "")
		temp = strings.Split(e.ChildText(".trow--event "), `//`)[1]
		temp = strings.ReplaceAll(temp, " ", "")
		temp = strings.ReplaceAll(temp, "\n", "")
		cleanMatch.Bet2.Title = strings.ReplaceAll(temp, " ", "")
		cleanMatch.Draw.Title = "DRAW "
		bet1 := strings.Split(e.ChildText(".hierarchy-outcome-price"), "\n")[0]
		bet1 = strings.ReplaceAll(bet1, " ", "")
		draw := strings.Split(e.ChildText(".hierarchy-outcome-price"), "\n")[1]
		bet2 := strings.Split(e.ChildText(".hierarchy-outcome-price"), "\n")[2]
		draw = strings.ReplaceAll(draw, " ", "")
		bet2 = strings.ReplaceAll(bet2, " ", "")

		cleanMatch.Bet1.Odd = cleanOdds(bet1)
		cleanMatch.Bet2.Odd = cleanOdds(bet2)
		cleanMatch.Draw.Odd = cleanOdds(draw)

		PmuData = append(PmuData, cleanMatch)

		// fmt.Println("odds ? = ", e.ChildText(".hierarchy-outcome-price"))
		// bet1 := strings.Split(e.ChildText(".hierarchy-outcome-price"), " ")[1]
		// bet2 := strings.Split(e.ChildText(".hierarchy-outcome-price"), " ")[1]

		// draw := strings.Split(bet2, " ")[1]
		// fmt.Println("draw  = ", draw)

		// PmuData[i].Bet2.Title = strings.Split(e.ChildText(".trow--event "), `//`)[1]

		i += 1

		// }

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://paris-sportifs.pmu.fr/pari/sport/25/football")
	return PmuData

}
func Pmu() PMUData {
	pmuData := get_page()
	return pmuData
}
