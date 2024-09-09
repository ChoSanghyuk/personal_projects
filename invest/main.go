package main

import (
	"invest/app"

	"invest/bot"
	"invest/config"
	"invest/db"
	"invest/event"
	"invest/scrape"
	"strconv"
	"time"

	"log"
)

func main() {
	// Create a new instance of the server

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	scraper := scrape.NewScraper(conf,
		scrape.WithKIS(conf.KisAppKey(), conf.KisAppSecret()),
	)

	db, err := db.NewStorage(conf.Dsn())
	if err != nil {
		panic(err)
	}
	event := event.NewEvent(db, scraper)

	c := make(chan string)

	chatId, err := strconv.ParseInt(conf.Telegram.ChatId, 10, 64)
	if err != nil {
		panic(err)
	}

	teleBot, err := bot.NewTeleBot(conf.Telegram.Token, chatId)
	if err != nil {
		panic(err)
	}

	go func() {
		for true {
			event.AssetEvent(c)
			time.Sleep(10 * time.Minute)
		}
	}()

	go func() {
		for true {
			event.RealEstateEvent(c)
			time.Sleep(10 * time.Minute)
		}
	}()

	go func() {
		app.Run(db, scraper)
	}()

	for true {
		msg := <-c
		teleBot.SendMessage(msg)
		log.Println(msg)
	}
}
