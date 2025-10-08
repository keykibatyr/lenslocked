package controllers

import (
	"html/template"
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


func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct{
		Question string
		Answer template.HTML
	}{
	{
		Question: "Is there a free version?",
		Answer: " Yes! We offer a free trial for 30 days on any paid plans.",
	},
	{
		Question: "What are your support hours?",
		Answer: " We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
	},
	{
		Question: "How do I contact support?",
		Answer: `Email us – <a href="mailto:keykibatyr@gmail.com">support@lenslocked.com</a>`,
	},
	}

	return func(w http.ResponseWriter, r *http.Request){
        tpl.Data = questions
		tpl.ExecuteTemp(w)
	}
}