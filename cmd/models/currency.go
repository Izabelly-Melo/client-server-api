package models

type CurrencyInfo struct {
	Id         int     `gorm:"primaryKey"`
	Code       string  `json:"code"`
	Codein     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high,string"`
	Low        float64 `json:"low,string"`
	VarBid     float64 `json:"varBid,string"`
	PctChange  float64 `json:"pctChange,string"`
	Bid        float64 `json:"bid,string"`
	Ask        float64 `json:"ask,string"`
	Timestamp  int64   `json:"timestamp,string"`
	CreateDate string  `json:"create_date"`
}

type CurrencyResponse struct {
	USDBRL CurrencyInfo `json:"USDBRL"`
}
