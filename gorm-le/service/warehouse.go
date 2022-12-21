package service

import (
	"fmt"
	"gorm-le/db"
	"gorm-le/model"
)

func Page() (model.IPage, error) {
	warehouse := model.WmsWarehouse{}
	wmsWarehouseMgr := model.WmsWarehouseMgr(db.MysqlDb)
	wmsWarehouseMgr.First(&warehouse)
	var warehouseCount int64
	wmsWarehouseMgr.Count(&warehouseCount)
	page, err := wmsWarehouseMgr.SelectPage(model.NewPage(10, 1))
	if err != nil {
		return nil, err
	}
	fmt.Println(warehouse)
	fmt.Println(warehouseCount)
	fmt.Println(page)
	return page, nil
}
