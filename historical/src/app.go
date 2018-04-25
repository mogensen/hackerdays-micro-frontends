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
	
	router.ServeFiles("/historical/static/*filepath", http.Dir("/go/src/app/static/"))
	
	router.GET("/historical/iot/", iotHandler)
	router.GET("/historical/iot/:type", iotHandler)

	server := manners.NewWithServer(&http.Server{Addr: ":3006", Handler: router})
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