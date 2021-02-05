package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DbInit() *gorm.DB {
	dbConfig := settingDB()

	postgres_conn_name := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		dbConfig.db_host, dbConfig.db_port, dbConfig.db_username, dbConfig.db_name, dbConfig.db_password)
	fmt.Println("conname \n", postgres_conn_name)

	db, err := gorm.Open("postgres", postgres_conn_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database")

	return db
}
