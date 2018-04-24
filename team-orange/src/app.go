package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)


func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var js = `{ "test" : 123}`
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	fmt.Fprintf(w, js)
}

func main() {

	router := httprouter.New()
	router.GET("/orange/api/", indexHandler)
	router.ServeFiles("/orange/static/*filepath", http.Dir("/go/src/app/static/"))
	router.GET("/orange/graph", renderHandler)
	
	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	log.Fatal(http.ListenAndServe(":3004", router))
}