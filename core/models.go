package core

import (
	"net/http"

	"gorm.io/gorm"
)

type InterfaceModel interface {
	Save(...any)
	Create(...any)
}

type BaseModel struct {
	DBLink *gorm.DB
	Req    *http.Request
	Wri    http.ResponseWriter
	InterfaceModel
}
