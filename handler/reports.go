package handler

import (
	"database/sql"
	"fmt"
	"hand_held/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetDocNo(c echo.Context) error {

	req := new(model.DocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	// return c.JSON(http.StatusOK, "test")

	var DocNo []model.Doc
	rows, err := h.db.Raw("EXEC GetSdDocNo @DevNo = ?, @TrSerial = ?,@StoreCode = ?;", req.DevNo, req.TrSerial, req.StoreCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var doc model.Doc
		err = rows.Scan(
			&doc.DocNo,
		)
		print(rows)
		if err != nil {
			return c.JSON(http.StatusOK, 1)
		}
		DocNo = append(DocNo, doc)
	}

	return c.JSON(http.StatusOK, DocNo[0].DocNo+1)
}

func (h *Handler) GetOpenDocs(c echo.Context) error {

	req := new(model.OpenDocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var OpenDocs []model.OpenDoc
	rows, err := h.db.Raw("EXEC GetOpenSdDocNo @StCode = ?,@DevNo = ?, @TrSerial = ?;", req.StCode, req.DevNo, req.TrSerial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var openDoc model.OpenDoc
		err = rows.Scan(
			&openDoc.DocNo,
			&openDoc.StoreCode,
			&openDoc.AccontSerial,
			&openDoc.TransSerial,
			&openDoc.AccountName,
			&openDoc.AccountCode,
			&openDoc.DevNo,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		OpenDocs = append(OpenDocs, openDoc)
	}

	return c.JSON(http.StatusOK, OpenDocs)
}
func (h *Handler) InsertOrder(c echo.Context) error {
	req := new(model.InsertOrder)
	if err := c.Bind(req); err != nil {
		return err
	}
	orderNoRows, err := h.db.Raw("EXEC GetSalesOrderDocNo @StoreCode = ?, @TrSerial = ?", req.StoreCode, 30).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var orderNo int32
	for orderNoRows.Next() {
		orderNoRows.Scan(&orderNo)
	}
	orderNo = orderNo + 1
	fmt.Println(orderNo)

	rows, err := h.db.Raw("EXEC InsertTr05 @DocNo = ?, @StoreCode = ? , @EmpCode = ? , @AccountSerial =? ", orderNo, req.StoreCode, req.EmpCode, req.AccountSerial).Rows()
	defer rows.Close()
	var serial int
	for rows.Next() {
		err = rows.Scan(&serial)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, serial)
}

func (h *Handler) CloseOrder(c echo.Context) error {
	type Req struct {
		Serial    int
		TotalCash float64
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return err
	}

	_, err := h.db.Raw("EXEC CloseTr05  @Serial = ? ", req.Serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "closed")
}

func (h *Handler) DevicesCheck(c echo.Context) error {
	type Req struct {
		DeviceId string
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp int
	rows, err := h.db.Raw("EXEC DevicesCheck @DeviceId = ?", req.DeviceId).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&resp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DevicesInsert(c echo.Context) error {
	type Req struct {
		DeviceId string
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp int
	rows, err := h.db.Raw("EXEC DevicesInsert @DeviceId = ?", req.DeviceId).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&resp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetOrderItems(c echo.Context) error {

	req := new(model.GetOrderItemsRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req.Serial)

	var items []model.OrderItem
	rows, err := h.db.Raw("EXEC GetDocItemData @Serial = ?", req.Serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.OrderItem
		err = rows.Scan(&item.Serial, &item.BarCode, &item.ItemName, &item.Qnt, &item.Price, &item.Total)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *Handler) InsertOrderItem(c echo.Context) error {

	req := new(model.InsertOrderItem)
	if err := c.Bind(req); err != nil {
		return err
	}
	rows, err := h.db.Raw("EXEC InsertTr06 @HeadSerial = ?, @ItemSerial = ? , @Qnt = ? , @Price = ? , @QntAntherUnit = ?", req.HeadSerial, req.ItemSerial, req.Qnt, req.Price, req.QntAntherUnit).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	type Resp struct {
		TotalPackages int
		TotalCash     float64
		Serial        int32
	}
	var resp = new(Resp)
	for rows.Next() {
		err = rows.Scan(&resp.TotalPackages, &resp.TotalCash)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	if rows.NextResultSet() {
		for rows.Next() {
			fmt.Println("asdasdasd")
			err = rows.Scan(&resp.Serial)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

	}

	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetOpenPrepareDocs(c echo.Context) error {
	var OpenDocs []model.PrepareDocResp
	rows, err := h.db.Raw("EXEC GetOpenPrepare;").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var openDoc model.PrepareDocResp
		err = rows.Scan(
			&openDoc.DocNo,
			&openDoc.AccountName,
			&openDoc.AccountCode,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		OpenDocs = append(OpenDocs, openDoc)
	}

	return c.JSON(http.StatusOK, OpenDocs)
}
func (h *Handler) GetDocItems(c echo.Context) error {

	req := new(model.DocItemsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var DocItems []model.DocItem
	rows, err := h.db.Raw("EXEC GetSdItems @DevNo = ?, @TrSerial = ?,@StoreCode = ? , @DocNo = ?;", req.DevNo, req.TrSerial, req.StoreCode, req.DocNo).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var docItem model.DocItem
		err = rows.Scan(
			&docItem.Serial,
			&docItem.Qnt,
			&docItem.Item_BarCode,
			&docItem.ItemName,
			&docItem.MinorPerMajor,
			&docItem.ByWeight,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		DocItems = append(DocItems, docItem)
	}

	return c.JSON(http.StatusOK, DocItems)
}

func (h *Handler) DeleteItem(c echo.Context) error {

	req := new(model.DeleteItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := h.db.Raw("EXEC DeleteSdItem  @Serial = ?; ", req.Serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) InsertItem(c echo.Context) error {

	req := new(model.InsertItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}

	_, err := h.db.Raw(
		"EXEC InsertSdDocNo  @DNo = ? ,@TrS = ? ,@AccS = ? ,@ItmS =?  ,@Qnt = ? ,@StCode = ? ,@InvNo = ? ,@ItmBarCode = ? ,@DevNo = ?,@StCode2 = ?,@ExpDate = ?, @SessionNo = ?; ", req.DNo, req.TrS, req.AccS, req.ItmS, req.Qnt, req.StCode, req.InvNo, req.ItmBarCode, req.DevNo, req.StCode2, req.ExpDate, req.SessionNo).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "inserted")
}

func (h *Handler) UpdatePrepareItem(c echo.Context) error {
	req := new(model.UpdatePrepareReq)
	fmt.Println(req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	rows, err := h.db.Raw(
		"EXEC UpdatePrepare  @QPrep = ? ,@ISerial = ? ,@HSerial = ? ,@EmpCode = ? ;", req.QPrep, req.ISerial, req.HSerial, req.EmpCode).Rows()
	if err != nil {
		print(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	var resp []model.UpdatePrepareResp
	var item model.UpdatePrepareResp
	for rows.Next() {
		err = rows.Scan(
			&item.Prepared,
			&item.QntPrepared,
			&item.Qnt,
			// &item.HeadPrepared,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		resp = append(resp, item)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ClosePrepareDoc(c echo.Context) error {
	req := new(model.ClosePrepareDocReq)
	if err := c.Bind(req); err != nil {
		fmt.Println("err1")
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	rows, err := h.db.Raw("EXEC ChkPrepare @HSerial = ?, @EmpCode = ?", req.HSerial, req.EmpCode).Rows()
	defer rows.Close()
	var resp []model.ClosePrepareDocResp

	for rows.Next() {
		var res model.ClosePrepareDocResp
		err = rows.Scan(&res.Close)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		resp = append(resp, res)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) InventorySession(c echo.Context) error {
	req := new(model.InventorySessionReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	rows, err := h.db.Raw("EXEC CheckGardSession @StoreCode = ?", req.StoreCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	var resp []model.InventorySessionResp

	for rows.Next() {
		var res model.InventorySessionResp
		err = rows.Scan(&res.SessionNo, &res.PartInv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		resp = append(resp, res)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)

}

func (h *Handler) RaseedInvInsert(c echo.Context) error {
	req := new(model.RaseedInvInsertReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	rows, err := h.db.Raw("EXEC RaseedInvInsert @ItemSerial = ?, @I = ? , @R = ? , @SessionNo = ? , @StoreCode = ? ", req.ItemSerial, req.I, req.R, req.SessionNo, req.StoreCode).Rows()
	defer rows.Close()
	var resp model.RaseedInvInsertResp

	for rows.Next() {
		err = rows.Scan(&resp.Differnce, &resp.Inserted)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)

}
func (h *Handler) CashTryStores(c echo.Context) error {
	var stores []model.CashtryStores
	// return c.JSON(http.StatusOK, "test")
	rows, err := h.db.Raw("EXEC GetStoreName").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer rows.Close()
	for rows.Next() {
		var store model.CashtryStores
		rows.Scan(&store.StoreCode, &store.StoreName)
		stores = append(stores, store)
	}

	return c.JSON(http.StatusOK, stores)
}

func (h *Handler) GetAccount(c echo.Context) error {

	req := new(model.GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var accounts []model.Account
	rows, err := h.db.Raw("EXEC GetAccount @Code = ?, @Name = ? , @Type = ?", req.Code, req.Name, req.Type).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		rows.Scan(&account.Serial, &account.AccountCode, &account.AccountName)
		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) CloseDoc(c echo.Context) error {
	req := new(model.CloseDocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println(req.DevNo, req.Trans, req.DocNo)
	rows, err := h.db.Raw("EXEC CloseSdDoc @DevNo = ?, @Trans = ? , @DocNo = ?", req.DevNo, req.Trans, req.DocNo).Rows()
	defer rows.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "closed")

}

func (h *Handler) GetItem(c echo.Context) error {

	req := new(model.GetItemRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var items []model.Item
	var rows *sql.Rows
	var err error
	if req.Name == "" {
		rows, err = h.db.Raw("EXEC GetItemData @BCode = ?, @StoreCode = ? ", req.BCode, req.StoreCode).Rows()
	} else {
		rows, err = h.db.Raw("EXEC GetItemData @BCode = ?, @StoreCode = ? , @Name = ? ", req.BCode, req.StoreCode, req.Name).Rows()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.Serial, &item.ItemName, &item.MinorPerMajor, &item.POSPP, &item.POSTP, &item.ByWeight, &item.WithExp, &item.ItemHasAntherUnit, &item.AvrWait, &item.Expirey, &item.I, &item.R)
		if err != nil {
			fmt.Println("Err2")
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *Handler) GetMsgs(c echo.Context) error {

	req := new(model.GetMsgsRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var msgs []string
	var rows *sql.Rows
	var err error
	rows, err = h.db.Raw("EXEC GetMsgs @EmpSerial = ?, @BonSerial = ? ", req.EmpSerial, req.BonSerial).Rows()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var msg string
		err = rows.Scan(&msg)
		if err != nil {
			fmt.Println("Err2")
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		msgs = append(msgs, msg)
	}

	return c.JSON(http.StatusOK, msgs)
}

func (h *Handler) GetAreaOrder(c echo.Context) error {
	type Req struct {
		Area   int
		Serial int
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var resp []model.UndestributedDoc

	rows, err := h.db.Raw("EXEC GetAreaOrder @Area = ? , @Serial = ? ", req.Area, req.Serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.UndestributedDoc
		err = rows.Scan(&rec.BonSerial, &rec.DocNo, &rec.AccountCode, &rec.AccountName, &rec.AccountAddress, &rec.AccountArea)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		resp = append(resp, rec)
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetUndestributedDoc(c echo.Context) error {
	var rec model.UndestributedDoc
	var resp []model.UndestributedDoc

	err := h.db.Raw("EXEC GetUndistributeDoc @BCode = ? ", c.Param("bcode")).Row().Scan(
		&rec.BonSerial,
		&rec.DocNo,
		&rec.AccountCode,
		&rec.AccountName,
		&rec.AccountAddress,
		&rec.AccountArea,
		&rec.AreaName,
		&rec.BonCount,
	)
	resp = append(resp, rec)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ReadMsgs(c echo.Context) error {

	req := new(model.GetMsgsRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var rows *sql.Rows
	var err error
	rows, err = h.db.Raw("EXEC ReadMsgs @EmpSerial = ?, @BonSerial = ? ", req.EmpSerial, req.BonSerial).Rows()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	return c.JSON(http.StatusOK, "updated")
}
func (h *Handler) IsItemInInvoice(c echo.Context) error {

	req := new(model.ItemInInvReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var err error
	resp := new(model.ItemInInvResp)

	rows := h.db.Raw("EXEC IsItemInInvoice @ItemSerial = ?, @BonSerial = ? ", req.ItemSerial, req.BonSerial).Row()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = rows.Scan(&resp.Found)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

	// return c.JSON(http.StatusInternalServerError, "error")

}

func (h *Handler) GetEmp(c echo.Context) error {
	req := new(model.EmpReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req.EmpCode)

	var employee []model.Emp
	rows, err := h.db.Raw("EXEC GetEmp @EmpCode = ?", req.EmpCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Emp
		err = rows.Scan(&item.EmpName, &item.EmpPassword, &item.EmpCode, &item.SecLevel, &item.FixEmpStore)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		employee = append(employee, item)
	}

	return c.JSON(http.StatusOK, employee)
}

func (h *Handler) UpdateDriver(c echo.Context) error {
	type Req struct {
		BonSerial  string
		DriverCode int
	}
	req := new(Req)
	if err := c.Bind(req); err != nil {
		return err
	}
	rows, err := h.db.Raw("EXEC Stktr03UpdateDriver @BonSerial = ? , @DriverCode = ? ", req.BonSerial, req.DriverCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	return c.JSON(http.StatusOK, "updated")
}

func (h *Handler) GetPrepareDoc(c echo.Context) error {
	req := new(model.InvReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	var items []model.InvoiceItem
	rows, err := h.db.Raw("EXEC GetPrepareDoc @BCode  = ?", req.BCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.InvoiceItem
		err = rows.Scan(&item.BonSer, &item.Qnt, &item.Price, &item.IsPrepared, &item.QntPrepare, &item.ItemCode, &item.GroupCode, &item.MinorPerMajor, &item.ItemName, &item.ItemSerial)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}
