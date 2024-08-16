package db

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"gorm.io/datatypes"
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
	db.AutoMigrate(&Fund{}, &Asset{}, &FundStatus{}, &InvestHistory{}, &CliIdx{}, &Market{})
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

/*
결국은 time.Time 객체인 것이 중요한게 아닌, string형 변환했을 때 DB 타입과 일치하는지가 중요함
time.Time{}.Local() => '0000-00-00 00:00:00' 라서 Date 타입 및 Timestamp 실패
time.Now() => '2024-08-16 08:47:20.346' Date 타입 및 Timestamp 성공
*/
func TestTime(t *testing.T) {
	db.AutoMigrate(&Sample{})

	// date, _ := time.Parse("2006-01-02", "2021-11-22")

	d := Sample{
		Date: datatypes.Date(time.Now()),
		Time: time.Now(),
	}

	db.Debug().Create(&d)
}
