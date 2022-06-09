package handler

import (
	"hand_held/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DirectOrderInsert(c echo.Context) error {

	req := new(model.InsertDirectOrderReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var resp int
	err := h.db.Raw(
		"EXEC StkTr01Insert  @AccountSerial = ?,@EmpCode = ?,@StoreCode = ?,@StoreCode2 = ?, @ComputerName = ?  ,@HeadSerial = ?,@TransSerial = ?,@ItemSerial = ?,@Qnt = ?,@Price =  ?,@Tax = ?,@MinorPerMajor = ?",
		req.AccountSerial,
		req.EmpCode,
		req.StoreCode,
		req.StoreCode2,
		req.ComputerName,
		req.HeadSerial,
		req.TransSerial,
		req.ItemSerial,
		req.Qnt,
		req.Price,
		req.Tax,
		req.MinorPerMajor,
	).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DirectOrderItemsDelete(c echo.Context) error {
	var resp int
	serial, _ := strconv.Atoi(c.Param("id"))
	err := h.db.Raw("EXEC StkTr02Delete  @Serial = ?", serial).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DirectOrdersList(c echo.Context) error {
	var resp []model.DirectOrder
	req := new(model.ListDirectOrdersReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	rows, err := h.db.Raw("EXEC StkTr01List  @TransSerial = ? , @StoreCode = ? , @ComputerName = ?", req.TransSerial, req.StoreCode, req.ComputerName).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.DirectOrder
		err := rows.Scan(
			&rec.Serial,
			&rec.StoreCode,
			&rec.DocNo,
			&rec.AccountSerial,
			&rec.TransSerial,
			&rec.TotalCash,
			&rec.AccountName,
			&rec.AccountCode,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't ssscan the values : "+err.Error())
		}
		resp = append(resp, rec)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DirectOrderPrintList(c echo.Context) error {
	var resp []model.DirectOrderPrint
	id, _ := strconv.Atoi(c.Param("id"))
	rows, err := h.db.Raw("EXEC StkTr01PrintItemsBySerial  @Serial = ?", id).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.DirectOrderPrint
		err := rows.Scan(
			&rec.DocDate,
			&rec.DocTime,
			&rec.AccountName,
			&rec.DocNo,
			&rec.ItemName,
			&rec.EmpName,
			&rec.Qnt,
			&rec.MinorPerMajor,
			&rec.Price,
			&rec.TotalPrice,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values : "+err.Error())
		}
		resp = append(resp, rec)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DirectOrderItemsList(c echo.Context) error {
	var resp []model.DocItem
	serial, _ := strconv.Atoi(c.Param("id"))
	rows, err := h.db.Raw("EXEC StkTr01ListItems  @Serial = ?", serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.DocItem
		err := rows.Scan(
			&rec.Serial,
			&rec.Qnt,
			&rec.Price,
			&rec.Item_BarCode,
			&rec.TotalCash,
			&rec.ItemName,
			&rec.MinorPerMajor,
			&rec.ItemTotal,
			&rec.ByWeight,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values : "+err.Error())
		}
		resp = append(resp, rec)
	}
	return c.JSON(http.StatusOK, resp)
}

// req := new(model.DocItemsReq)
// if err := c.Bind(req); err != nil {
// 	return c.JSON(http.StatusBadRequest, "ERROR binding request")
// }
// var DocItems []model.DocItem
// rows, err := h.db.Raw("EXEC GetSdItems @DevNo = ?, @TrSerial = ?,@StoreCode = ? , @DocNo = ?;", req.DevNo, req.TrSerial, req.StoreCode, req.DocNo).Rows()
// if err != nil {
// 	return c.JSON(http.StatusInternalServerError, err)
// }
// defer rows.Close()
// for rows.Next() {
// 	var docItem model.DocItem
// 	err = rows.Scan(
// 		&docItem.Serial,
// 		&docItem.Qnt,
// 		&docItem.Item_BarCode,
// 		&docItem.ItemName,
// 		&docItem.MinorPerMajor,
// 		&docItem.ByWeight,
// 	)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "can't scan the values")
// 	}
// 	DocItems = append(DocItems, docItem)
// }

// return c.JSON(http.StatusOK, DocItems)
