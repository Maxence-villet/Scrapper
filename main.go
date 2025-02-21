package main

import (
	argumentmanager "scrap.com/argumentManager"
	"scrap.com/bot"
	"scrap.com/data"
	"scrap.com/scrap"
)

func main() {

	args := argumentmanager.GetArguments()
	args = argumentmanager.FilterArguments(args)

	dataHandler := data.NewData()
	dataHandler.RemoveCache()
	dataHandler.CreateCache()

	scrapHandler := scrap.NewScrap()
	scrapHandler.Scrap(args)

	dataHandler.RemoveDuplicates()

	botHandler := bot.NewBot()
	botHandler.SendMessage()

}
