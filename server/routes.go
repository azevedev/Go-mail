package main

import (
	"net/http"
)


func initRoutes(){
	http.HandleFunc("/", index) // setting router rule
    http.HandleFunc("/login", login)
}