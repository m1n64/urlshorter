package controllers

import (
	"fmt"
	"net/http"
	"redirector/services"
)

func RedirectLink(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	paramValues, found := queryParams["link"]
	if !found || len(paramValues) < 1 {
		fmt.Fprint(w, "Url Param 'link' is missing")
		return
	}

	if paramValues[0] == "" {

	}

	services.PublishQueue(paramValues[0])

	link := services.ListenLink()

	if link != "" {
		http.Redirect(w, r, link, http.StatusFound)
	}

	// Выводим значение параметра на страницу
	fmt.Fprintf(w, "Url Param 'link' is: %s", paramValues[0])
}
