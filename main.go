package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"encoding/json"
)
type item struct {
  Name string `json:"name"`
  Price string `json:"price"`
  ImgUrl string `json:"ImgUrl"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("j2store.net"),
	)
	var items[]  item


	// On every a element which has href attribute call callback
	c.OnHTML("div[itemprop=itemListElement]", func(e *colly.HTMLElement) {
		item:=item{
			Name:e.ChildText("h2.product-title"),
			Price: e.ChildText("div.sale-price"),
			ImgUrl: e.ChildAttr("img","src"),
		}
	  //	link := e.Attr("href")
	//	c.Visit(e.Request.AbsoluteURL(link))
		items = append(items,item)
	})

	c.OnHTML("[title=Next]", func(e *colly.HTMLElement){
		next_page := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request){
		fmt.Println(r.URL.String())
	})
	// Before making a request print "Visiting ..."
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL.String())
	//})

	// Start scraping on https://hackerspaces.org
  c.Visit("http://j2store.net/demo/index.php/shop")
//   fmt.Println(items)
		content, err := json.Marshal(items)

  if err != nil {
	fmt.Println(err.Error())
  }

  os.WriteFile("products.json", content, 0644)


}
