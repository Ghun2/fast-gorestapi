package test

import (
	"encoding/json"
	"github.com/Ghun2/fast-gorestapi/api/controller"
	"github.com/Ghun2/fast-gorestapi/config"
	"github.com/Ghun2/fast-gorestapi/db"
	"github.com/Ghun2/fast-gorestapi/model"
	"github.com/Ghun2/fast-gorestapi/router"
	"github.com/Ghun2/fast-gorestapi/service/store"
	"github.com/Ghun2/fast-gorestapi/service/user"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"testing"
)

var (
	d    	*gorm.DB
	us   	user.Store
	ctrl 	*controller.Controller
	e    	*echo.Echo
)

func TestMain(m *testing.M) {
	config.ReadTestConfig()
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	d = db.TestDB()
	db.AutoMigrate(d)

	us = store.NewUserStore(d)

	ctrl = controller.NewController(us)
	e = router.New()
	if err := loadFixtures(); err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	if err := db.DropTestDB(d); err != nil {
		log.Fatal(err)
	}
	_ = d.Close()
}

func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m[key].(map[string]interface{})
}

func loadFixtures() error {
	u1 := model.User{
		UserName: 	"user1",
		AuthID:		"kakao:1234",
		Email:    	"user1@ghun2ee.com",
		Birth:      "920730",
		Sex:    	"1",
		Phone: 		"01012345678",
	}
	if err := us.Create(&u1); err != nil {
		return err
	}

	u2 := model.User{
		UserName: 	"user2",
		AuthID:		"apple:5678",
		Email:    	"user2@ghun2ee.com",
		Birth:      "920526",
		Sex:    	"2",
		Phone: 		"01098765432",
	}
	if err := us.Create(&u2); err != nil {
		return err
	}
}
