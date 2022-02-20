package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//Home is the Home Page Handle
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "index.html")
}

// About is the About Page Handle
func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+ tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return 
	}
}