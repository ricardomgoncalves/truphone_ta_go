package routes

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
)

func OperationsRouter(router *mux.Router, service service.Service) {
	// swagger:operation GET /accumulators listAccumulatedFamilies
	//
	// Lists the most accumulated aged Families.
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
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/ListAccumulatedFamiliesResponse"
	router.HandleFunc("/accumulators", ListAccumulatedFamiliesHandler(service)).Methods("GET")
	// swagger:operation GET /growths listGrowingFamilies
	//
	// Lists the fastest growing Families.
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
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/ListFastestGrowingFamiliesResponse"
	router.HandleFunc("/growths", ListFastestGrowingFamiliesHandler(service)).Methods("GET")
	// swagger:operation GET /duplicates listDuplicates
	//
	// Lists the possible duplicated members.
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
	// responses:
	//   default:
	//     schema:
	//       "$ref": "#/definitions/ListPossibleDuplicatesMembersResponse"
	router.HandleFunc("/duplicates", ListPossibleDuplicatesMembersHandler(service)).Methods("GET")
}

func ListAccumulatedFamiliesHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := service.ListAccumulatedFamiliesRequest{}
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

		resp, err := svc.ListAccumulatedFamilies(ctx, &req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func ListFastestGrowingFamiliesHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := service.ListFastestGrowingFamiliesRequest{}
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

		resp, err := svc.ListFastestGrowingFamilies(ctx, &req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}

func ListPossibleDuplicatesMembersHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := service.ListPossibleDuplicatesMembersRequest{}
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

		resp, err := svc.ListPossibleDuplicatesMembers(ctx, &req)
		if err != nil {
			WriteError(ctx, w, err)
			return
		}

		Write(ctx, w, resp.Code, resp)
	}
}
