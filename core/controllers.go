package core

import (
	"context"
	"database/sql"
	"net/http"
)

type BaseController struct {
	DBLink *sql.DB
	Ctx    context.Context
	Req    *http.Request
	Wri    http.ResponseWriter
}
