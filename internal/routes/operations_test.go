package routes

import (
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

func TestListAccumulatedFamiliesHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		svcResp := service.ListAccumulatedFamiliesResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: []family.AgeFamily{
				{
					Family: family.Family{
						Id:          "id",
						Name:        "Family 1",
						CountryCode: "PT",
					},
					Age: 1,
				},
			},
		}

		svc.EXPECT().
			ListAccumulatedFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/accumulators?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.ListAccumulatedFamiliesResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		svcResp := service.ListAccumulatedFamiliesResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: []family.AgeFamily{
				{
					Family: family.Family{
						Id:          "id",
						Name:        "Family 1",
						CountryCode: "PT",
					},
					Age: 1,
				},
			},
		}

		svc.EXPECT().
			ListAccumulatedFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/accumulators?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.ListAccumulatedFamiliesResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		svc.EXPECT().
			ListAccumulatedFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/accumulators?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		svc.EXPECT().
			ListAccumulatedFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/accumulators?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}

func TestListFastestGrowingFamiliesHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		svcResp := service.ListFastestGrowingFamiliesResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: []family.AgeFamily{
				{
					Family: family.Family{
						Id:          "id",
						Name:        "Family 1",
						CountryCode: "PT",
					},
					Age: 1,
				},
			},
		}

		svc.EXPECT().
			ListFastestGrowingFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/growths?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.ListFastestGrowingFamiliesResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		svcResp := service.ListFastestGrowingFamiliesResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: []family.AgeFamily{
				{
					Family: family.Family{
						Id:          "id",
						Name:        "Family 1",
						CountryCode: "PT",
					},
					Age: 1,
				},
			},
		}

		svc.EXPECT().
			ListFastestGrowingFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/growths?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.ListFastestGrowingFamiliesResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		svc.EXPECT().
			ListFastestGrowingFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/growths?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		svc.EXPECT().
			ListFastestGrowingFamilies(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/growths?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}

func TestListPossibleDuplicatesMembersHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	svc := service.NewMockService(ctl)

	t.Run("should get response code 200", func(t *testing.T) {
		svcResp := service.ListPossibleDuplicatesMembersResponse{
			Id:      "request_id",
			Code:    200,
			Message: http.StatusText(200),
			Result: []family.Member{
				{
					Id: "id",
				},
			},
		}

		svc.EXPECT().
			ListPossibleDuplicatesMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/duplicates?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		outResp := service.ListPossibleDuplicatesMembersResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get response code of service", func(t *testing.T) {
		svcResp := service.ListPossibleDuplicatesMembersResponse{
			Id:      "request_id",
			Code:    301,
			Message: http.StatusText(301),
			Result: []family.Member{
				{
					Id: "id",
				},
			},
		}

		svc.EXPECT().
			ListPossibleDuplicatesMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(&svcResp, nil)

		req, err := http.NewRequest("GET", "/truphone/duplicates?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 301, recorder.Code)
		outResp := service.ListPossibleDuplicatesMembersResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResp)
		require.Nil(t, err)
		assert.Equal(t, svcResp, outResp)
	})

	t.Run("should get service error code", func(t *testing.T) {
		svc.EXPECT().
			ListPossibleDuplicatesMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, family.ErrorFamilyNotFound)

		req, err := http.NewRequest("GET", "/truphone/duplicates?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("should get 500 error code", func(t *testing.T) {
		svc.EXPECT().
			ListPossibleDuplicatesMembers(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, errors.New("random"))

		req, err := http.NewRequest("GET", "/truphone/duplicates?limit=1&offset=1", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		router := Router(svc)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, 500, recorder.Code)
	})
}
