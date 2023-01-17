package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"attrtour/database"
	"attrtour/routes"

	"attrtour/app/combines"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbn *sql.DB
var ctx context.Context

func init() {
	var err error

	db := database.SQL{}
	dbn, err = db.ConnMySQL()

	if err != nil {
		panic(fmt.Sprintf("Database connection error:\n%v", err))
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbn,
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Database connection with GORM error:\n%v", err))
	}

	models := combines.ModelCombine{}
	models.Build()

	gormDB.AutoMigrate(models.Items...)

	api := routes.Api{
		DBLink: gormDB,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Run(w, r)
	})
}

func main() {
	http.ListenAndServe(":9000", nil)
}
