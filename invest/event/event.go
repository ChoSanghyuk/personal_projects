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

/*
작업 1. 자산의 현재가와 자산의 매도/매수 기준 비교하여 알림 전송
 - 보유 자산 list
 - 자산 정보
 - 현재가
작업 2. 자금별/종목별 현재 총액 갱신 + 최저가/최고가 갱신
 - investSummary list
 - 현재가
 - 환율
 - 자산 정보
직업 3. 현재 시장 단계에 맞는 변동 자산을 가지고 있는지 확인하여 알림 전송. 대상 시, 우선처분 대상 및 보유 자산 현환 전송
 - 시장 단계
 - 갱신된 investSummary list
*/

func (e Event) buySellMsg(assetId uint, pm map[uint]float64) (msg string, err error) {

	// 자산 정보 조회
	a, err := e.stg.RetrieveAsset(assetId)
	if err != nil {
		return "", fmt.Errorf("[AssetEvent] RetrieveAsset 시, 에러 발생. %w", err)
	}

	// 자산별 현재 가격 조회
	category, err := m.ToCategory(a.Category)
	if err != nil {
		return "", fmt.Errorf("[AssetEvent] ToCategory시, 에러 발생. %w", err)
	}
	cp, err := e.scraper.CurrentPrice(category, a.Code)
	if err != nil {
		return "", fmt.Errorf("[AssetEvent] CurrentPrice 시, 에러 발생. %w", err)
	}

	log.Printf("%s 현재 가격 %.3f", a.Name, cp)

	pm[assetId] = cp

	// 자산 매도/매수 기준 비교 및 알림 여부 판단. (알림 전송)
	if a.BuyPrice >= cp {
		msg = fmt.Sprintf("BUY %s. LOWER BOUND : %f. CURRENT PRICE :%f", a.Name, a.BuyPrice, cp)
	} else if a.SellPrice <= cp {
		msg = fmt.Sprintf("SELL %s. UPPER BOUND : %f. CURRENT PRICE :%f", a.Name, a.SellPrice, cp)
	}

	return
}

func (e Event) updateFundSummarys(list []m.InvestSummary, pm map[uint]float64) (err error) {
	for i := range len(list) {
		is := &list[i]
		is.Sum = pm[is.AssetID] * float64(is.Count)

		err = e.stg.UpdateInvestSummarySum(is.FundID, is.AssetID, is.Sum)
		if err != nil {
			return
		}
	}
	return nil
}

func (e Event) portfolioMsg(ivsmLi []m.InvestSummary, pm map[uint]float64) (msg string, err error) {
	// 현재 시장 단계 조회
	market, err := e.stg.RetrieveMarketStatus("")
	if err != nil {
		msg = fmt.Sprintf("[AssetEvent] RetrieveMarketStatus 시, 에러 발생. %s", err)
		return
	}
	marketLevel := m.MarketLevel(market.Status)

	// 환율까지 계산하여 원화로 변환
	ex := e.scraper.ExchageRate()
	if ex == 0 {
		msg = "[AssetEvent] ExchageRate 시 환율 값 0 반환"
		return
	}

	keySet := make(map[uint]bool)
	stable := make(map[uint]float64)
	volatile := make(map[uint]float64)
	priority := make(map[uint][]m.InvestSummary)
	extra := make(map[uint][]m.InvestSummary)

	for i := range len(ivsmLi) {

		ivsm := &ivsmLi[i]

		keySet[ivsm.FundID] = true

		var v float64
		if ivsm.Asset.Currency == m.USD.String() {
			v = ivsm.Sum * ex
		} else {
			v = ivsm.Sum
		}

		category, err := m.ToCategory(ivsm.Asset.Category)
		if err != nil {
			msg = fmt.Sprintf("[AssetEvent] investSummary loop내 ToCategory시, 에러 발생. %s", err)
			return msg, err
		}

		if category.IsStable() {
			stable[ivsm.FundID] = stable[ivsm.FundID] + v
		} else {
			volatile[ivsm.FundID] = volatile[ivsm.FundID] + v
		}

		if !category.IsStable() && ivsm.Asset.Top <= pm[ivsm.AssetID] {
			priority[ivsm.FundID] = append(priority[ivsm.FundID], *ivsm)
		} else {
			extra[ivsm.FundID] = append(extra[ivsm.FundID], *ivsm)
		}
	}

	var sb strings.Builder
	for k := range keySet {
		if volatile[k]+stable[k] == 0 {
			continue
		}

		r := volatile[k] / (volatile[k] + stable[k])
		if r > marketLevel.VolatileAssetRate() {
			sb.WriteString(strings.Repeat("=", 20))
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("자금 %d 변동 자산 비중 초과. 변동 자산 비율 : %.2f. 현재 시장 단계 : %s(%.1f)\n\n", k, r, marketLevel.String(), marketLevel.VolatileAssetRate()))

			sb.WriteString("=====우선 처분 종목 정보=====\n")
			for _, p := range priority[k] {
				sb.WriteString(fmt.Sprintf("종목 %d %s %.2f %.2f\n", p.Asset.ID, p.Asset.Name, p.Asset.Top, pm[p.AssetID]))
			}
			sb.WriteString("=====보유 종목 정보=====\n")
			for _, e := range extra[k] {
				sb.WriteString(fmt.Sprintf("종목 %d %s %.2f %.2f\n", e.Asset.ID, e.Asset.Name, e.Asset.Top, pm[e.AssetID]))
			}

		}

	}

	msg = sb.String()
	return
}

func (e Event) AssetEvent(c chan<- string) {

	// 등록 자산 목록 조회
	assetList, err := e.stg.RetrieveAssetList()
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] RetrieveAssetList 시, 에러 발생. %s", err)
		return
	}
	priceMap := make(map[uint]float64) // assetId => price

	// 등록 자산 매수/매도 기준 충족 시, 채널로 메시지 전달
	for _, a := range assetList {
		msg, err := e.buySellMsg(a.ID, priceMap)
		if err != nil {
			c <- fmt.Sprintf("[AssetEvent] buySellMsg시, 에러 발생. %s", err)
			return
		}
		if msg != "" {
			c <- msg
		}
	}

	// 자금별 종목 투자 내역 조회
	ivsmLi, err := e.stg.RetreiveFundsSummaryOrderByFundId()
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] RetreiveFundsSummaryOrderByFundId 시, 에러 발생. %s", err)
		return
	}
	if len(ivsmLi) == 0 {
		return
	}

	// 자금별/종목별 현재 총액 갱신
	err = e.updateFundSummarys(ivsmLi, priceMap)
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] updateFundSummary 시, 에러 발생. %s", err)
		return
	}

	// 현재 시장 단계 이하로 변동 자산을 가지고 있는지 확인. (알림 전송)
	msg, err := e.portfolioMsg(ivsmLi, priceMap)
	if err != nil {
		c <- fmt.Sprintf("[AssetEvent] portfolioMsg시, 에러 발생. %s", err)
	}
	if msg != "" {
		c <- msg
	}
}

func (e Event) RealEstateEvent(c chan<- string) {

	rtn, err := e.scraper.RealEstateStatus()
	if err != nil {
		c <- fmt.Sprintf("크롤링 시 오류 발생. %s", err.Error())
		return
	}

	if rtn != "예정지구 지정" {
		c <- fmt.Sprintf("연신내 재개발 변동 사항 존재. 예정지구 지정 => %s", rtn)
	} else {
		log.Printf("연신내 변동 사항 없음. 현재 단계: %s", rtn)
	}
}
