package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Write([]byte("ok"))
}
