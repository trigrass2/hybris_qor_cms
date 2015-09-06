package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type DeviceHolder interface {
	HolderName() string
	HolderID() uint
	HolderType() string
}

func (e Employee) HolderName() string {
	return e.Name
}
func (e Employee) HolderID() uint {
	return e.ID
}
func (e Employee) HolderType() string {
	return "Employee"
}

func (e Warehouse) HolderName() string {
	return e.Name
}
func (e Warehouse) HolderID() uint {
	return e.ID
}
func (e Warehouse) HolderType() string {
	return "Warehouse"
}

func (dOut DeviceOut) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(dOut.FromReportItemID, dOut.ToWhomID, "Employee", int(dOut.Quantity))
	return
}

func (dOut DeviceOut) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(dOut.FromReportItemID, dOut.ToWhomID, "Employee", -1*int(dOut.Quantity))
	return
}

func (dIn DeviceIn) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse", dIn.Quantity)
	return
}

func (dIn DeviceIn) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse", -1*dIn.Quantity)
	return
}

func (cIn ConsumableIn) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(cIn.ReportItemID, 0, "Warehouse", -1*cIn.Quantity)
	return
}

func (cIn ConsumableIn) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(cIn.ReportItemID, 0, "Warehouse", cIn.Quantity)
	return
}

func (cOut ConsumableOut) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(cOut.ReportItemID, 0, "Warehouse", cOut.Quantity)
	return
}

func (cOut ConsumableOut) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(cOut.ReportItemID, 0, "Warehouse", -1*cOut.Quantity)
	return
}

func moveDeviceByID(fromReportItemID uint, toHolderId uint, toHolderType string, quantity int) (err error) {
	var d Device
	var from, to DeviceHolder
	from, to, d, err = fromToDevice(fromReportItemID, toHolderId, toHolderType)

	err = moveDevice(from, to, &d, quantity)
	return
}

func fromToDevice(fromReportItemID uint, toHolderId uint, toHolderType string) (from DeviceHolder, to DeviceHolder, d Device, err error) {
	d, from, err = deviceAndHolderByReportItem(fromReportItemID)
	if err != nil {
		log.Println(err)
		return
	}

	if toHolderId > 0 {
		to, err = holderByIDType(toHolderId, toHolderType)

		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func moveDevice(from DeviceHolder, to DeviceHolder, device *Device, quantity int) (err error) {
	var fromRi, toRi *ReportItem
	fromRi, err = getOrCreateReportItem(from, device, 0)
	if err != nil {
		log.Println(err)
		return
	}

	fcount := fromRi.Count - quantity
	tcount := 0
	if to != nil {
		toRi, err = getOrCreateReportItem(to, device, 0)
		if err != nil {
			log.Println(err)
			return
		}
		tcount = toRi.Count + quantity
	}

	if fcount < 0 || tcount < 0 {
		err = errors.New(fmt.Sprintf("数量输入有误"))
		return
	}

	err = DB.Model(&fromRi).Where("id = ?", fromRi.ID).UpdateColumn("count", gorm.Expr("count - ?", quantity)).Error
	if err != nil {
		log.Println(err)
		return
	}

	if to != nil {
		err = DB.Model(&toRi).Where("id = ?", toRi.ID).UpdateColumn("count", gorm.Expr("count + ?", quantity)).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func holderByIDType(id uint, t string) (h DeviceHolder, err error) {
	switch t {
	case "Employee":
		employee := Employee{}
		err = DB.Find(&employee, id).Error
		h = employee
	case "Warehouse":
		warehouse := Warehouse{}
		err = DB.Find(&warehouse, id).Error
		h = warehouse
	}
	return
}

func (d Device) AfterCreate(db *gorm.DB) (err error) {

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, d.WarehouseID).Error
	if err != nil {
		return
	}

	_, err = getOrCreateReportItem(warehouse, &d, d.TotalQuantity)
	return
}

func (d Device) BeforeUpdate(db *gorm.DB) (err error) {

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, d.WarehouseID).Error
	if err != nil {
		return
	}

	var ri *ReportItem
	ri, err = getOrCreateReportItem(warehouse, &d, d.TotalQuantity)
	if err != nil {
		return
	}

	var oldDev Device
	err = DB.Find(&oldDev, d.ID).Error
	if err != nil {
		return
	}

	inc := d.TotalQuantity - oldDev.TotalQuantity

	if ri.Count+int(inc) < 0 {
		err = errors.New(fmt.Sprintf("更新后的库存数量不能小于零，在库%d", ri.Count))
		return
	}

	err = DB.Model(&ReportItem{}).Where(&ReportItem{ID: ri.ID}).UpdateColumns(&ReportItem{Count: ri.Count + int(inc)}).Error
	if err != nil {
		return
	}

	err = DB.Model(&ReportItem{}).Where(&ReportItem{DeviceID: d.ID}).UpdateColumns(&ReportItem{DeviceCode: d.Code, DeviceName: d.Name}).Error
	return
}

func (d Device) BeforeDelete(db *gorm.DB) (err error) {

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, d.WarehouseID).Error
	if err != nil {
		return
	}

	var ri *ReportItem
	ri, err = getOrCreateReportItem(warehouse, &d, 0)
	if uint(ri.Count) != d.TotalQuantity {
		err = errors.New(fmt.Sprintf("有人带出设备%s，不能删除，当前库存数量%d，总数量%d", d.Name, ri.Count, d.TotalQuantity))
		return
	}

	err = DB.Delete(&ri).Error
	return
}

func deviceAndHolderByReportItem(reportItemID uint) (device Device, holder DeviceHolder, err error) {

	var ri ReportItem

	err = DB.Find(&ri, reportItemID).Error

	if err != nil {
		log.Println(err)
		return
	}

	err = DB.Find(&device, ri.DeviceID).Error
	if err != nil {
		log.Println(err)
		return
	}

	holder, err = holderByIDType(ri.WhoHasThemID, ri.WhoHasThemType)
	if err != nil {
		log.Println(err)
	}
	return
}

func getOrCreateReportItem(holder DeviceHolder, device *Device, count uint) (r *ReportItem, err error) {
	var reportItem ReportItem

	DB.Where(&ReportItem{DeviceID: device.ID, WhoHasThemID: holder.HolderID(), WhoHasThemType: holder.HolderType()}).Find(&reportItem)

	if reportItem.ID > 0 {
		r = &reportItem
		return
	}

	reportItem = ReportItem{
		WhoHasThemName:   holder.HolderName(),
		WhoHasThemID:     holder.HolderID(),
		WhoHasThemType:   holder.HolderType(),
		DeviceID:         device.ID,
		DeviceCode:       device.Code,
		DeviceName:       device.Name,
		DeviceCategoryID: device.CategoryID,
		Count:            int(count),
	}

	err = DB.Create(&reportItem).Error
	if err != nil {
		log.Println(err)
	}
	r = &reportItem
	return
}

func (cdOut ClientDeviceOut) AfterCreate(db *gorm.DB) (err error) {
	err = DB.Where(&ReportItem{ClientDeviceInID: cdOut.ClientDeviceInID}).Delete(&ReportItem{}).Error
	return
}

func (cdOut ClientDeviceOut) BeforeDelete(db *gorm.DB) (err error) {
	err = DB.Model(&ReportItem{}).Unscoped().Where(&ReportItem{ClientDeviceInID: cdOut.ClientDeviceInID}).UpdateColumn("deleted_at", nil).Error
	return
}

func (cdIn ClientDeviceIn) AfterCreate(db *gorm.DB) (err error) {
	err = DB.Unscoped().Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Delete(&ReportItem{}).Error
	if err != nil {
		return
	}

	err = createOrUpdateReportItem(
		cdIn.ID,
		cdIn.WarehouseID,
		cdIn.DeviceName,
		cdIn.ClientName,
		cdIn.ByWhomID,
		cdIn.Quantity)
	return
}

func (cdIn ClientDeviceIn) BeforeDelete(db *gorm.DB) (err error) {
	var cdOut ClientDeviceOut
	err = DB.Where(&ClientDeviceOut{ClientDeviceInID: cdIn.ID}).Find(&cdOut).Error
	if cdOut.ID > 0 {
		err = errors.New("设备已经还回，不能删除。")
		return
	}

	err = DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Delete(&ReportItem{}).Error
	return
}

func createOrUpdateReportItem(clientDeviceInID uint, warehouseId uint, deviceName string, clientName string, operatedByWhomId uint, quantity int) (err error) {

	var reportItem ReportItem

	bywhom := Employee{}
	err = DB.Find(&bywhom, operatedByWhomId).Error
	if err != nil {
		log.Println(err)
		return
	}

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, warehouseId).Error
	if err != nil {
		log.Println(err)
		return
	}

	if DB.Where(&ReportItem{ClientDeviceInID: clientDeviceInID}).Find(&reportItem).RecordNotFound() {
		reportItem := ReportItem{
			WhoHasThemName:     warehouse.Name,
			WhoHasThemID:       warehouse.ID,
			WhoHasThemType:     "Warehouse",
			ClientName:         clientName,
			DeviceName:         deviceName,
			Count:              quantity,
			OperatedByWhomID:   bywhom.ID,
			OperatedByWhomName: bywhom.Name,
			ClientDeviceInID:   clientDeviceInID,
		}
		err = DB.Create(&reportItem).Error
		if err != nil {
			log.Println(err)
		}
		return
	}

	if quantity > 0 {
		err = DB.Model(&reportItem).UpdateColumn("count", gorm.Expr("count + ?", quantity)).Error
	} else {
		err = DB.Model(&reportItem).UpdateColumn("count", gorm.Expr("count - ?", quantity)).Error
	}
	if err != nil {
		log.Println(err)
	}

	return
}
