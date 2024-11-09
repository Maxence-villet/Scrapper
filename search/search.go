package search

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Search permet de rechercher dans un fichier
func Search(filename string) error {

	// Ouvrir le fichier en lecture
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var list_search = [4]string{"project", "build", "apps", "scrum"}
	var blacklist_search = [1]string{"project-management"}

	// Créer un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ici, vous pouvez ajouter votre logique de recherche
		for _, element := range list_search {
			// fmt.Println(strings.ToUpper(line), strings.ToUpper(element))
			for _, element_blacklist := range blacklist_search {
				if strings.Contains(strings.ToUpper(line), strings.ToUpper(element_blacklist)) {

				} else if strings.Contains(strings.ToUpper(line), strings.ToUpper(element)) {
					f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
					if err != nil {
						fmt.Println("Erreur lors de l'ouverture du fichier :", err)
					}
					defer f.Close()
					if _, err := f.Write([]byte("comidoc.net" + line + "\n")); err != nil {
						fmt.Println("Erreur lors de l'écriture :", err)
					}
					fmt.Println(line)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
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
