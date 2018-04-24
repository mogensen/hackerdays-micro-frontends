package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var js = `/* eslint-disable no-use-before-define, no-console, class-methods-use-this */
	/* globals HTMLElement, window, CustomEvent 123*/
	
	class OrangeRecos extends HTMLElement {
	  static get observedAttributes() {
		return ['sku'];
	  }
	  connectedCallback() {
		const sku = this.getAttribute('sku');
		this.log('connected', sku);
		this.render();
	  }
	  attributeChangedCallback(attr, oldValue, newValue) {
		this.log('attributeChanged', attr, newValue);
		this.render();
	  }
	  render() {
		const sku = this.getAttribute('sku');
		this.innerHTML = '<h2>TEST</h1>';
	  }
	  disconnectedCallback() {
		const sku = this.getAttribute('sku');
		this.log('disconnected', sku);
	  }
	  log(...args) {
		console.log('🖼️ orange-recos', ...args);
	  }
	}
	
	window.customElements.define('orange-recos', OrangeRecos);
	`
	w.Header().Set("Content-Type", "application/javascript; charset=UTF-8")

	fmt.Fprintf(w, js)
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/orange/fragments.js", indexHandler)
	
	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":3004", router)
}