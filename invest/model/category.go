package model

type Category uint

const (
	Cash Category = iota + 1
	Gold
	ShortTermBond
	DomesticStock
	DomesticCoin
)

var CategoryList = []string{"현금", "금", "단기채권", "국내주식", "국내코인"}

func (c Category) String() string {
	return CategoryList[c-1]
}
