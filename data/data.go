package data

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"scrap.com/config"
)

// Interface définissant les opérations sur les données
type DataHandler interface {
	GetFilename() string
	RemoveCache() error
	RemoveDuplicates() error
	CreateCache() error
}

// Structure implémentant l'interface DataHandler
type Data struct {
	filename      string
	filename_sort string
}

func (d *Data) GetFilename() string {
	return d.filename_sort
}

func (d *Data) RemoveCache() error {
	err := os.Remove(d.filename)
	if err != nil {
		log.Println("Erreur lors de la suppression du cache:", err)
		return err
	}

	err = os.Remove(d.filename_sort)
	if err != nil {
		log.Println("Erreur lors de la suppression du cache:", err)
		return err
	}

	d.CreateCache()
	return nil
}
func (d *Data) CreateCache() error {
	f, err := os.Create(d.filename)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	f, err = os.Create(d.filename_sort)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	return nil
}

func (d *Data) RemoveDuplicates() error {

	// Ouvrir les fichiers en lecture et écriture
	file, err := os.Open(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file_sort, err := os.OpenFile(d.filename_sort, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file_sort.Close()

	// Utiliser un map pour stocker les lignes déjà rencontrées
	seen := make(map[string]struct{})

	// Créer un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	_, err = fmt.Fprintln(file_sort, config.NewConfig().GetImage())
	if err != nil {
		return err
	}

	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := seen[line]; !ok {
			_, err := fmt.Fprintln(file_sort, line)
			if err != nil {
				return err
			}
			seen[line] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func NewData() *Data {
	d := &Data{
		filename:      "data.txt",
		filename_sort: "dataSort.txt",
	}
	return d
}
