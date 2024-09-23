package db

import "time"

type assetMsgCache struct {
	isMsgSent bool
	sentTime  time.Time
}

var assetIsMsgCache map[int]assetMsgCache

func IsAssetSent(assetId int) bool {

	return false
}
