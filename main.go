package main

import (
	"log"
	"net/http"

	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/i18n"
	"github.com/qor/qor/i18n/backends/database"
	"github.com/qor/qor/roles"
	"github.com/theplant/device_management/db"
)

func main() {
	roles.Register("admin", func(req *http.Request, cu qor.CurrentUser) bool {
		return true
	})

	adm := admin.New(&qor.Config{DB: &db.DB})

	adm.SetAuth(&Auth{})
	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"数据维护"}})
	device.Meta(&admin.Meta{Name: "CategoryID", Type: "select_one", Collection: db.DeviceCategories})
	device.Meta(&admin.Meta{Name: "WarehouseID", Type: "select_one", Collection: db.WarehouseCollection})
	device.EditAttrs("Name", "Code", "TotalQuantity")
	device.IndexAttrs("Name", "Code", "TotalQuantity")

	I18nBackend := database.New(&db.DB)
	// config.I18n = i18n.New(I18nBackend)
	adm.AddResource(i18n.New(I18nBackend), &admin.Config{Menu: []string{"系统设置"}, Invisible: true})

	adm.MountTo("/admin", http.DefaultServeMux)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/admin/report_items", 302)
	// })

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/admin/report_items"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/admin/report_items"
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
	return "管理员"
}
