package main

import (
	DATA "Arbitrage_Betting/Clean_All_Data"
	"fmt"
)

func main() {
	data := DATA.CleanAllData()
	fmt.Println("Final Result =>", data)

}
