package main

import (
	argumentmanager "scrap.com/argumentManager"
	"scrap.com/bot"
	"scrap.com/data"
	"scrap.com/scrap"
)

func main() {

	argsNotFiltered := argumentmanager.GetArguments()
	args, blacklist := argumentmanager.FilterArguments(argsNotFiltered)

	dataHandler := data.NewData()
	dataHandler.RemoveCache()
	dataHandler.CreateCache()

	scrapHandler := scrap.NewScrap()
	scrapHandler.Scrap(args, blacklist)

	dataHandler.RemoveDuplicates()

	botHandler := bot.NewBot()
	botHandler.SendMessage()

}
