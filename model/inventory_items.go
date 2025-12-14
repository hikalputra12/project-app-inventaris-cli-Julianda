package model

import "time"

/* Inventory represents an inventory item. */
type Inventory struct {
	Model
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
}

// total barang  yang di pakai lebih dari 100 hari
type TotalUsageDays struct {
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	TotalUsageDays int     `json:"total_usage_days"`
}

type TotalInvestmentValue struct {
	TotalInvestmentValue float64 `json:"total_investment_value"`
}

type Depreciation struct {
	Name            string
	InitialPrice    float64
	InvestmentValue float64
	Depreciation    float64
}
