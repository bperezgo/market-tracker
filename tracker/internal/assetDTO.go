package domain

import "time"

type AssetDTO struct {
	ID       string    `json:"id"`
	Date     time.Time `json:"date"`
	Exchange string    `json:"exchange"`
	Price    float32   `json:"price"`
}
