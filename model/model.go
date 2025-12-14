package model

import "time"

type Categories struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Items struct {
	Id                int       `json:"id"`
	IdCategory        int       `json:"id_category"`
	Category          string    `json:"category"`
	Name              string    `json:"name"`
	Price             float32   `json:"price"`
	PurchaseDate      time.Time `json:"purchase_date"`
	Replaced          bool      `json:"replaced"`
	TotalUsage        int       `json:"total_usage"`
	AgeItem           int       `json:"age_item"`
	CurrentInvestment float32   `json:"current_investment"`
	DepreciationValue float32   `json:"depreciation_value"`
}

type TotalInvest struct {
	TotalItem          int     `json:"total_item"`
	TotalPurchasePrice int     `json:"total_purchase_price"`
	TotalInvest        float32 `json:"total_invest"`
	TotalDepreciation  float32 `json:"total_depreciation"`
}
