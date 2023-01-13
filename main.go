package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"attrtour/database"
	"attrtour/routes"
)

var dbn *sql.DB
var ctx context.Context

func init() {
	var err error

	db := database.SQL{}
	ctx, dbn, err = db.ConnMySQL("rw_mysql_dbqueue")

	if err != nil {
		panic(fmt.Sprintf("Database connection error:\n%v", err))
	}

	api := routes.Api{
		DBLink: dbn,
		Ctx:    ctx,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Run(w, r)
	})
}

func main() {
	http.ListenAndServe(":9000", nil)
}
