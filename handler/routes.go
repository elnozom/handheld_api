package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	cashtry := v1.Group("/cashtry")
	cashtry.GET("/stores", h.CashTryStores)
	v1.GET("/employee", h.GetEmp)
	v1.GET("/doc/undestributed/:bcode", h.GetUndestributedDoc)
	v1.GET("/doc/area", h.GetAreaOrder)

	v1.POST("/invenetory", h.InventorySession)
	v1.GET("/invoice", h.GetPrepareDoc)
	v1.GET("/invoice/item", h.IsItemInInvoice)
	v1.GET("/invoice/open", h.GetOpenPrepareDocs)
	v1.POST("/invoice/close", h.ClosePrepareDoc)
	v1.POST("/invoice", h.UpdatePrepareItem)
	v1.POST("/raseed", h.RaseedInvInsert)
	v1.POST("/devices/check", h.DevicesCheck)
	v1.POST("/devices/insert", h.DevicesInsert)
	v1.POST("/orders/close", h.CloseOrder)
	v1.GET("/msgs", h.GetMsgs)
	v1.PUT("/msgs/read", h.ReadMsgs)
	v1.POST("/get-account", h.GetAccount)
	v1.POST("/get-item", h.GetItem) // done
	v1.POST("/get-doc", h.GetDocNo)
	v1.POST("/get-doc-items", h.GetDocItems)
	v1.POST("/insert-item", h.InsertItem)
	v1.POST("/delete-item", h.DeleteItem)
	v1.POST("/get-docs", h.GetOpenDocs)
	v1.POST("/close-doc", h.CloseDoc)
	v1.POST("/update-driver", h.UpdateDriver)

	// order
	v1.POST("/orders", h.InsertOrder)
	v1.POST("/orders/item", h.InsertOrderItem)
	v1.GET("/orders/items", h.GetOrderItems)

	// product
	v1.POST("/products/info", h.ProductCreateInitial)
	v1.GET("/products/maxCode/:group", h.ProductGetMaxCode)
	v1.GET("/products/types/:group", h.ItemTypeByGroup)
	v1.GET("/products/find", h.ProductFindByCode)

}
