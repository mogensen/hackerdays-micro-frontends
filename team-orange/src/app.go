package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"os/signal"
	"github.com/braintree/manners"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var js = `{ "test" : 123}`
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, js)
}

func main() {
	
	shutdown := make(chan int)
	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	router := httprouter.New()
	router.GET("/orange/api/weather", getWeatherDataHandler)
	
	router.GET("/orange/api/", indexHandler)
	router.ServeFiles("/orange/static/*filepath", http.Dir("/go/src/app/static/"))
	
	router.GET("/orange/graph/", renderHandler)
	router.GET("/orange/graph/:type", renderHandler)

	router.GET("/orange/iot/", iotHandler)
	router.GET("/orange/iot/:type", iotHandler)

	server := manners.NewWithServer(&http.Server{Addr: ":3004", Handler: router})
	go func() {		
		fmt.Println("Http server up and running")
		server.ListenAndServe()
		shutdown <- 1
	}()

	//register for interupt (Ctrl+C) and SIGTERM (docker)
	signal.Notify(sigChan)
	go func() {
		fmt.Println("Waiting for shut down...")
		<-sigChan
		fmt.Println("Shutting down...")
		server.Close()
	}()

	<-shutdown
}