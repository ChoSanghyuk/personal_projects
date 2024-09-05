package event

import (
	"fmt"
	m "invest/model"
	"log"
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

func (e Event) AssetEvent(c chan<- string) {
	// 현재 시장 단계 조회
	market, _ := e.stg.RetrieveMarketStatus("") // TODO. 에러 시 처리
	marketLevel := m.MarketLevel(market.Status)

	// 환율까지 계산하여 원화로 변환
	ex := e.scraper.ExchageRate()

	// 보유 자산 목록 조회
	assetList, _ := e.stg.RetrieveAssetList()
	assets := make([]*m.Asset, len(assetList))
	priceMap := make(map[uint]float64)

	for i := range len(assetList) {

		// 자산 정보 조회
		a, _ := e.stg.RetrieveAsset(assetList[i].ID)
		assets[i] = a

		// 자산별 현재 가격 조회
		cp, _ := e.scraper.CurrentPrice(a.Name)
		// url, header := e.tm.ApiInfo(a.Name)
		// p, _ := e.scraper.CallApi(url, header)                          // TODO. 조회 메소드 갱신 필요
		// cp, _ := strconv.ParseFloat(strings.ReplaceAll(p, ",", ""), 64) // TODO , 없애는 롤 누구 소유인지 판단
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
	investSummary, _ := e.stg.RetreiveFundsSummaryOrderByFundId()
	if len(investSummary) == 0 {
		return
	}

	length := investSummary[len(investSummary)-1].FundID + 1
	stable := make([]float64, length)
	volatile := make([]float64, length)

	for _, s := range investSummary {
		s.Sum = priceMap[s.AssetID] * float64(s.Count)
		e.stg.UpdateInvestSummarySum(s.FundID, s.AssetID, s.Sum)

		var v float64
		if s.Asset.Currency == m.USD.String() {
			v = s.Sum * ex
		} else {
			v = s.Sum
		}

		if s.Asset.Category <= 3 { // TODO. 단계 변수화
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
		}
	}
}

func (e Event) RealEstateEvent(c chan<- string) {

	// url, cssPath := e.tm.CrawlInfo("estate")

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
