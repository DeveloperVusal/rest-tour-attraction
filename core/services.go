package core

import (
	"net/http"

	"gorm.io/gorm"
)

type BaseService struct {
	DBLink *gorm.DB
	Req    *http.Request
	Wri    http.ResponseWriter
}

func (bs *BaseService) RequestInit(db *gorm.DB, wri http.ResponseWriter, req *http.Request) {
	bs.Req = req
	bs.Wri = wri
	bs.DBLink = db
}
