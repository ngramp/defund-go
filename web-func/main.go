package main

import (
	"defund/components"
	"github.com/a-h/templ"
	"github.com/akrylysov/algnhsa"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

var isHostedOnAWS bool

func init() {
	task := len(os.Getenv("LAMBDA_TASK_ROOT"))
	ex := len(os.Getenv("AWS_EXECUTION_ENV"))
	isHostedOnAWS = task > 0 || ex > 0
}

func main() {
	r := chi.NewRouter()
	component := components.Home("World")
	r.Get("/", templ.Handler(component).ServeHTTP)
	if isHostedOnAWS {
		algnhsa.ListenAndServe(r, nil)
	} else {
		fs := http.FileServer(http.Dir("../static"))
		r.Handle("/css/*", http.StripPrefix("/css/", fs))
		http.ListenAndServe(":8080", r)
	}
}
