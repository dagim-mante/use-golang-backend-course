package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

func Parse(filename string) (Template, error) {
	tpl, err := template.ParseFiles(filename)
	if err != nil {
		return Template{}, fmt.Errorf("error parsing: %w\n", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing the template: %v\n", err)
		http.Error(w, "Failed to execute the template.", http.StatusInternalServerError)
		return
	}
}
