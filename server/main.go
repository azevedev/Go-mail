package main

import (
    "html/template"
	"net/http"
	"database/sql"
	// "fmt"
	"log"
)


var views *template.Template
var db *sql.DB
func init(){
	views = template.Must(template.ParseGlob("../client/views/*.html"))
	db = getCon()
	initRoutes()
}

func main() {
	
    err := http.ListenAndServe(":8080", nil) 
    if err != nil {
        log.Fatal("FatalErr: ", err)
	}
	
	defer db.Close()
}
