package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPackStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/pack/", controllers.CreatePack).Methods("POST")
	router.HandleFunc("/pack/", controllers.GetPack).Methods("GET")
	router.HandleFunc("/pack/{packId}", controllers.GetPackById).Methods("GET")
	router.HandleFunc("/pack/{packId}", controllers.UpdatePack).Methods("PUT")
	router.HandleFunc("/pack/{packId}", controllers.DeletePack).Methods("DELETE")
}
