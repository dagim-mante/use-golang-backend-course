package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dagim-mante/golang-backend-course/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := views.Parse(filePath)
	if err != nil {
		log.Printf("Parsing the template: %v\n", err)
		http.Error(w, "Failed to parse the template.", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.gohtml")
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", router)
}
