package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/braintree/manners"
)


func main() {
	
	shutdown := make(chan int)
	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	router := httprouter.New()
	router.GET("/orange/api/weather", getWeatherDataHandler)
	
	router.ServeFiles("/orange/static/*filepath", http.Dir("/go/src/app/static/"))
	
	router.GET("/orange/graph/", renderHandler)
	router.GET("/orange/graph/:type", renderHandler)

	server := manners.NewWithServer(&http.Server{Addr: ":3004", Handler: router})
	go func() {		
		fmt.Println("Http server up and running")
		server.ListenAndServe()
		shutdown <- 1
	}()

	//register for interupt (Ctrl+C) and SIGTERM (docker)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("Shutting down....")

		server.Close()
	}()

	<-shutdown
}