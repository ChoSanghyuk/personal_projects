package db

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage() (*Storage, error) {
	dsn := "root:root@tcp(127.0.0.1:3300)/lolche?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s Storage) SaveMain(name string) error {
	main := main{
		Name: name,
	}

	result := s.db.Create(&main)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s Storage) SavePbe(name string) error {
	main := pbe{
		Name: name,
	}

	result := s.db.Create(&main)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s Storage) DeleteAllMain() error {
	result := s.db.Unscoped().Where("1 = 1").Delete(&main{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s Storage) DeleteAllPbe(name string) error {
	result := s.db.Unscoped().Where("1 = 1").Delete(&pbe{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s Storage) AllMain() ([]string, error) {

	var mains []main

	result := s.db.Model(&main{}).Select("name").Find(&mains)
	if result.Error != nil {
		return nil, result.Error
	}

	decs := make([]string, len(mains))
	for i := 0; i < len(mains); i++ {
		decs[i] = mains[i].Name
	}
	return decs, nil
}

func (s Storage) AllPbe() ([]string, error) {

	var pbes []pbe

	result := s.db.Model(&pbe{}).Select("name").Find(&pbes)
	if result.Error != nil {
		return nil, result.Error
	}

	decs := make([]string, len(pbes))
	for i := 0; i < len(pbes); i++ {
		decs[i] = pbes[i].Name
	}
	return decs, nil
}
