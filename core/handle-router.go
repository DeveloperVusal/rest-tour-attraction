package core

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type HandleRouter struct {
	Req *http.Request
	Wri http.ResponseWriter
}

func (hr *HandleRouter) TunnelControl(routes *map[string][]interface{}) {
	isNotFound := true

	for path, interf := range *routes {
		hPath, paramsMap, paramsSlice := hr.HandlePath(&path)

		if hr.IsValidPath(hPath, paramsSlice) {
			isNotFound = false
			method := interf[0].(string)

			if hr.IsMethod(method) {
				var appendArgs = map[string]interface{}{}

				for key, val := range paramsMap {
					appendArgs[key] = val
				}

				f_ptr, ok := interf[1].(func(map[string]interface{}))

				if !ok {
					hr.Wri.WriteHeader(http.StatusInternalServerError)
					panic(fmt.Sprintf("Invalid object type: expected `func()`, turned out to be `%T`", interf[1]))
				}

				f_ptr(appendArgs)
				hr.Wri.WriteHeader(http.StatusOK)
			} else {
				hr.Wri.WriteHeader(http.StatusMethodNotAllowed)
				hr.Wri.Write([]byte("Method Not Allowed\n"))
				hr.Wri.Write([]byte(method + ", must be passed\n"))
			}
		}
	}

	if isNotFound {
		hr.Wri.WriteHeader(http.StatusNotFound)
		hr.Wri.Write([]byte("404 Not Found\n"))
	}
}

func (hr *HandleRouter) IsMethod(method string) bool {
	if hr.Req.Method == method {
		return true
	} else {
		return false
	}
}

func (hr *HandleRouter) IsValidPath(path string, params []string) bool {
	concat := ""

	for _, val := range params {
		concat = concat + "/" + val
	}

	if len(concat) > 1 {
		if concat[0:1] == "/" {
			concat = concat[1:]
		}
	}

	hPath := path + concat

	if hr.Req.URL.Path == hPath {
		return true
	} else {
		return false
	}
}

func (hr *HandleRouter) HandlePath(path *string) (string, map[string]string, []string) {
	var params map[int]string = map[int]string{}
	var paramsSlice []string
	var paramsMap map[string]string = map[string]string{}
	var pathBuild []string

	sections1 := strings.Split(hr.Req.URL.Path, "/")

	for indx, item := range sections1 {
		params[indx] = item
	}

	sections2 := strings.Split(*path, "/")
	r, err := regexp.Compile("^{.*}$")

	if err != nil {
		hr.Wri.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	for indx, item := range sections2 {
		if r.MatchString(item) {
			item = item[:len(item)-1]
			item = item[1:]
			paramsMap[item] = params[indx]
			paramsSlice = append(paramsSlice, params[indx])
		} else {
			pathBuild = append(pathBuild, item)
		}
	}

	return strings.Join(pathBuild, "/"), paramsMap, paramsSlice
}
