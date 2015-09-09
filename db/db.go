package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/validations"
	"github.com/theplant/hybris_qor_cms/config"
)

var DB gorm.DB
var AllModels = []interface{}{
	&Page{},
	&admin.AssetManager{},
}

func init() {
	host := config.MySQLHost
	database := config.MySQLDatabase

	var err error
	connect := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=true", config.MySQLRootPassword, host, config.MySQLPort, database)
	// fmt.Println(connect)
	DB, err = gorm.Open("mysql", connect)

	if err != nil {
		panic(fmt.Sprintf("Error when connect database: '%v'", err))
	}

	DB.LogMode(config.Verbose)

	models := AllModels

	for _, m := range models {
		err = DB.AutoMigrate(m).Error
		if err != nil {
			panic(fmt.Sprintf("Error when migrate: '%v'", err))
		}
	}

	validations.RegisterCallbacks(&DB)
}
