package main

import (
	"scrap.com/bot"
	"scrap.com/data"
	"scrap.com/scrap"
)

func main() {

	dataHandler := data.NewData()
	dataHandler.RemoveCache()
	dataHandler.CreateCache()

	scrapHandler := scrap.NewScrap()
	scrapHandler.Scrap()

	dataHandler.RemoveDuplicates()

	botHandler := bot.NewBot()
	botHandler.SendMessage()

}
