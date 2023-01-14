package core

import (
	"fmt"
	"net/http"
)

type BaseModel struct {
	Req *http.Request
	Wri http.ResponseWriter
}

func (bm *BaseModel) RequestInit(wri http.ResponseWriter, req *http.Request) {
	bm.Req = req
	bm.Wri = wri
}

func (bm *BaseModel) Save(a ...any) {
	fmt.Println(a...)
}
