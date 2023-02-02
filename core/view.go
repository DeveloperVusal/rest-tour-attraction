package core

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type View struct{}

func (vw *View) Render(w http.ResponseWriter, filename string) {
	fileHtml := vw.getFile(filename)

	fileHtml.Execute(w, nil)
}

func (vw *View) getFile(filename string) template.Template {
	ex, _ := os.Executable()
	filepathJoin := filepath.Join(path.Dir(ex), filename)

	tmpl, _ := template.ParseFiles(filepathJoin)

	return *tmpl
}
