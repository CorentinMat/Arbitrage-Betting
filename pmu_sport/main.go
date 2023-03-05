package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func get_page() string {
	var PmuData PMUData
	var data []string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"),
	)
	var i int
	// Find and visit all links
	c.OnHTML("div.pmu-event-list-grid-highlights-formatter-row", func(e *colly.HTMLElement) {
		// fmt.Print("Link found -> ", e.Text)
		// if e.Text != "\\s" {

		if e.Text != "" && e.Text != "\n" {
			data = append(data, strings.ReplaceAll(e.Text, " ", ""))

		}
		PmuData[i].Title = e.ChildText(".trow--event ")
		PmuData[i].Bet1.Title = strings.Split(e.ChildText(".trow--event "), `//`)[0]
		// fais planté le code
		fmt.Println("odds ? = ", e.ChildText("a.hierarchy-outcome-price"))

		PmuData[i].Bet2.Title = strings.Split(e.ChildText(".trow--event "), `//`)[1]

		i += 1

		// }
		// fmt.Println(data)

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
