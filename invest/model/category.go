package model

import "errors"

type Category uint

const (
	Won Category = iota + 1
	Dollar
	Gold
	ShortTermBond
	DomesticETF
	DomesticStock
	DomesticCoin
	Leverage
)

var categoryList = []string{"현금", "달러", "금", "단기채권", "국내ETF", "국내주식", "국내코인", "레버리지"}

func (c Category) String() string {
	return categoryList[c-1]
}

func ToCategory(s string) (Category, error) {

	for i, c := range categoryList {
		if s == c {
			return Category(i + 1), nil
		}
	}
	return 0, errors.New("존재하지 않는 카테고리 번호. 입력 값 :" + s)
}

func (c Category) IsStable() bool {
	if c <= 4 {
		return true
	} else {
		return false
	}
}

func CategoryLength() uint64 {
	return uint64(len(categoryList))
}
