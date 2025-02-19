package bot

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"scrap.com/data"
	"scrap.com/scrap"

	"github.com/joho/godotenv"
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

	sendMessage(sess, "votre_channel_id")

	dataHandler := data.NewData()
	dataHandler.RemoveCache()

	scrapHandler := scrap.NewScrap()
	scrapHandler.Scrap()

	dataHandler.RemoveDuplicates()

	fmt.Println("the bot is online")

	// Ouvre une websocket vers Discord et bloque jusqu'à ce que vous appuyez sur Ctrl+C
	err = sess.Open()
	if err != nil {
		log.Fatal("Cannot open the session,", err)
	}
	defer sess.Close()

}

func NewBot() *Bot {

	// Charger me fichier .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Récupérer le token et le channel_id
	apiKey := os.Getenv("API_KEY")
	channelId := os.Getenv("CHANNEL_ID")

	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}

	if channelId == "" {
		log.Fatal("CHANNEL_ID is not set")
	}

	b := &Bot{
		token:      apiKey,
		timeReload: 3,
		channel_id: channelId,
	}

	return b
}
