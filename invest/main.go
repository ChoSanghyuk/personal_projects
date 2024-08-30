package main

import (
	"invest/alarm/telegram"
	"invest/config"
	"invest/event"
	"invest/scrape"
	"strconv"

	"log"
)

func main() {
	// Create a new instance of the server

	scraper := scrape.Scraper{}
	conf := &config.ConfigInfo

	bh := event.BitcoinEventHandler{
		Scraper: &scrape.Scraper{
			ScrapeOption: scrape.BitcoinApi(conf.Bitcoin.API.Url),
		},
		UpperBound: conf.Bitcoin.Bound.Upper,
		LowerBound: conf.Bitcoin.Bound.Lower,
	}

	rh := event.RealEstateEventHandler{
		Scraper: &scraper,
		Url:     conf.RealEstate.Crawl.Url,
		Csspath: conf.RealEstate.Crawl.CssPath,
	}

	c := make(chan string)
	// get bit price
	go bh.PullingEvent(c)
	go rh.PullingEvent(c)

	chatId, err := strconv.ParseInt(conf.Telegram.ChatId, 10, 64)
	if err != nil {
		panic(err)
	}

	teleBot, err := telegram.NewTeleBot(conf.Telegram.Token, chatId)
	if err != nil {
		panic(err)
	}

	for true {
		msg := <-c
		teleBot.SendMessage(msg)
		log.Println(msg)
	}
}
