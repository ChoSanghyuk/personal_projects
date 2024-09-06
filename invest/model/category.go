package model

type Category uint

const (
	Cash Category = iota + 1
	Gold
	ShortTermBond
	DomesticStock
	DomesticCoin
)

var categoryList = []string{"현금", "금", "단기채권", "국내주식", "국내코인"}

func (c Category) String() string {
	return categoryList[c-1]
}

func ToCategory(s string) Category {

	for i, c := range categoryList {
		if s == c {
			return Category(i + 1)
		}
	}
	return 0
}
