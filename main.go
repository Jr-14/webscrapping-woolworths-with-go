package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

const WOOLWORTHS_DOMAIN = "https://www.woolworths.com.au/"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(WOOLWORTHS_DOMAIN),
	)

	c.OnHTML("div[class=ng-tns-c356-4 product-grid-v2--tile ng-star-inserted]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://www.woolworths.com.au/shop/browse/fruit-veg")
}
