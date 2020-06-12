package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
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
	// swagger:operation GET /families listFamilies
	//
	// Lists Families.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: offset
	//   in: query
	//   description: offset number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// - name: limit
	//   in: query
	//   description: maximum number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// - name: country
	//   in: query
	//   description: country code to filter
	//   required: false
	//   type: string
	//   format: string
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/ListFamiliesResponse"
	subRouter.HandleFunc("", ListFamiliesHandler(service)).Methods("GET")
	// swagger:operation GET /families/{id} getFamily
	//
	// Gets a Family.
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
	//       "$ref": "#/definitions/GetFamilyResponse"
	subRouter.HandleFunc("/{id}", GetFamilyHandler(service)).Methods("GET")
	//subRouter.HandleFunc("", ProductsHandler).Methods("GET")
	// swagger:operation PUT /families/{id} updatedFamily
	//
	// Updates a Family.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   required: true
	// - name: family
	//   in: body
	//   description: family request
	//   schema:
	//     "$ref": "#/definitions/UpdateFamilyRequest"
	//   required: true
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/UpdateFamilyResponse"
	subRouter.HandleFunc("/{id}", UpdateFamilyHandler(service)).Methods("PUT")
	// swagger:operation DELETE /families/{id} deleteFamily
	//
	// Deletes a Family.
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

func ListFamiliesHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := service.ListFamiliesRequest{}
		if limit := r.URL.Query().Get("limit"); limit != "" {
			if val, err := strconv.Atoi(limit); err == nil {
				o := uint32(val)
				req.Limit = &o
			}
		}

		if offset := r.URL.Query().Get("offset"); offset != "" {
			if val, err := strconv.Atoi(offset); err == nil {
				o := uint32(val)
				req.Offset = &o
			}
		}

		if countryCode := r.URL.Query().Get("country"); countryCode != "" {
			req.CountryCode = &countryCode
		}

		resp, err := svc.ListFamilies(ctx, &req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func GetFamilyHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		id, _ := vars["id"]
		resp, err := svc.GetFamily(ctx, &service.GetFamilyRequest{Id: id})
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func UpdateFamilyHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		req := &service.UpdateFamilyRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			WriteError(ctx, w, errors.Wrap(family.ErrorFamilyBadRequest, err))
			return
		}

		id, _ := vars["id"]
		req.Id = id

		resp, err := svc.UpdateFamily(ctx, req)
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
