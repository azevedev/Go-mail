package main

import (
	"net/http"
)


func index(w http.ResponseWriter, r *http.Request){
	views.ExecuteTemplate(w, "index.html", nil)
}
func login(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		http.Redirect(w, r, "http://localhost:8080/", 200)
	}
	user := r.FormValue("username")
	pass := r.FormValue("password")
	data := struct{
		User string
		Password string
		Logged bool
	}{
		User: user,
		Password: pass,
		Logged: false,
	}
	if(user == "test" && pass == "123"){
		data.Logged = true
	}
	views.ExecuteTemplate(w, "login.html", data)
}