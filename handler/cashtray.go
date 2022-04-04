package handler

import (
	"hand_held/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CashTryClose(c echo.Context) error {
	req := new(model.CloseCashtrayReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp int
	err := h.db.Raw("EXEC CashtrayClose  @Serial = ? ,@Exceed = ? ,@Shortage = ? ,@Amount = ?", req.Serial, req.Exceed, req.Shortage, req.Amount).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) CashTryFind(c echo.Context) error {

	var resp model.CashtryData

	err := h.db.Raw("EXEC CashTryData @CashTrySerial = ?;", c.Param("serial")).Row().Scan(
		&resp.CashTryNo,
		&resp.SessionNo,
		&resp.EmpCode,
		&resp.OpenDate,
		&resp.CloseDate,
		&resp.OpenTime,
		&resp.CloseTime,
		&resp.StartCash,
		&resp.TotalCash,
		&resp.ComputerName,
		&resp.TotalOrder,
		&resp.TotalHome,
		&resp.TotalIn,
		&resp.TotalOut,
		&resp.TotalVisa,
		&resp.TotalShar,
		&resp.TotalVoid,
		&resp.Halek,
		&resp.EndCash,
		&resp.Paused,
		&resp.CasherMoney,
		&resp.PayLater,
		&resp.HomeIn,
		&resp.HomeOutCashTry,
		&resp.CashTryTypeCode,
		&resp.Final,
		&resp.CasherCashTrySerial,
		&resp.Exceed,
		&resp.Shortage,
		&resp.StoreCode,
		&resp.TotalVat,
		&resp.DiscValue,
		&resp.TotalVoidCash,
		&resp.TotalVoidCrdt,
		&resp.TotalVoidVisa,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp.DeliveryNonReturn = resp.TotalHome - resp.HomeIn
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) PausedCashTry(c echo.Context) error {
	var pausedCashtries []model.PausedCashtry
	rows, err := h.db.Raw("EXEC CashTrayPaused").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cashtry model.PausedCashtry
		err = rows.Scan(
			&cashtry.Serial,
			&cashtry.EmpCode,
			&cashtry.EmpName,
			&cashtry.OpenDate,
			&cashtry.OpenTime,
			&cashtry.ComputerName,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values"+err.Error())

		}
		cashtry.OpenDate = strings.Split(cashtry.OpenDate, "T")[0]
		pausedCashtries = append(pausedCashtries, cashtry)
	}

	return c.JSON(http.StatusOK, pausedCashtries)
}
