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

func TestCreateFamilyHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.CreateFamilyRequest{Family: family.Family{
			Name:        "Family 1",
			CountryCode: "PT",
		}}

		svcResp := service.CreateFamilyResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result:  "family_id",
		}

		svc.EXPECT().
			CreateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/families", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.CreateFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.CreateFamilyRequest{Family: family.Family{
			Name:        "Family 1",
			CountryCode: "PT",
		}}

		svcResp := service.CreateFamilyResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result:  "family_id",
		}

		svc.EXPECT().
			CreateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/families", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.CreateFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.CreateFamilyRequest{Family: family.Family{
			Name:        "Family 1",
			CountryCode: "PT",
		}}

		svc.EXPECT().
			CreateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/families", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.CreateFamilyRequest{Family: family.Family{
			Name:        "Family 1",
			CountryCode: "PT",
		}}

		svc.EXPECT().
			CreateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/families", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("should get response code 400", func(t *testing.T) {
		svc.EXPECT().
			CreateFamily(gomock.Any(), gomock.Any()).
			Times(0)

		body, err := json.Marshal("invalid")
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/truphone/families", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 400, recorder.Code)
	})
}

func TestGetFamilyHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.GetFamilyRequest{Id: "id"}

		svcResp := service.GetFamilyResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: family.Family{
				Id:          "id",
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svc.EXPECT().
			GetFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.GetFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.GetFamilyRequest{Id: "id"}

		svcResp := service.GetFamilyResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: family.Family{
				Id:          "id",
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svc.EXPECT().
			GetFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.GetFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.GetFamilyRequest{Id: "id"}

		svc.EXPECT().
			GetFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.GetFamilyRequest{Id: "id"}

		svc.EXPECT().
			GetFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}

func TestListFamiliesHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		svcResp := service.ListFamiliesResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: []family.Family{
				{
					Id:          "id",
					Name:        "Family 1",
					CountryCode: "PT",
				},
			},
		}

		svc.EXPECT().
			ListFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/families?limit=1&offset=1&country=PT", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.ListFamiliesResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		svcResp := service.ListFamiliesResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: []family.Family{
				{
					Id:          "id",
					Name:        "Family 1",
					CountryCode: "PT",
				},
			},
		}

		svc.EXPECT().
			ListFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/families?limit=1&offset=1&country=PT", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.ListFamiliesResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		svc.EXPECT().
			ListFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/families?limit=1&offset=1&country=PT", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		svc.EXPECT().
			ListFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/families?limit=1&offset=1&country=PT", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}

func TestUpdateFamilyHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svcResp := service.UpdateFamilyResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: family.Family{
				Id:          "id",
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svc.EXPECT().
			UpdateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/families/id", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.UpdateFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svcResp := service.UpdateFamilyResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: family.Family{
				Id:          "id",
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svc.EXPECT().
			UpdateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/families/id", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.UpdateFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svc.EXPECT().
			UpdateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/families/id", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Name:        "Family 1",
				CountryCode: "PT",
			},
		}

		svc.EXPECT().
			UpdateFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/families/id", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("should get response code 400", func(t *testing.T) {
		svc.EXPECT().
			UpdateFamily(gomock.Any(), gomock.Any()).
			Times(0)

		body, err := json.Marshal("invalid")
		require.Nil(t, err)

		req, err := http.NewRequest("PUT", "/truphone/families/id", bytes.NewBuffer(body))

		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 400, recorder.Code)
	})
}

func TestDeleteFamilyHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		reqBody := service.DeleteFamilyRequest{Id: "id"}

		svcResp := service.DeleteFamilyResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
		}

		svc.EXPECT().
			DeleteFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("DELETE", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.DeleteFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		reqBody := service.DeleteFamilyRequest{Id: "id"}

		svcResp := service.DeleteFamilyResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
		}

		svc.EXPECT().
			DeleteFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("DELETE", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.DeleteFamilyResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		reqBody := service.DeleteFamilyRequest{Id: "id"}

		svc.EXPECT().
			DeleteFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("DELETE", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		reqBody := service.DeleteFamilyRequest{Id: "id"}

		svc.EXPECT().
			DeleteFamily(gomock.Any(), gomock.Eq(&reqBody)).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("DELETE", "/truphone/families/id", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}
