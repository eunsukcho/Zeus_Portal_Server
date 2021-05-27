package dblayer

import (
	"fmt"
	"zeus/models"

	"github.com/jinzhu/gorm"
)

type DBLayer interface {
	// init zeus env
	GetAllEnvData() ([]models.Env_setting_Tbls, error)
	UpdateEnvData(models.Env_setting_Tbls) (models.Env_setting_Tbls, error)

	GetLogCode() ([]models.LogType_Code, error)

	MenuLayer
	SmtpLayer
	AuthLayer
	DevServerLayer
}

type DBORM struct {
	*gorm.DB
}

func NewDBInit() (*DBORM, error) {
	dbConfig := settingDB()

	postgres_conn_name := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		dbConfig.db_host, dbConfig.db_port, dbConfig.db_username, dbConfig.db_name, dbConfig.db_password)
	fmt.Println("conname \n", postgres_conn_name)

	db, err := gorm.Open("postgres", postgres_conn_name)

	return &DBORM{
		DB: db,
	}, err
}
