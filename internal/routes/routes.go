package routes

import (
	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/middleware"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
)

func Router(service service.Service) *mux.Router {
	router := mux.
		NewRouter().
		PathPrefix("/truphone").
		Subrouter()

	// Middleware
	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.RequestIdMiddleware)

	// Routes
	FamilyRouter(router, service)
	MemberRouter(router, service)
	OperationsRouter(router, service)

	return router
}
