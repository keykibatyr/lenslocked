package controllers

import (
	"net/http"

	"github.com/keykibatyr/lenslocked/views"
)

func StaticHandler(tpl views.Template, getData func(r *http.Request) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if getData != nil {
            tpl.Data = getData(r)
        }
		tpl.ExecuteTemp(w)
	}
}
