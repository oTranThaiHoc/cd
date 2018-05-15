package handlers

import (
	"net/http"
	"fmt"
)

var EventHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	payload := r.PostFormValue("payload")
	fmt.Println(payload)

	w.WriteHeader(http.StatusOK)
})
