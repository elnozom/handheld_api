package main

import (
	"fmt"
	"hand_held/config"
	"hand_held/db"
	"hand_held/handler"
	"hand_held/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	db.InitDatabase()
	db := db.DBConn
	h := handler.NewHandler(db)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	r.Logger.Fatal(r.Start(port))
}
