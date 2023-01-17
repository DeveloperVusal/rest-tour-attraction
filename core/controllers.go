package core

import (
	"net/http"

	"gorm.io/gorm"
)

type BaseController struct {
	DBLink *gorm.DB
	Req    *http.Request
	Wri    http.ResponseWriter
}
