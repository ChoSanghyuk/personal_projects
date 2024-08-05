package main

import (
	"invest/alarm/mail"
	"invest/config"
	"invest/event"
	"invest/scrape"

	"log"
)

func main() {
	// Create a new instance of the server

	scraper := scrape.Scraper{}
	conf := &config.ConfigInfo

	bh := event.BitcoinEventHandler{
		Scraper:    &scraper,
		Url:        conf.Bitcoin.Crawl.Url,
		Csspath:    conf.Bitcoin.Crawl.CssPath,
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

	for true {
		msg := <-c
		mail.SendEmail(conf.Email.SMTP, msg, conf.Email.Target, msg)
		log.Println(msg)
	}
}
