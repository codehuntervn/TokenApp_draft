package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserWalletMTDStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/userWalletMTD/", controllers.CreateUserWalletMTD).Methods("POST")
	router.HandleFunc("/userWalletMTD/", controllers.GetUserWalletMTD).Methods("GET")
	router.HandleFunc("/userWalletMTD/{userAddress}", controllers.GetUserWalletMTDById).Methods("GET")
	router.HandleFunc("/userWalletMTD/{userAddress}", controllers.UpdateUserWalletMTD).Methods("PUT")
	router.HandleFunc("/userWalletMTD/{userAddress}", controllers.DeleteUserWalletMTD).Methods("DELETE")
}
