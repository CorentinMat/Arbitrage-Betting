package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func get_page() string {
	var PmuData PMUData

	var data []string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"),
	)
	i := 0
	// Find and visit all links
	c.OnHTML("div.pmu-event-list-grid-highlights-formatter-row", func(e *colly.HTMLElement) {
		var cleanMatch CleanMatch
		var err error

		if e.Text != "" && e.Text != "\n" {
			data = append(data, strings.ReplaceAll(e.Text, " ", ""))

		}
		cleanMatch.Title = e.ChildText(".trow--event ")
		cleanMatch.Bet1.Title = strings.Split(e.ChildText(".trow--event "), `//`)[0]
		bet1 := strings.Split(e.ChildText(".hierarchy-outcome-price"), " ")[0]
		cleanMatch.Bet1.Odd, err = strconv.ParseFloat(bet1, 32)
		if err != nil {
			log.Fatalln(err)
		}
		// bet2 := strings.Split(e.ChildText(".hierarchy-outcome-price"), " ")[1]
		// draw := strings.Split(bet2, " ")[1]
		// bet2 = strings.Split(bet2, " ")[0]

		fmt.Println("bets 1  = ", bet1)
		// fmt.Println("bets 2  = ", bet2)
		// fmt.Println("draw  = ", draw)

		PmuData = append(PmuData, cleanMatch)

		// fmt.Println("odds ? = ", e.ChildText(".hierarchy-outcome-price"))
		// bet1 := strings.Split(e.ChildText(".hierarchy-outcome-price"), " ")[1]
		// bet2 := strings.Split(e.ChildText(".hierarchy-outcome-price"), " ")[1]

		// draw := strings.Split(bet2, " ")[1]
		// fmt.Println("draw  = ", draw)

		// PmuData[i].Bet2.Title = strings.Split(e.ChildText(".trow--event "), `//`)[1]

		i += 1

		// }
		fmt.Println(PmuData)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://paris-sportifs.pmu.fr/?_ga=2.268433678.1296634928.1677849818-1205291434.1677849818")
	return "temp"
}
func main() {
	fmt.Println(get_page())
}
