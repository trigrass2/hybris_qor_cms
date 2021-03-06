package db

import "github.com/jinzhu/gorm"

// master data
type Page struct {
	gorm.Model
	Url          string
	Section1     string     `sql:"size:10000"`
	ProductCodes string     `sql:"size:10000"`
	Products     []*Product `gorm:"many2many:page_products"`
}

type Product struct {
	gorm.Model
	Code          string `json:"@code"`
	Name          string
	Description   string `sql:"size:10000"`
	Picture       *Picture
	Price         *Price
	Europe1Prices *Europe1Price
}

type Europe1Price struct {
	PriceRow []*PriceRow
}
type PriceRow struct {
	Uri string `json:"@uri"`
}

type Picture struct {
	Code        string `json:"@code"`
	DownloadURL string `json:"@downloadURL"`
}

type Price struct {
	Price    string
	Currency *Currency
}

type Currency struct {
	Isocode string `json:"@isocode"`
}
