package database

import (
	"database/sql"

	"attrtour/config"

	_ "github.com/go-sql-driver/mysql"
)

type SQL struct{}

// Метод подключения к БД MySQL
func (_sql *SQL) ConnMySQL() (*sql.DB, error) {
	loadCfg := &config.DatabaseConfig{}
	cfg := loadCfg.ConfigLoad()

	var dataSourceName string = ""

	dataSourceName = cfg["username"] + ":" + cfg["password"] + "@tcp(" + cfg["host"] + ":" + cfg["port"] + ")/" + cfg["dbname"] + "?parseTime=true"

	_db, err := _sql.DriverMySQL(&dataSourceName)

	_db.SetMaxOpenConns(100)

	return _db, err
}

// Вызов драйвера mysql
func (_sql *SQL) DriverMySQL(dataSourceame *string) (*sql.DB, error) {
	return sql.Open("mysql", *dataSourceame)
}
