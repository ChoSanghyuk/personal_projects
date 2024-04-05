package util

import "log"

func CheckErr(e error, msg ...string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}
