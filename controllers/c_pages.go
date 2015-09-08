package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/theplant/hybris_qor_cms/db"
	"github.com/theplant/hybris_qor_cms/hybris"

	"github.com/gin-gonic/gin"
)

func ViewPageHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var page = &db.Page{}
	if err := db.DB.Find(&page, id).Error; err != nil {
		panic(err)
	}

	// if err := DB.Model(&book).Related(&book.Authors, "Authors").Error; err != nil {
	// 	panic(err)
	// }
	codes := strings.Split(strings.TrimSpace(page.ProductCodes), ",")
	var products []*db.Product
	for _, code := range codes {
		products = append(products, hybris.GetProduct(code))
	}

	log.Printf("product %+v", products)

	ctx.HTML(
		http.StatusOK,
		"page.tmpl",
		gin.H{
			"page":     page,
			"products": products,
			"raw": func(html string) template.HTML {
				return template.HTML(html)
			},
		},
	)
}

func ViewPageHandler2(ctx *gin.Context) {
	page := &db.Page{}
	if err := db.DB.Where("url = ?", strings.TrimPrefix(ctx.Request.URL.RequestURI(), "/")).Find(&page).Error; err != nil {
		panic(err)
	}

	// if err := DB.Model(&book).Related(&book.Authors, "Authors").Error; err != nil {
	// 	panic(err)
	// }

	codes := strings.Split(strings.TrimSpace(page.ProductCodes), ",")
	var products []*db.Product
	for _, code := range codes {
		p := hybris.GetProduct(code)
		p.Price = hybris.GetPrice(p.Europe1Prices.PriceRow[1].Uri)
		// log.Printf("product %+v", p)
		// log.Printf("Picture %+v", p.Picture)
		// log.Printf("Picture %+v", p.Europe1Prices.PriceRow[1].Uri)
		// log.Printf("Picture %+v", p.Price)
		// log.Printf("Picture %+v", p.Price.Currency)
		products = append(products, p)
	}

	ctx.HTML(
		http.StatusOK,
		"page.tmpl",
		gin.H{
			"page":     page,
			"products": products,
			"raw": func(html string) template.HTML {
				return template.HTML(html)
			},
		},
	)
}
