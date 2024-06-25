package controllers

import (
	"net/http"
)

type Controller struct {}

func New() (*Controller) {
	return &Controller{}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api" {
		c.getResponse(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func (c *Controller) getResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Go API Server has started!"))
}
