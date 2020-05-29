package main

import (
	"github.com/Ghun2/fast-gorestapi/api/controller"
	"github.com/Ghun2/fast-gorestapi/config"
	"github.com/Ghun2/fast-gorestapi/db"
	"github.com/Ghun2/fast-gorestapi/router"
	"github.com/Ghun2/fast-gorestapi/service/store"
)

func main() {
	config.ReadConfig()
	r := router.New()

	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)

	c := controller.NewController(us)
	c.Register(v1)
	r.Logger.Fatal(r.Start("localhost:3000"))
}