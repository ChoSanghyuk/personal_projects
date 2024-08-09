package db

import (
	"database/sql"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	dsn := "root:root@tcp(127.0.0.1:3300)/investdb?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestMigration(t *testing.T) {
	db.AutoMigrate(&Fund{}, &Asset{}, &FundStatus{}, &InvestHistory{})
}

func TestCreate(t *testing.T) {
	fund := Fund{
		Name:   "개인",
		Amount: 1000000,
	}

	result := db.Create(&fund)

	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Log("ID", fund.Id)
	t.Log("Rows Affected", result.RowsAffected)
}
