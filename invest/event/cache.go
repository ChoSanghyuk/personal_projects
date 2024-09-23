package event

import "time"

type assetMsg struct {
	assetId uint
	isSell  bool
}

type assetMsgSentInfo struct {
	isMsgSent bool
	price     float64
	sentTime  time.Time
}

var assetMsgCache map[assetMsg]*assetMsgSentInfo

func init() {
	assetMsgCache = make(map[assetMsg]*assetMsgSentInfo)
}

func hasMsgCache(assetId uint, isSell bool, price float64) bool {

	cache := assetMsgCache[assetMsg{
		assetId: assetId,
		isSell:  isSell,
	}]

	if cache == nil { // || cache.sentTime.IsZero() sentTime이 미존재할 경우는 없음
		return false
	}

	diff := time.Since(cache.sentTime)
	if diff <= 24*time.Hour { // 유효한 캐시
		return cache.price == price && cache.isMsgSent
	}

	return false
}

func setMsgCache(assetId uint, isSell bool, price float64) {

	k := assetMsg{
		assetId: assetId,
		isSell:  isSell,
	}

	cache := assetMsgCache[k]

	if cache == nil {
		assetMsgCache[k] = &assetMsgSentInfo{
			isMsgSent: true,
			price:     price,
			sentTime:  time.Now(),
		}
	} else {
		cache.price = price
		cache.sentTime = time.Now()
	}

}
