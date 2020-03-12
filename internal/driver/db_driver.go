package driver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //To bypass mysql error
	_ "github.com/lib/pq"
)

// DBDriver ...
type DBDriver struct {
	Dsn  string
	Name string
}

// NewDbDriver ...
func NewDbDriver(dsn string, name string) *DBDriver {
	return &DBDriver{
		Dsn:  dsn,
		Name: name,
	}
}

//ConnectDatabase will create connection to database
func (m *DBDriver) ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(m.Name, m.Dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
