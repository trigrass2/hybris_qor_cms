package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/i18n"
	"github.com/qor/qor/i18n/backends/database"
	"github.com/qor/qor/roles"
	"github.com/theplant/hybris_qor_cms/config"
	"github.com/theplant/hybris_qor_cms/controllers"
	"github.com/theplant/hybris_qor_cms/db"
)

func main() {
	roles.Register("admin", func(req *http.Request, cu qor.CurrentUser) bool {
		return true
	})
	mux := http.NewServeMux()
	adm := admin.New(&qor.Config{DB: &db.DB})

	adm.SetAuth(&Auth{})
	page := adm.AddResource(&db.Page{}, &admin.Config{Menu: []string{"Cms"}})
	// adm.AddResource(&db.Product{}, &admin.Config{Menu: []string{"Cms"}})
	page.Meta(&admin.Meta{
		Name:     "Section1",
		Type:     "rich_editor",
		Resource: adm.AddResource(&admin.AssetManager{}, &admin.Config{Invisible: true}),
	})

	page.Meta(&admin.Meta{Name: "Url", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		url := resource.(*db.Page).Url
		url = fmt.Sprintf(`<a href="%s/page%s">%s</a>`, config.CmsHost, url, url)
		return template.HTML(url)
	}})

	page.NewAttrs("-ProductCodes")
	page.EditAttrs("-ProductCodes")
	page.IndexAttrs("-ProductCodes")
	I18nBackend := database.New(&db.DB)
	// config.I18n = i18n.New(I18nBackend)
	adm.AddResource(i18n.New(I18nBackend), &admin.Config{Menu: []string{"系统设置"}, Invisible: true})

	adm.MountTo("/admin", mux)
	// frontend routes
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	bindPages(router)
	router.StaticFS("/system/", http.Dir("public/system"))
	// books
	pageRoutes := router.Group("/page")
	{
		// listing
		// pageRoutes.GET("", controllers.ListBooksHandler)
		// pageRoutes.GET("/", controllers.ListBooksHandler)
		pageRoutes.GET("/:url", controllers.ViewPageHandler)
	}
	mux.Handle("/", router)

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func bindPages(router *gin.Engine) {
	pages := []db.Page{}
	db.DB.Find(&pages)
	for _, page := range pages {
		if !strings.HasPrefix(page.Url, "/") {
			page.Url = "/" + page.Url
		}
		router.GET(page.Url, controllers.ViewPageHandler)
	}
}

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/admin/pages"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/admin/pages"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {

	return &User{}
}

type User struct {
}

func (u User) AvailableLocales() []string {
	return []string{"zh_CN"}
}

func (u User) DisplayName() string {
	return "Admin"
}
