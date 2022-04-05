package handler

import (
	"hand_held/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ProductCreateInitial(c echo.Context) error {
	req := new(model.ProductCreateInitialReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp int
	err := h.db.Raw("EXEC StkMs01CreateInitial  @ItemCode = ?, @GroupCode = ?, @BarCode = ?, @Name = ?, @MinorPerMajor = ?, @AccountSerial = ?, @ActiveItem = ?, @ItemTypeID = ?, @ItemHaveSerial = ?, @MasterItem = ?, @ItemHaveAntherUint = ?, @StoreCode = ?, @LastBuyPrice = ?, @POSTP = ?, @POSPP = ?, @Ratio1 = ?, @Ratio2 = ? ", req.ItemCode, req.GroupCode, req.BarCode, req.Name, req.MinorPerMajor, req.AccountSerial, req.ActiveItem, req.ItemTypeID, req.ItemHaveSerial, req.MasterItem, req.ItemHaveAntherUint, req.StoreCode, req.LastBuyPrice, req.POSTP, req.POSPP, req.Ratio1, req.Ratio2).Row().Scan(&resp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ProductFindByCode(c echo.Context) error {

	var item model.ProductCreateInitialReq
	err := h.db.Raw("EXEC StkMs01FindByCode  @BCode = ?, @StoreCode = ? , @Name = ?", c.FormValue("BCode"), c.FormValue("StoreCode"), c.FormValue("Name")).Row().Scan(
		&item.ItemCode,
		&item.GroupCode,
		&item.BarCode,
		&item.Name,
		&item.ItemTypeID,
		&item.MinorPerMajor,
		&item.AccountSerial,
		&item.ActiveItem,
		&item.ItemHaveSerial,
		&item.MasterItem,
		&item.ItemHaveAntherUint,
		&item.StoreCode,
		&item.LastBuyPrice,
		&item.POSTP,
		&item.POSPP,
		&item.Ratio1,
		&item.Ratio2,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, item)
}

func (h *Handler) ProductGetMaxCode(c echo.Context) error {
	var resp model.GroupCodeAndMaxItem
	err := h.db.Raw("EXEC StkMs01MacItemCodeByGroup @GroupCode = ?", c.Param("group")).Row().Scan(&resp.MaxCode, &resp.GroupName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ItemTypeByGroup(c echo.Context) error {
	var resp []model.ItemType
	rows, err := h.db.Raw("EXEC ItemTypeByGroup @GroupCode = ?", c.Param("group")).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.ItemType
		err = rows.Scan(&item.ItemTypeID, &item.ItemTypeName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		resp = append(resp, item)
	}
	return c.JSON(http.StatusOK, resp)

}
