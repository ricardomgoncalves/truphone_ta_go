package routes

import (
	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
)

func MemberRouter(router *mux.Router, service service.Service) {
	//subRouter := router.PathPrefix("/member").Subrouter()
	//subRouter.HandleFunc("", ProductsHandler).Methods("POST")
	//subRouter.HandleFunc("", ProductsHandler).Methods("GET")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("GET")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("PUT")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("DELETE")
}
