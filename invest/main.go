package main

import (
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

	scraper := scrape.NewScraper(nil)
	db, err := db.NewStorage("root:root@tcp(127.0.0.1:3300)/investdb?charset=utf8mb4&parseTime=True&loc=Local") // TODO. dsn config화
	if err != nil {
		panic(err)
	}
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	event := event.NewEvent(db, scraper, conf)

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

	for true {
		msg := <-c
		teleBot.SendMessage(msg)
		log.Println(msg)
	}
}
