package repo

import (
	"hand_held/model"

	"github.com/jinzhu/gorm"
)

type AccountsRepo struct {
	db *gorm.DB
}

func NewAccountsRepo(db *gorm.DB) AccountsRepo {
	return AccountsRepo{
		db: db,
	}
}

func (ar *AccountsRepo) AccountsTransactionInsert(req *model.AccountTransactionReq) (*int, error) {
	var resp int
	err := ar.db.Raw("EXEC AccTr01Insert  @AccMoveSerial = ? , @Stc = ? , @UserCode = ?, @AccType = ? ,@AccountSerial = ? ,@AccountSerial2 = ? ,@Amount =? ; ", req.TransactionType, req.Store, 1, req.AccType, req.Safe, req.AccSerial, req.Amount).Row().Scan(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
