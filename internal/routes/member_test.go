package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateMemberHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.CreateMemberRequest{Member: family.Member{
			FirstName: "Name",
		}}

		svcResp := service.CreateMemberResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result:  "member_id",
		}

		svc.EXPECT().
			CreateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/members", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.CreateMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.CreateMemberRequest{Member: family.Member{
			FirstName: "Name",
		}}

		svcResp := service.CreateMemberResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result:  "member_id",
		}

		svc.EXPECT().
			CreateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/members", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.CreateMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.CreateMemberRequest{Member: family.Member{
			FirstName: "Name",
		}}

		svc.EXPECT().
			CreateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorMemberNotFound)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/members", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.CreateMemberRequest{Member: family.Member{
			FirstName: "Name",
		}}

		svc.EXPECT().
			CreateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/members", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("should get response code 400", func(t *testing.T) {
		svc.EXPECT().
			CreateMember(gomock.Any(), gomock.Any()).
			Times(0)

		body, err := json.Marshal("invalid")
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/members", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 400, recorder.Code)
	})
}

func TestListMembersHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		svcResp := service.ListMembersResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: []family.Member{
				{
					Id:        "id",
					FirstName: "Member",
				},
			},
		}

		svc.EXPECT().
			ListMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/members?limit=1&offset=1&family=id&parent=id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.ListMembersResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		svcResp := service.ListMembersResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: []family.Member{
				{
					Id:          "id",
					FirstName:        "Member",
				},
			},
		}

		svc.EXPECT().
			ListMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/members?limit=1&offset=1&family=id&parent=id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.ListMembersResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		svc.EXPECT().
			ListMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/members?limit=1&offset=1&family=id&parent=id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		svc.EXPECT().
			ListMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/members?limit=1&offset=1&family=id&parent=id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}

func TestGetMemberHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.GetMemberRequest{Id: "id"}

		svcResp := service.GetMemberResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: family.Member{
				Id: "id",
			},
		}

		svc.EXPECT().
			GetMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.GetMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.GetMemberRequest{Id: "id"}

		svcResp := service.GetMemberResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: family.Member{
				Id: "id",
			},
		}

		svc.EXPECT().
			GetMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.GetMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.GetMemberRequest{Id: "id"}

		svc.EXPECT().
			GetMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.GetMemberRequest{Id: "id"}

		svc.EXPECT().
			GetMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}

func TestUpdateMemberHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.UpdateMemberRequest{
			Id: "id",
			Member: family.Member{
				FirstName: "Member",
			},
		}

		svcResp := service.UpdateMemberResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: family.Member{
				Id:        "id",
				FirstName: "Member",
			},
		}

		svc.EXPECT().
			UpdateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/members/id", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.UpdateMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.UpdateMemberRequest{
			Id: "id",
			Member: family.Member{
				FirstName: "Member",
			},
		}

		svcResp := service.UpdateMemberResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: family.Member{
				Id:        "id",
				FirstName: "Member",
			},
		}

		svc.EXPECT().
			UpdateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/members/id", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.UpdateMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.UpdateMemberRequest{
			Id: "id",
			Member: family.Member{
				FirstName: "Member",
			},
		}

		svc.EXPECT().
			UpdateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorMemberNotFound)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/members/id", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.UpdateMemberRequest{
			Id: "id",
			Member: family.Member{
				FirstName: "Member",
			},
		}

		svc.EXPECT().
			UpdateMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/members/id", bytes.NewBuffer(body))
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("should get response code 400", func(t *testing.T) {
		svc.EXPECT().
			UpdateMember(gomock.Any(), gomock.Any()).
			Times(0)

		body, err := json.Marshal("invalid")
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/members/id", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 400, recorder.Code)
	})
}

func TestDeleteMemberHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.DeleteMemberRequest{Id: "id"}

		svcResp := service.DeleteMemberResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
		}

		svc.EXPECT().
			DeleteMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("DELETE", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.DeleteMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.DeleteMemberRequest{Id: "id"}

		svcResp := service.DeleteMemberResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
		}

		svc.EXPECT().
			DeleteMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("DELETE", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.DeleteMemberResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.DeleteMemberRequest{Id: "id"}

		svc.EXPECT().
			DeleteMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("DELETE", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.DeleteMemberRequest{Id: "id"}

		svc.EXPECT().
			DeleteMember(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("DELETE", "/truphone/members/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}
