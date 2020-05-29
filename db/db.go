package db

import (
	"fmt"
	"github.com/Ghun2/fast-gorestapi/config"
	"github.com/Ghun2/fast-gorestapi/model"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/Ghun2/fast-gorestapi/model"
)

func New() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
			"loc":"Local",
		},
	}
	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(5)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               "TEST",
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
			"loc":"Local",
		},
	}
	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		fmt.Println("test storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB(db *gorm.DB) error {
	db.DropTableIfExists(
		&model.User{},
	)
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
