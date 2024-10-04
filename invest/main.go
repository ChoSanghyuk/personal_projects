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
	AssetSpec = "0 */15 * * * *" // todo. 우선 코인때문에 주말에도 로직 실행. 이후 분리?
	IndexSpec = "0 0 9 * * *"    // todo. 주말 제외 및 월요일일때는 금요일과 index 비교
	EmaSpec   = "0 0 9 * * 2-6"  // 화~토
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
	c.AddFunc(AssetSpec, func() { event.AssetEvent(ch) })
	c.AddFunc(AssetSpec, func() { event.RealEstateEvent(ch) })
	c.AddFunc(IndexSpec, func() { event.IndexEvent(ch) })
	c.AddFunc(EmaSpec, func() { event.EmaUpdateEvent(ch) })
	c.Start()

	go func() {
		app.Run(db, scraper)
	}()

	for true {
		msg := <-ch
		teleBot.SendMessage(msg)
		log.Println(msg)
	}
}
