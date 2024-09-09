package event

import (
	"fmt"
	m "invest/model"
	"log"
	"strings"
)

type Event struct {
	stg     Storage
	scraper Scraper
}

func NewEvent(stg Storage, scraper Scraper) *Event {
	return &Event{
		stg:     stg,
		scraper: scraper,
	}
}

// Todo. 작업 단위로 함수 분리
func (e Event) AssetEvent(c chan<- string) {
	// 현재 시장 단계 조회
	market, err := e.stg.RetrieveMarketStatus("")
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] RetrieveMarketStatus 시, 에러 발생. %s", err)
	}
	marketLevel := m.MarketLevel(market.Status)

	// 환율까지 계산하여 원화로 변환
	ex := e.scraper.ExchageRate()
	if ex == 0 {
		c <- "[AssetEvent] ExchageRate 시 환율 값 0 반환"
	}

	// 보유 자산 목록 조회
	assetList, err := e.stg.RetrieveAssetList()
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] RetrieveAssetList 시, 에러 발생. %s", err)
	}
	assets := make([]*m.Asset, len(assetList))
	priceMap := make(map[uint]float64) // assetId => price

	for i := range len(assetList) {

		// 자산 정보 조회
		a, err := e.stg.RetrieveAsset(assetList[i].ID)
		if err != nil {
			c <- fmt.Sprintf("[AssetEvent] RetrieveAsset 시, 에러 발생. %s", err)
		}
		assets[i] = a

		// 자산별 현재 가격 조회
		category, err := m.ToCategory(a.Category)
		if err != nil {
			c <- fmt.Sprintf("[AssetEvent] ToCategory시, 에러 발생. %s", err)
		}
		cp, err := e.scraper.CurrentPrice(category, a.Code)
		if err != nil {
			c <- fmt.Sprintf("[AssetEvent] CurrentPrice 시, 에러 발생. %s", err)
		}
		log.Printf("%s 현재 가격 %.3f", a.Name, cp)
		priceMap[a.ID] = cp

		// 자산 매도/매수 기준 비교 및 알림 여부 판단. (알림 전송)
		if a.BuyPrice >= cp {
			c <- fmt.Sprintf("BUY %s. LOWER BOUND : %f. CURRENT PRICE :%f", a.Name, a.BuyPrice, cp)
		} else if a.SellPrice <= cp {
			c <- fmt.Sprintf("SELL %s. UPPER BOUND : %f. CURRENT PRICE :%f", a.Name, a.SellPrice, cp)
		}
	}

	// 자금별/종목별 현재 총액 갱신
	investSummary, err := e.stg.RetreiveFundsSummaryOrderByFundId()
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] RetreiveFundsSummaryOrderByFundId 시, 에러 발생. %s", err)
	}
	if len(investSummary) == 0 {
		return
	}

	length := investSummary[len(investSummary)-1].FundID + 1
	stable := make([]float64, length)
	volatile := make([]float64, length)

	for _, s := range investSummary {
		s.Sum = priceMap[s.AssetID] * float64(s.Count) // todo. 이거 s 정보 바뀌는지 확인 필요
		e.stg.UpdateInvestSummarySum(s.FundID, s.AssetID, s.Sum)

		var v float64
		if s.Asset.Currency == m.USD.String() {
			v = s.Sum * ex
		} else {
			v = s.Sum
		}

		category, err := m.ToCategory(s.Asset.Category)
		if err != nil {
			c <- fmt.Sprintf("[AssetEvent] investSummary loop내 ToCategory시, 에러 발생. %s", err)
		}
		if category.IsStable() {
			stable[s.FundID] = stable[s.FundID] + v
		} else {
			volatile[s.FundID] = volatile[s.FundID] + v
		}
	}

	// 현재 시장 단계 이하로 변동 자산을 가지고 있는지 확인. (알림 전송)
	for i := range length {
		if volatile[i]+stable[i] == 0 {
			continue
		}

		r := volatile[i] / (volatile[i] + stable[i])
		if r > marketLevel.VolatileAssetRate() {
			c <- fmt.Sprintf("자금 %d 변동 자산 비중 초과. 변동 자산 비율 : %f. 현재 시장 단계 : %s", i, r, marketLevel.String())

			priority, extra := divideFundAsset(i, investSummary, priceMap)
			var sb strings.Builder
			sb.WriteString("=====우선 처분 자산 정보=====\n")
			for _, is := range priority {
				sb.WriteString(fmt.Sprintf("%+v\n", is))
			}
			sb.WriteString("=====그외 자산 정보=====\n")
			for _, is := range extra {
				sb.WriteString(fmt.Sprintf("%+v\n", is))
			}
			c <- sb.String()
		}
	}
}

func (e Event) RealEstateEvent(c chan<- string) {

	rtn, err := e.scraper.RealEstateStatus()
	if err != nil {
		c <- fmt.Sprintf("크롤링 시 오류 발생. %s", err.Error())
	}

	if rtn != "예정지구 지정" {
		c <- fmt.Sprintf("연신내 재개발 변동 사항 존재. 예정지구 지정 => %s", rtn)
	} else {
		log.Printf("연신내 변동 사항 없음. 현재 단계: %s", rtn)
	}
}

func divideFundAsset(id uint, base []m.InvestSummary, priceMap map[uint]float64) (priority []m.InvestSummary, extra []m.InvestSummary) {

	for _, is := range base {
		if is.FundID == id {
			// 해당 Asset의 현재가가 최고가 넘겼는지 확인.
			if priceMap[is.AssetID] >= is.Asset.Top {
				priority = append(priority, is)
			} else {
				extra = append(extra, is)
			}
		}
	}
	return
}
