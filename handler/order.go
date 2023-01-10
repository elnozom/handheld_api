package handler

import (
	"hand_held/model"
	"net/http"
	"strconv"
	"strings"

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
	var fromDate *string
	var toDate *string
	// if *req.FromDate != "null" {
	// 	fromDate = req.FromDate
	// }
	// if *req.ToDate != "null" {
	// 	toDate = req.ToDate
	// }
	rows, err := h.db.Raw("EXEC StkTr01List  @TransSerial = ? , @StoreCode = ? , @ComputerName = ? , @isClosed = ? , @fromDate = ?, @toDate = ?", req.TransSerial, req.StoreCode, req.ComputerName, req.IsClosed, fromDate, toDate).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.DirectOrder
		err := rows.Scan(
			&rec.Serial,
			&rec.DocDate,
			&rec.Discount,
			&rec.StoreCode,
			&rec.DocNo,
			&rec.AccountSerial,
			&rec.TransSerial,
			&rec.Total,
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
	var resp model.PrintResponse
	id, _ := strconv.Atoi(c.Param("id"))
	rows, err := h.db.Raw("EXEC StkTr01PrintItemsBySerial  @Serial = ?", id).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	var grandTotal float64
	var wholeQntSum float64
	var partQntSum float64
	for rows.Next() {
		var rec model.DirectOrderPrint
		err := rows.Scan(
			&rec.DocDate,
			&rec.AccountName,
			&rec.DocNo,
			&rec.ItemName,
			&rec.EmpName,
			&rec.StoreName,
			&rec.WholeQnt,
			&rec.PartQnt,
			&rec.MinorPerMajor,
			&rec.Price,
			&rec.TotalPrice,
		)
		grandTotal += rec.TotalPrice
		wholeQntSum += rec.WholeQnt
		partQntSum += rec.PartQnt
		rec.ItemName = _truncateText(rec.ItemName, 30)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values : "+err.Error())
		}
		resp.Items = append(resp.Items, rec)
	}
	resp.Totals.GrandTotal = grandTotal
	resp.Totals.WholeQnt = wholeQntSum
	resp.Totals.PartQnt = partQntSum
	resp.Info = *h.info
	return c.JSON(http.StatusOK, resp)
}

func _truncateText(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:strings.LastIndex(s[:max], " ")]
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
