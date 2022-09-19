package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoundStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/round/", controllers.CreateRound).Methods("POST")
	router.HandleFunc("/round/", controllers.GetRound).Methods("GET")
	router.HandleFunc("/round/{roundId}", controllers.GetRoundById).Methods("GET")
	router.HandleFunc("/round/{roundId}", controllers.UpdateRound).Methods("PUT")
	router.HandleFunc("/round/{roundId}", controllers.DeleteRound).Methods("DELETE")
}
