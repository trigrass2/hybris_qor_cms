package db

import "github.com/jinzhu/gorm"

// master data
type Page struct {
	gorm.Model
	Url          string
	Section1     string `sql:"size:10000"`
	ProductCodes string
}

type Product struct {
	Code        string
	Description string
	Detail      string
	Image       string
}
