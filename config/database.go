package config

import (
	"os"

	"github.com/bondhan/godddnews/internal/driver"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// NewDbConfig ...
func NewDbConfig() *gorm.DB {
	//init postgresql database
	// postgre := driver.NewDbDriver(os.Getenv("POSGRE_DSN"), "postgres")
	postgreDsn := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=" + os.Getenv("DB_SSLMODE")
	postgre := driver.NewDbDriver(postgreDsn, "postgres")
	db, err := postgre.ConnectDatabase()
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	return db
}
