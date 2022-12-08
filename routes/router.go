package routes

import (
	"crud-mux-example/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", controllers.Create).Methods("POST")
	router.HandleFunc("/read", controllers.Read).Methods("GET")
	router.HandleFunc("/update", controllers.Update).Methods("PUT")
	router.HandleFunc("/delete", controllers.Delete).Methods("DELETE")
	return router
}
