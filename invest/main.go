package main

import (
	"invest/app"

	"invest/bot"
	"invest/config"
	"invest/db"
	"invest/event"
	"invest/scrape"
	"strconv"

	"log"

	"github.com/robfig/cron"
)

const (
	Every15Min    = "0 */15 * * * *"
	Every9Am      = "0 0 9 * * *"
	PortfolioSpec = "0 5 10,22 * * *"
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
	event := event.NewEvent(db, scraper, scraper)

	ch := make(chan string)

	chatId, err := strconv.ParseInt(conf.Telegram.ChatId, 10, 64)
	if err != nil {
		panic(err)
	}

	teleBot, err := bot.NewTeleBot(conf.Telegram.Token, chatId)
	if err != nil {
		panic(err)
	}

	c := cron.New()
	c.AddFunc(Every15Min, func() { event.AssetEvent(ch) })
	c.AddFunc(Every15Min, func() { event.RealEstateEvent(ch) })
	c.AddFunc(Every9Am, func() { event.IndexEvent(ch) })
	c.AddFunc(PortfolioSpec, func() { event.PortfolioEvent(ch) })
	c.Start()

	// go func() {
	// 	for true {
	// 		event.AssetEvent(ch)
	// 		time.Sleep(10 * time.Minute)
	// 	}
	// }()

	// go func() {
	// 	for true {
	// 		event.RealEstateEvent(ch)
	// 		time.Sleep(10 * time.Minute)
	// 	}
	// }()

	go func() {
		app.Run(db, scraper)
	}()

	for true {
		msg := <-ch
		teleBot.SendMessage(msg)
		log.Println(msg)
	}
}
