package scrap

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type ScrapHandler interface {
	Scrap() error
}

type Scrap struct {
}

func (b *Scrap) Scrap(keyWord []string) error {
	c := colly.NewCollector()

	c.OnHTML(".group", func(e *colly.HTMLElement) {
		e.ForEach(".group.relative a[href]", func(_ int, link *colly.HTMLElement) {
			href := link.Attr("href")
			if href[0:7] == "/udemy/" {
				var blacklist_search = [1]string{"project-management"}
				for _, element := range keyWord {
					for _, element_blocked := range blacklist_search {
						if strings.Contains(strings.ToUpper(href), strings.ToUpper(element_blocked)) {

						} else if strings.Contains(href, element) {
							data := []byte("comidoc.net" + href + "\n")
							f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
							if err != nil {
								fmt.Println("Erreur lors de l'ouverture du fichier :", err)
								return
							}
							defer f.Close()

							if _, err := f.Write(data); err != nil {
								fmt.Println("Erreur lors de l'écriture :", err)
								return
							}
						} else {
						}
					}
				}
			} else {
				// instructeur
			}

		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request to %s failed with error: %v", r.Request.URL.String(), err)
	})

	return c.Visit("https://comidoc.net/daily")

}

func NewScrap() *Scrap {
	s := &Scrap{}
	return s
}
