package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/theplant/hybris_qor_cms/db"

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

	ctx.HTML(
		http.StatusOK,
		"page.tmpl",
		gin.H{
			"page": page,
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

	ctx.HTML(
		http.StatusOK,
		"page.tmpl",
		gin.H{
			"page": page,
			"raw": func(html string) template.HTML {
				return template.HTML(html)
			},
		},
	)
}
