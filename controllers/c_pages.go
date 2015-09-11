package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/theplant/hybris_qor_cms/config"
	"github.com/theplant/hybris_qor_cms/db"
	"github.com/theplant/hybris_qor_cms/hybris"

	"github.com/gin-gonic/gin"
)

func ViewPageHandler(ctx *gin.Context) {
	page := &db.Page{}
	url := ctx.Params.ByName("url")
	if err := db.DB.Where("url = ?", "/"+url).Preload("Products").Find(&page).Error; err != nil {
		panic(err)
	}
	// log.Printf("page %+v", page)
	// if err := DB.Model(&book).Related(&book.Authors, "Authors").Error; err != nil {
	// 	panic(err)
	// }

	if strings.TrimSpace(page.ProductCodes) != "" {
		codes := strings.Split(strings.TrimSpace(page.ProductCodes), ",")
		for _, code := range codes {
			p := hybris.GetProduct(code)
			if p == nil {
				continue
			}
			ps := *p
			db.DB.Where(db.Product{Code: p.Code}).Assign(ps).FirstOrCreate(&ps)
		}
	}
	var products []*db.Product
	for _, p := range page.Products {
		p := hybris.GetProduct(p.Code)
		if p.Europe1Prices != nil && len(p.Europe1Prices.PriceRow) == 2 {
			p.Price = hybris.GetPrice(p.Europe1Prices.PriceRow[1].Uri)
		}

		// log.Printf("product %+v", p)
		// log.Printf("Picture %+v", p.Picture)
		// log.Printf("Picture %+v", p.Europe1Prices.PriceRow[1].Uri)
		// log.Printf("Picture %+v", p.Price)
		// log.Printf("Picture %+v", p.Price.Currency
		products = append(products, p)
	}

	ctx.HTML(
		http.StatusOK,
		"page.tmpl",
		gin.H{
			"page":     page,
			"products": products,
			"host":     template.URL(config.Host),
			"raw": func(html string) template.HTML {
				return template.HTML(html)
			},
		},
	)
}
