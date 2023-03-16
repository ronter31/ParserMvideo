package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)


func main() {
URL := "http://www.mvideo.ru/noutbuki-planshety-komputery-8/noutbuki-118"

geziyor.NewGeziyor(&geziyor.Options{
	StartURLs: []string{URL},
	ParseFunc: quatesParse,
	Exporters: []export.Exporter{&export.JSON{}},
}).Start()

}

func quatesParse(g *geziyor.Geziyor,r *client.Response){
	r.HTMLDoc.Find("div.product-cards-layout product-cards-layout--list").Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"text": strings.TrimSpace(s.Find("a.product-title__text").Text()),
			"price": s.Find("span.price__main-value").Text(),
		}
	})

	
}