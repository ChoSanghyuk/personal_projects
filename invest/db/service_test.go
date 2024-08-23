package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var stg *Storage

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

	stg = &Storage{
		db: db,
	}

}
func TestRetrieveFundAmount(t *testing.T) {

	rst, err := stg.RetrieveFundAmount()
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))

}

func TestRetrieveFundAmountById(t *testing.T) {

	rst, err := stg.RetrieveFundAmountById(1)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))

}

func TestRetreiveInvestHistOfFund(t *testing.T) {

	rst, err := stg.RetreiveInvestHistOfFund()
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}

func TestRetreiveAssetOfFundById(t *testing.T) {
	rst, err := stg.RetreiveAssetOfFundById(1)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}

func TestRetrieveAssetList(t *testing.T) {

	rst, err := stg.RetrieveAssetList()
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))

}

func TestRetrieveAssetInfo(t *testing.T) {
	rst, err := stg.RetrieveAssetInfo(1)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}

func TestRetrieveAssetHist(t *testing.T) {
	rst, err := stg.RetrieveAssetHist(1)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}

func TestRetrieveMarketSituation(t *testing.T) {
	rst, err := stg.RetrieveMarketSituation("2024-08-20")
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}

func TestRetrieveInvestHist(t *testing.T) {
	rst, err := stg.RetrieveInvestHist(1, 0, "", "")
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(rst, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j))
}
