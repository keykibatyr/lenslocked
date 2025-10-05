package views

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

type Template struct {
	htmlTpl *template.Template
	Data interface{}
}

func (t *Template) ExecuteTemp(w http.ResponseWriter) { // I could add interafce instead of *Info. Probably will do that
	err := t.htmlTpl.Execute(w, t.Data)
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

func Parse(filepath string, data interface{}) (Template, error){
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		return Template{}, fmt.Errorf("could not parse the page")
	}
	return Template{
		htmlTpl: tpl,
		Data: data,
	}, nil
}

