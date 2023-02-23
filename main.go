package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"attrtour/core"
	"attrtour/database"
	"attrtour/routes"

	"attrtour/app/combines"

	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbn *sql.DB

var mux *http.ServeMux
var httpCors *cors.Cors

func init() {
	env := core.Helpers{}
	env.LoadEnv()

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

	isFindRoute := false
	api := routes.Api{
		DBLink:      gormDB,
		IsFindRoute: &isFindRoute,
	}

	mux = http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Run(w, r)
	})

	httpCors = cors.New(cors.Options{
		AllowedOrigins:   []string{env.Env("APP_HOST_ORIGIN")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Accept", "Accept-Language", "Content-Type", "X-Requested-With", "crossDomain"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
}

func main() {
	handler := httpCors.Handler(mux)
	http.ListenAndServe(":9000", handler)
}
