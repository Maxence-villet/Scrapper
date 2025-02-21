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

func (b *Scrap) WriteData(href string) {
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
}

func (b *Scrap) Scrap(keyWord []string, blacklist []string) error {
	c := colly.NewCollector()

	c.OnHTML(".group", func(e *colly.HTMLElement) {
		e.ForEach(".group.relative a[href]", func(_ int, link *colly.HTMLElement) {
			href := link.Attr("href")
			blacklistcheck := true

			if href[0:7] == "/udemy/" {

				arg := strings.Replace(href, "/udemy/", "", -1)
				args := strings.Split(arg, "-")
				for _, a := range args {
					if blacklist != nil {
						for _, element_blacklisted := range blacklist {
							if a == element_blacklisted {
								blacklistcheck = false
							}
						}
					}
					if keyWord != nil {
						for _, k := range keyWord {
							if a == k {
								if blacklistcheck == true {
									b.WriteData(href)
									blacklistcheck = false
								}
							}
						}
					}
				}

				if keyWord == nil && blacklistcheck == true {
					b.WriteData(href)
					blacklistcheck = false
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
