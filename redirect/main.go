package main

import (
	"net/http"
	"redirector/controllers"
)

func main() {
	http.HandleFunc("/", controllers.RedirectLink)

	http.ListenAndServe(":9900", nil)
}
