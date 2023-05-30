package main

import (
	"fmt"
	"hand_held/config"
	"hand_held/db"
	"hand_held/handler"
	"hand_held/repo"
	"hand_held/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	db.InitDatabase()
	db := db.DBConn
	companyRepo := repo.NewCompanyRepo(db)
	permissionsRepo := repo.NewPermissionsRepo(db)
	accountsRepo := repo.NewAccountsRepo(db)
	companyInfo, err := companyRepo.Find()
	if err != nil {
		panic(err)
	}
	h := handler.NewHandler(db, companyInfo, permissionsRepo, accountsRepo)
	h.Register(v1)
	port := fmt.Sprintf("192.168.1.40:%s", config.Config("PORT"))
	r.Logger.Fatal(r.Start(port))
}
