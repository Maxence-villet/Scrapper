package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"scrap.com/m/scrap"
	"scrap.com/m/search"
)

func main() {
	scrap.Srap()
	search.Search("data.txt")

	sess, err := discordgo.New("")
	if err != nil {
		log.Fatal(err)
	}

	// Fonction pour envoyer le contenu du fichier dans un canal
	sendMessage := func(s *discordgo.Session, channelID string) {
		file, err := os.Open("result.txt")
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		if err := scanner.Err(); err != nil {
			fmt.Println("Erreur lors de la lecture du fichier:", err)
		}
		for scanner.Scan() {
			s.ChannelMessageSend("1304593207732080685", scanner.Text())
		}

	}

	// Boucle infinie pour envoyer le message toutes les 30 secondes
	go func() {
		for {
			sendMessage(sess, "votre_channel_id")
			time.Sleep(900 * time.Second)

		}
	}()

	fmt.Println("the bot is online")

	// Ouvre une websocket vers Discord et bloque jusqu'à ce que vous appuyez sur Ctrl+C
	err = sess.Open()
	if err != nil {
		log.Fatal("Cannot open the session,", err)
	}
	defer sess.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	<-make(chan struct{})
}

func remove_cache(filename string) {
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}
