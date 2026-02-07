package controllers

import (
	"html/template"
	"net/http"

	"github.com/dagim-mante/golang-backend-course/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions := []struct {
			Question string
			Answer   template.HTML
		}{
			{
				Question: "Is this free?",
				Answer:   "Yes this is free forever.",
			},
			{
				Question: "What is your email?",
				Answer:   `<a href="mailto:dagimawimantefardo@gmail.com">dagimawimantefardo@gmail.com</a>`,
			},
		}
		tpl.Execute(w, questions)
	}
}
