package db

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/validations"
)

func (device Device) Validate(db *gorm.DB) {
	var deviceInDb Device
	db.Where("code = ?", device.Code).First(&deviceInDb)

	if deviceInDb.ID != 0 && deviceInDb.ID != device.ID {
		db.AddError(validations.NewError(device, "Code", "代码已经存在了，不能重复"))
	}

	if device.Name == "" {
		db.AddError(validations.NewError(device, "Name", "设备名称不能为空"))
	}
}

// func (cdIn *ClientDeviceIn) Validate(db *gorm.DB) {
// 	if len(strings.TrimSpace(cdIn.DeviceName)) == 0 {
// 		db.AddError(validations.NewError(cdIn, "DeviceName", "收入设备名不能为空"))
// 	}
// 	if len(strings.TrimSpace(cdIn.ClientName)) == 0 {
// 		db.AddError(validations.NewError(cdIn, "ClientName", "收入客户名不能为空"))
// 	}
// 	if cdIn.Warehouse.ID == 0 {
// 		db.AddError(validations.NewError(cdIn, "Wharehouse", "收入到的仓库不能为空"))
// 	}
// 	if cdIn.Quantity <= 0 {
// 		db.AddError(validations.NewError(cdIn, "Quantity", "收入设备的数量要大于0"))
// 	}
// 	if cdIn.ByWhom.ID == 0 {
// 		db.AddError(validations.NewError(cdIn, "ByWhom", "请选择操作员"))
// 	}
// }

// func (dOut *DeviceOut) Validate(db *gorm.DB) {
// 	if dOut.FromReportItemID == 0 {
// 		db.AddError(validations.NewError(dOut, "FromReportItemID", "带出设备不能为空"))
// 	}
// 	if dOut.ToWhomID == 0 {
// 		db.AddError(validations.NewError(dOut, "ToWhomID", "设备带出人不能为空"))
// 	}
// 	if dOut.Quantity <= 0 {
// 		db.AddError(validations.NewError(dOut, "Quantity", "带出设备的数量要大于0"))
// 	}
// 	if dOut.ByWhomID == 0 {
// 		db.AddError(validations.NewError(dOut, "ByWhomID", "请选择操作员"))
// 	}

// 	from, to, d, err := fromToDevice(dOut.FromReportItemID, dOut.ToWhomID, "Employee")
// 	if err != nil {
// 		db.AddError(validations.NewError(dOut, "FromReportItemID", err.Error()))
// 		return
// 	}

// 	dOut.DeviceName = d.Name
// 	dOut.ToWhomName = to.HolderName()
// 	dOut.FromWarehouseName = from.HolderName()

// 	byWhom, _ := holderByIDType(dOut.ByWhomID, "Employee")
// 	dOut.ByWhomName = byWhom.HolderName()

// }

// func (dIn *DeviceIn) Validate(db *gorm.DB) {
// 	if dIn.FromReportItemID == 0 {
// 		db.AddError(validations.NewError(dIn, "FromReportItemID", "带出设备不能为空"))
// 	}
// 	if dIn.ToWarehouseID == 0 {
// 		db.AddError(validations.NewError(dIn, "ToWhomID", "设备还回仓库不能为空"))
// 	}
// 	if dIn.Quantity <= 0 {
// 		db.AddError(validations.NewError(dIn, "Quantity", "还回设备的数量要大于0"))
// 	}
// 	if dIn.ByWhomID == 0 {
// 		db.AddError(validations.NewError(dIn, "ByWhomID", "请选择操作员"))
// 	}

// 	from, to, d, err := fromToDevice(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse")
// 	if err != nil {
// 		db.AddError(validations.NewError(dIn, "FromReportItemID", err.Error()))
// 		return
// 	}

// 	dIn.DeviceName = d.Name
// 	dIn.ToWarehouseName = to.HolderName()
// 	dIn.FromWhomName = from.HolderName()

// 	byWhom, _ := holderByIDType(dIn.ByWhomID, "Employee")
// 	dIn.ByWhomName = byWhom.HolderName()

// }

// func (cOut *ConsumableOut) Validate(db *gorm.DB) {
// 	if cOut.ReportItemID == 0 {
// 		db.AddError(validations.NewError(cOut, "ReportItemID", "消耗品不能为空"))
// 	}
// 	if cOut.ToWhomID == 0 {
// 		db.AddError(validations.NewError(cOut, "ToWhomID", "消耗品使用人不能为空"))
// 	}
// 	if cOut.Quantity <= 0 {
// 		db.AddError(validations.NewError(cOut, "Quantity", "带出消耗品的数量要大于0"))
// 	}
// 	if cOut.ByWhomID == 0 {
// 		db.AddError(validations.NewError(cOut, "ByWhomID", "请选择操作员"))
// 	}

// 	from, to, d, err := fromToDevice(cOut.ReportItemID, cOut.ToWhomID, "Employee")
// 	if err != nil {
// 		db.AddError(validations.NewError(cOut, "FromReportItemID", err.Error()))
// 		return
// 	}

// 	cOut.DeviceName = d.Name
// 	cOut.ToWhomName = to.HolderName()
// 	cOut.WarehouseName = from.HolderName()

// 	byWhom, _ := holderByIDType(cOut.ByWhomID, "Employee")
// 	cOut.ByWhomName = byWhom.HolderName()

// }

// func (cIn *ConsumableIn) Validate(db *gorm.DB) {
// 	if cIn.ReportItemID == 0 {
// 		db.AddError(validations.NewError(cIn, "ReportItemID", "购买的消耗品不能为空"))
// 	}
// 	if cIn.Quantity <= 0 {
// 		db.AddError(validations.NewError(cIn, "Quantity", "购买的数量要大于0"))
// 	}
// 	if cIn.ByWhomID == 0 {
// 		db.AddError(validations.NewError(cIn, "ByWhomID", "请选择操作员"))
// 	}

// 	from, _, d, err := fromToDevice(cIn.ReportItemID, 0, "Employee")
// 	if err != nil {
// 		db.AddError(validations.NewError(cIn, "ReportItemID", err.Error()))
// 		return
// 	}

// 	cIn.DeviceName = d.Name
// 	cIn.WarehouseName = from.HolderName()
// 	byWhom, _ := holderByIDType(cIn.ByWhomID, "Employee")
// 	cIn.ByWhomName = byWhom.HolderName()

// }

// func (cdOut *ClientDeviceOut) Validate(db *gorm.DB) {

// 	if cdOut.ClientDeviceInID == 0 {
// 		db.AddError(validations.NewError(cdOut, "ClientDeviceInID", "选择收入过的客户设备"))
// 	}
// 	if cdOut.ByWhom.ID == 0 {
// 		db.AddError(validations.NewError(cdOut, "ByWhom", "请选择操作员"))
// 	}
// 	// panic(fmt.(cdOut.ClientDeviceInID))
// 	if cdOut.ClientDeviceInID > 0 {
// 		cdIn := &ClientDeviceIn{}
// 		DB.Preload("Warehouse").Find(cdIn, cdOut.ClientDeviceInID)
// 		// log.Println(cdIn)
// 		// panic("stop")

// 		cdOut.DeviceName = cdIn.DeviceName
// 		cdOut.ClientName = cdIn.ClientName
// 		cdOut.Quantity = cdIn.Quantity
// 		cdOut.WarehouseName = cdIn.Warehouse.Name
// 	}
// }
