package main

import (
	"net/http"
	"socialx/internal/store"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	app.store.Posts.Create(r.Context(), &store.Post{})
}
