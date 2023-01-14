package core

import (
	"context"
	"database/sql"
	"net/http"
)

type InterfaceModel interface {
	Save(...any)
	Create(...any)
}

type BaseModel struct {
	DBLink *sql.DB
	Ctx    context.Context
	Req    *http.Request
	Wri    http.ResponseWriter
	InterfaceModel
}

func (bm *BaseModel) RequestInit(db *sql.DB, ctx context.Context, wri http.ResponseWriter, req *http.Request) {
	bm.Req = req
	bm.Wri = wri
	bm.DBLink = db
	bm.Ctx = ctx
}
