package models

import "time"

type ConfigInfo struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

type Product struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	LargeDescription string    `json:"large_description"`
	RegularPrice     float32   `json:"regular_price"`
	SalesPrice       float32   `json:"sales_price"`
	Color            string    `json:"color"`
	Size             string    `json:"size"`
	Stock            int       `json:"stock"`
	Images           string    `json:"images"`
	Code             string    `json:"code"`
	SKU              string    `json:"sku"`
	Category         string    `json:"category"`
	Gender           string    `json:"gender"`
	Published        bool      `json:"published"`
	Active           bool      `json:"active"`
	CreationDate     time.Time `json:"creation_date"`
	UpdateDate       time.Time `json:"update_date"`
}
