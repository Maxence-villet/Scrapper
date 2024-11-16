package bot

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"scrap.com/config"
	"scrap.com/crypto"
	"scrap.com/data"
	"scrap.com/scrap"
)

type BotHandler interface {
	SendMessage() error
}

type Bot struct {
	token      string
	timeReload int
	channel_id string
}

func (b *Bot) SendMessage() {
	sess, err := discordgo.New(b.token)
	if err != nil {
		log.Fatal(err)
	}

	// Fonction pour envoyer le contenu du fichier dans un canal
	sendMessage := func(s *discordgo.Session, channelID string) {

		dataHandler := data.NewData()

		file, err := os.Open(dataHandler.GetFilename())
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
			s.ChannelMessageSend(b.channel_id, scanner.Text())
		}

	}

	// Boucle infinie pour envoyer le message toutes les 900 secondes
	go func() {
		for {

			sendMessage(sess, "votre_channel_id")
			time.Sleep(config.NewConfig().GetTimeReload())

			dataHandler := data.NewData()
			dataHandler.RemoveCache()

			scrapHandler := scrap.NewScrap()
			scrapHandler.Scrap()

			dataHandler.RemoveDuplicates()

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

func NewBot() *Bot {

	tokenTmp, err := crypto.Encrypt(config.NewConfig().GetKey(), config.NewConfig().GetToken())
	if err != nil {
		log.Fatal(err)
	}

	tokenDecryptTmp, err := crypto.Decrypt(config.NewConfig().GetKey(), tokenTmp)

	b := &Bot{
		token:      string(tokenDecryptTmp),
		timeReload: 3,
		channel_id: config.NewConfig().GetChannelId(),
	}

	return b
}
