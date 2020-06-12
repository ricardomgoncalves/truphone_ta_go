package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"net/http"
)

func MemberRouter(router *mux.Router, service service.Service) {
	subRouter := router.PathPrefix("/members").Subrouter()
	// swagger:operation POST /members createMembers
	//
	// Creates a new Member.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: member
	//   in: body
	//   description: member request
	//   schema:
	//     "$ref": "#/definitions/CreateMemberRequest"
	//   required: true
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/CreateMemberResponse"
	subRouter.HandleFunc("", CreateMemberHandler(service)).Methods("POST")
	//subRouter.HandleFunc("", ProductsHandler).Methods("GET")
	// swagger:operation POST /members/{id} getMember
	//
	// Gets a Member.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: member id
	//   required: true
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/GetMemberResponse"
	subRouter.HandleFunc("/{id}", GetMemberHandler(service)).Methods("GET")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("PUT")
	//subRouter.HandleFunc("/{id}", ProductsHandler).Methods("DELETE")
}

func CreateMemberHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := &service.CreateMemberRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			WriteError(ctx, w, errors.Wrap(family.ErrorMemberBadRequest, err))
			return
		}

		resp, err := svc.CreateMember(ctx, req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func GetMemberHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		id, _ := vars["id"]
		resp, err := svc.GetMember(ctx, &service.GetMemberRequest{Id: id})
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}