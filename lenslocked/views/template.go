package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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

func (t Template) ExecuteTemp(w http.ResponseWriter, r *http.Request, data interface{}) { // I could add interafce instead of *Info. Probably will do that
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "could not clone thhe template", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing: %v", err)
		http.Error(
			w,
			"Failed at executing the page",
			http.StatusInternalServerError,
		)
		return
	}
	io.Copy(w, &buf)

}

func ParseFileSys(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("could not implement the csrf")
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
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
