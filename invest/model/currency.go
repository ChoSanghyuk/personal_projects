package model

import "slices"

type Currency uint

const (
	WON Currency = iota
	USD
)

var currencyList = []string{"WON", "USD"}

func (c Currency) String() string {
	return currencyList[c]
}

func IsCurrency(t string) bool {
	return slices.Contains(currencyList, t)
}
