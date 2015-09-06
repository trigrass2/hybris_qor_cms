package db

import "time"

// master data
type Device struct {
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"unique_index:idx_code_deleted_at"`
	Name          string
	Code          string `sql:"unique_index:idx_code_deleted_at"`
	TotalQuantity uint
	WarehouseID   uint
	CategoryID    uint
}

// type Warehouse struct {
// 	gorm.Model
// 	Name    string
// 	Address string
// }

// type Employee struct {
// 	gorm.Model
// 	Name   string
// 	Mobile string
// }

// // operations data
// type DeviceIn struct {
// 	gorm.Model
// 	FromReportItemID uint
// 	FromWhomName     string
// 	DeviceName       string
// 	Quantity         int
// 	ToWarehouseID    uint
// 	ToWarehouseName  string
// 	ByWhomID         uint
// 	ByWhomName       string
// 	Date             time.Time
// }

// type DeviceOut struct {
// 	gorm.Model
// 	FromReportItemID  uint
// 	FromWarehouseName string
// 	DeviceName        string
// 	Quantity          uint
// 	ToWhomID          uint
// 	ToWhomName        string
// 	ByWhomID          uint
// 	ByWhomName        string
// 	Date              time.Time
// }

// type ClientDeviceIn struct {
// 	gorm.Model
// 	DeviceName  string
// 	ClientName  string
// 	Quantity    int
// 	Date        time.Time
// 	WarehouseID uint
// 	Warehouse   Warehouse
// 	ByWhomID    uint
// 	ByWhom      Employee
// }

// type ClientDeviceOut struct {
// 	gorm.Model
// 	ClientDeviceInID uint
// 	DeviceName       string
// 	ClientName       string
// 	Quantity         int
// 	WarehouseName    string
// 	Date             time.Time
// 	ByWhomID         uint
// 	ByWhom           Employee
// }

// type ConsumableIn struct {
// 	gorm.Model
// 	ReportItemID  uint
// 	DeviceName    string
// 	Quantity      int
// 	WarehouseName string
// 	ByWhomID      uint
// 	ByWhomName    string
// 	Date          time.Time
// }

// type ConsumableOut struct {
// 	gorm.Model
// 	ReportItemID  uint
// 	DeviceName    string
// 	Quantity      int
// 	WarehouseName string
// 	ToWhomID      uint
// 	ToWhomName    string
// 	ByWhomID      uint
// 	ByWhomName    string
// 	Date          time.Time
// }

// // report data
// type ReportItem struct {
// 	ID                 uint `gorm:"primary_key"`
// 	CreatedAt          time.Time
// 	UpdatedAt          time.Time
// 	DeletedAt          *time.Time
// 	WhoHasThemName     string
// 	WhoHasThemID       uint
// 	WhoHasThemType     string
// 	ClientName         string
// 	DeviceID           uint
// 	DeviceName         string
// 	DeviceCode         string
// 	DeviceCategoryID   uint
// 	OperatedByWhomID   uint
// 	OperatedByWhomName string
// 	Count              int
// 	ClientDeviceInID   uint
// }
