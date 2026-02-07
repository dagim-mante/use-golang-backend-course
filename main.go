package main

import (
	"fmt"
	"net/http"

	"github.com/dagim-mante/golang-backend-course/controllers"
	"github.com/dagim-mante/golang-backend-course/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", controllers.StaticHandler(views.Must(views.Parse("templates/home.gohtml"))))

	router.Get("/contact", controllers.StaticHandler(views.Must(views.Parse("templates/contact.gohtml"))))

	router.Get("/faq", controllers.StaticHandler(views.Must(views.Parse("templates/faq.gohtml"))))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", router)
}
