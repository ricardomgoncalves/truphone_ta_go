package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"net/http"
	"strconv"
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
	// swagger:operation GET /members listMembers
	//
	// Lists members.
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
	// - name: family
	//   in: query
	//   description: family id to filter
	//   required: false
	//   type: string
	//   format: string
	// - name: parent
	//   in: query
	//   description: parent id to filter
	//   required: false
	//   type: string
	//   format: string
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/ListMembersResponse"
	subRouter.HandleFunc("", ListMembersHandler(service)).Methods("GET")
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
	// swagger:operation PUT /members/{id} updateMember
	//
	// Update a Member.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: member id
	//   required: true
	// - name: member
	//   in: body
	//   description: member request
	//   schema:
	//     "$ref": "#/definitions/UpdateMemberRequest"
	//   required: true
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/UpdateMemberResponse"
	subRouter.HandleFunc("/{id}", UpdateMemberHandler(service)).Methods("PUT")
	// swagger:operation DELETE /members/{id} deleteMember
	//
	// Delete a Member.
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
	//       "$ref": "#/definitions/DeleteMemberResponse"
	subRouter.HandleFunc("/{id}", DeleteMemberHandler(service)).Methods("DELETE")
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

func ListMembersHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := service.ListMembersRequest{}
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

		if familyId := r.URL.Query().Get("family"); familyId != "" {
			req.FamilyId = &familyId
		}

		if parentId := r.URL.Query().Get("family"); parentId != "" {
			req.ParentId = &parentId
		}

		resp, err := svc.ListMembers(ctx, &req)
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

func UpdateMemberHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		req := &service.UpdateMemberRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			WriteError(ctx, w, errors.Wrap(family.ErrorMemberBadRequest, err))
			return
		}

		id, _ := vars["id"]
		req.Id = id

		resp, err := svc.UpdateMember(ctx, req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func DeleteMemberHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		id, _ := vars["id"]
		resp, err := svc.DeleteMember(ctx, &service.DeleteMemberRequest{Id: id})
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}