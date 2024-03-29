package service

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name	string
	Method	string
	Pattern	string
	HandlerFunc	http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs
type Routes	[]Route

// Initialize our routes
var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		func (w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application.json; charset=UTF-8")
				params := mux.Vars(r)
				accountId := params["accountId"]
				account := DBClient.QueryAccount(accountId)
				print(account)
				w.Write(account)
		},
	},
	Route{
		"UpdateAccount",
		"PUT",
		"/accounts/{accountId}",
		func (w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application.json; charset=UTF-8")
				params := mux.Vars(r)
				accountId := params["accountId"]
				account, err := DBClient.UpdateAccount(accountId)
				if err != nil {
					log.Fatal(err)
				} else {
					w.Write(account)
				}
		},
	},
}
