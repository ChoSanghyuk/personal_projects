package main

import (
	"fmt"
	"invest/alarm/mail"
	"invest/config"
	"invest/web_scrape/crawl"
	"strconv"
	"strings"
	"time"

	"log"
)

func main() {
	// Create a new instance of the server

	conf := &config.ConfigInfo

	for true {

		// get bit price

		rtn, err := crawl.Crawl(conf.Bitcoin.Crawl.Url, conf.Bitcoin.Crawl.CssPath)
		if err != nil {
			fmt.Println(err.Error())

		}

		bitPrice, err := strconv.ParseFloat(strings.ReplaceAll(rtn, ",", ""), 64)
		if err != nil {
			log.Print(err.Error())
		}

		var msg string
		if bitPrice >= conf.Bitcoin.Bound.Upper {
			msg = fmt.Sprintf("SELL BIT. UPPER BOUND : %f. CURRENT PRICE :%f", conf.Bitcoin.Bound.Upper, bitPrice)

		} else if bitPrice <= conf.Bitcoin.Bound.Lower {
			msg = fmt.Sprintf("BUY BIT. LOWER BOUND : %f. CURRENT PRICE :%f", conf.Bitcoin.Bound.Upper, bitPrice)
		} else {
			// msg = fmt.Sprintf("STAY. CURRENT PRICE : %f", bitPrice)
		}

		mail.SendEmail(&config.ConfigInfo.Email.SMTP, msg, config.ConfigInfo.Email.Target, msg)
		log.Println(msg)

		time.Sleep(time.Minute * 5)
	}

}
