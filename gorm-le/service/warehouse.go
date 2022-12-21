package service

import (
	"fmt"
	"gorm-le/db"
	"gorm-le/model"
)

func Page() model.IPage {
	warehouse := model.WmsWarehouse{}
	wmsWarehouseMgr := model.WmsWarehouseMgr(db.MysqlDb)
	wmsWarehouseMgr.First(&warehouse)
	var warehouseCount int64
	wmsWarehouseMgr.Count(&warehouseCount)
	page, err := wmsWarehouseMgr.SelectPage(model.NewPage(2, 2))
	if err != nil {
		panic(err)
	}
	fmt.Println(warehouse)
	fmt.Println(warehouseCount)
	fmt.Println(page)
	return page
}
