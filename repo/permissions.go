package repo

import (
	"github.com/jinzhu/gorm"
)

type PermissionsRepo struct {
	db *gorm.DB
}

func NewPermissionsRepo(db *gorm.DB) PermissionsRepo {
	return PermissionsRepo{
		db: db,
	}
}

func (ur *PermissionsRepo) LoadEmpPermissions(code *int) (*[]string, error) {
	var resp []string
	rows, err := ur.db.Raw("EXEC LoadEmpPermissions @EmpCode = ?", code).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var rec string
		rows.Scan(&rec)
		resp = append(resp, rec)
	}
	return &resp, nil
}
