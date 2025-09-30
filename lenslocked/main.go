package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my Website!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch email me at <a href= \"mailto:keykibatyr@gmail.com\">keykibatyr@gmail.com</a> ")
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		w.WriteHeader(404)
// 		fmt.Fprintln(w, "page not found")
// 	}
// }

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		w.WriteHeader(404)
// 		fmt.Fprintln(w, "page not found")
// 	}
// }

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
	fmt.Fprint(w, `<h2><pre>
Q: Is there a free version?
A: Yes! We offer a free trial for 30 days on any paid plans.

Q: What are your support hours?
A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.

Q: How do I contact support?
A: Email us – <a href="mailto:keykibatyr@gmail.com">support@lenslocked.com</a>
</pre></h2>`)

}

func main() {
	// var router Router
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(
			w,
			"page not found",
			404,
		)
	})
	fmt.Println("Listening to port :7000...")
	http.ListenAndServe(":7000", r)
}
