package model

import "time"

type FlipTrack struct {
	FoundTime time.Time `json:"foundTime"`
	Sell      struct {
		Uuid      string  `json:"uuid"`
		SellPrice float64 `json:"sellPrice"`
	} `json:"sell"`
	OriginAuction string `json:"originAuction"`
	Seller        string `json:"seller"`
}
