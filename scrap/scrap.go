package scrap

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

/*
Cette fonction scrappe les url de comidoc et les insères dans un fichier txt

Paramètres:
-----------

	Rien

Retournes:
----------

	Rien
*/
func Srap() error {
	remove_cache("data.txt")
	c := colly.NewCollector()

	c.OnHTML(".group", func(e *colly.HTMLElement) {
		e.ForEach(".group.relative a[href]", func(_ int, link *colly.HTMLElement) {
			href := link.Attr("href")
			if href[0:7] == "/udemy/" {
				var list_search = [4]string{"project", "build", "apps", "scrum"}
				var blacklist_search = [1]string{"project-management"}
				for _, element := range list_search {
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

/*
Cette fonction supprime le fichier 'example.txt'

Paramètres:
-----------

	filename (string) : le nom du fichier (par exemple: data.txt)

Retournes:
----------

	Rien
*/
func remove_cache(filename string) {
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}
