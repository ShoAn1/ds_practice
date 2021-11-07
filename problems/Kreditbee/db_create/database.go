package db_create

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db_conncet_object *gorm.DB
var Once sync.Once

func initialiseDb() (*gorm.DB, error) {
	//connection_string := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=asdf"
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	fmt.Println(db_username, db_password, db_name)
	dsn := fmt.Sprintf(`%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local`, db_username, db_password, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDb_instance() *gorm.DB {
	var err error
	if db_conncet_object == nil {
		Once.Do(func() {
			db_conncet_object, err = initialiseDb()
		})
		if err != nil {
			log.Fatal("error in connecting to database")
		}
	}
	return db_conncet_object
}
