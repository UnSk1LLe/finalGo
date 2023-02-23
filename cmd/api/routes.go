package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createReplayHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showReplayHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateReplayHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteReplayHandler)
	return router
}
