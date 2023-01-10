package repo

import (
	"hand_held/model"

	"github.com/jinzhu/gorm"
)

type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) CompanyRepo {
	return CompanyRepo{
		db: db,
	}
}

func (ur *CompanyRepo) Find() (*model.CompanyInfo, error) {
	var resp model.CompanyInfo
	err := ur.db.Raw("EXEC CompanyInfoFind").Row().Scan(
		&resp.CurChrc,
		&resp.ReportTitle,
		&resp.BonMsg1,
		&resp.BonMsg2,
		&resp.BonMsg3,
		&resp.BonMsg4,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
