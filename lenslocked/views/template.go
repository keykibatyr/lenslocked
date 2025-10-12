package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func (t Template) ExecuteTemp(w http.ResponseWriter, data interface{}) { // I could add interafce instead of *Info. Probably will do that
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing: %v", err)
		http.Error(
			w,
			"Failed at executing the page",
			http.StatusInternalServerError,
		)
		return
	}

}

func ParseFileSys(fs fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		log.Printf("parsing template: %v", err)
		return Template{}, fmt.Errorf("could not parse the page %v", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// func Parse(filepath string) (Template, error) {
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		return Template{}, fmt.Errorf("could not parse the page %v", err)
// 	}
// 	return Template{
// 		htmlTpl: tpl,
// 		Data:    nil,
// 	}, nil
// }
