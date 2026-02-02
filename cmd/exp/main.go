package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  template.HTML
}

func main() {
	t, err := template.ParseFiles("home.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Dagim",
		Bio:  `<script>alert("you have been hacked!")</script>`,
	}

	err = t.Execute(os.Stdout, user)
}
