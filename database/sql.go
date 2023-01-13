package database

import (
	"context"
	"database/sql"

	"attrtour/config"

	_ "github.com/go-sql-driver/mysql"
)

type SQL struct{}

// Метод подключения к БД MySQL
func (_sql *SQL) ConnMySQL(serverName string) (context.Context, *sql.DB, error) {
	loadCfg := &config.DatabaseConfig{}
	cfg := loadCfg.ConfigLoad()

	var dataSourceName string = ""

	dataSourceName = cfg["username"] + ":" + cfg["password"] + "@tcp(" + cfg["host"] + ":" + cfg["port"] + ")/" + cfg["dbname"]

	_db, err := _sql.DriverMySQL(&dataSourceName)

	_db.SetMaxOpenConns(30)

	return context.Background(), _db, err
}

// Вызов драйвера mysql
func (_sql *SQL) DriverMySQL(dataSourceame *string) (*sql.DB, error) {
	return sql.Open("mysql", *dataSourceame)
}
