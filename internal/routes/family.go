package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"log"
	"net/http"
)

func FamilyRouter(router *mux.Router, service service.Service) {
	subRouter := router.PathPrefix("/families").Subrouter()
	// swagger:operation POST /families createFamily
	//
	// Creates a new Family.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: family
	//   in: body
	//   description: family request
	//   schema:
	//     "$ref": "#/definitions/CreateFamilyRequest"
	//   required: true
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/CreateFamilyResponse"
	subRouter.HandleFunc("", CreateFamilyHandler(service)).Methods("POST")
	//subRouter.HandleFunc("", ProductsHandler).Methods("GET")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("GET")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("PUT")
	// swagger:operation DELETE /families/{id} createFamily
	//
	// Creates a new Family.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   required: true
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/DeleteFamilyResponse"
	subRouter.HandleFunc("/{id}", DeleteFamilyHandler(service)).Methods("DELETE")
}

func CreateFamilyHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := &service.CreateFamilyRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			WriteError(ctx, w, errors.Wrap(family.ErrorFamilyBadRequest, err))
			return
		}

		resp, err := svc.CreateFamily(ctx, req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func DeleteFamilyHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		id, _ := vars["id"]
		resp, err := svc.DeleteFamily(ctx, &service.DeleteFamilyRequest{Id: id})
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func CreateFamilyHandler2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
