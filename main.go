package main

import (
	"fmt"
	"kenchen0625/webscrapingpractice/invoice"
	"log"
)

func main() {
	invoice := invoice.Invoice{
		URL: "https://invoice.etax.nat.gov.tw",
	}
	_, err := invoice.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	prizes, err := invoice.GetJackpot()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prizes)
}
