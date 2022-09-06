package handler

import (
	"hand_held/model"
	"hand_held/repo"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	db              *gorm.DB
	info            *model.CompanyInfo
	permissionsRepo repo.PermissionsRepo
	accountsRepo    repo.AccountsRepo
}

func NewHandler(databaase *gorm.DB, companyInfo *model.CompanyInfo, permissionsRepo repo.PermissionsRepo, accountsRepo repo.AccountsRepo) *Handler {
	return &Handler{
		db:              databaase,
		info:            companyInfo,
		permissionsRepo: permissionsRepo,
		accountsRepo:    accountsRepo,
	}
}
