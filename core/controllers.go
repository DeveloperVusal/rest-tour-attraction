package core

import "net/http"

type BaseController struct {
	Req *http.Request
	Wri http.ResponseWriter
}
