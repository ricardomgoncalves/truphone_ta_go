package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/requestid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFamilyService(t *testing.T) {
	t.Run("should return service with repos", func(t *testing.T) {
		service, err := NewFamilyService(nil, nil)
		require.Nil(t, err)
		assert.Equal(t, nil, service.familyRepo)
		assert.Equal(t, nil, service.memberRepo)

		ctl := gomock.NewController(t)
		famRepo := repo.NewMockFamilyRepo(ctl)
		memRepo := repo.NewMockMemberRepo(ctl)

		service, err = NewFamilyService(famRepo, memRepo)
		require.Nil(t, err)
		assert.Equal(t, famRepo, service.familyRepo)
		assert.Equal(t, memRepo, service.memberRepo)
	})
}

func TestFamilyService_CreateFamily(t *testing.T) {
	ctl := gomock.NewController(t)
	famRepo := repo.NewMockFamilyRepo(ctl)
	memRepo := repo.NewMockMemberRepo(ctl)
	service, err := NewFamilyService(famRepo, memRepo)
	require.Nil(t, err)
	ctx := requestid.WithRequestId(context.Background(), "request_id")

	t.Run("should return error name should have a character different of a space", func(t *testing.T) {
		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "        ",
			CountryCode: "PT",
		}})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "name should have a character different of a space"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should be at least 3 characters", func(t *testing.T) {
		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "aa",
			CountryCode: "PT",
		}})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at least 3 characters"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should be at maximum 30 characters", func(t *testing.T) {
		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "this name is too long to be in a family name",
			CountryCode: "PT",
		}})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at maximum 30 characters"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error invalid country code", func(t *testing.T) {
		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "NOT_A_COUNTRY_CODE",
		}})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "invalid country code"), err)
		require.Nil(t, resp)

		resp, err = service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "AA",
		}})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "invalid country code"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on repository error", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		famRepo.EXPECT().
			CreateFamily(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(errToReturn)

		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}})
		require.Equal(t, errToReturn, err)
		require.Nil(t, resp)
	})

	t.Run("should return response with another id", func(t *testing.T) {
		famRepo.EXPECT().
			CreateFamily(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(nil)

		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}})

		require.Nil(t, err)
		require.NotEqual(t, "id", resp.Result)
		require.Equal(t, "request_id", resp.Id)
	})
}

func TestFamilyService_DeleteFamily(t *testing.T) {
	ctl := gomock.NewController(t)
	famRepo := repo.NewMockFamilyRepo(ctl)
	memRepo := repo.NewMockMemberRepo(ctl)
	service, err := NewFamilyService(famRepo, memRepo)
	require.Nil(t, err)
	ctx := requestid.WithRequestId(context.Background(), "request_id")

	t.Run("should return error name should have a character different of a space", func(t *testing.T) {
		resp, err := service.DeleteFamily(ctx, &DeleteFamilyRequest{Id: ""})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on repository error", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		famRepo.EXPECT().
			DeleteFamilyById(gomock.Eq(ctx), gomock.Eq("id")).
			Times(1).
			Return(errToReturn)

		resp, err := service.DeleteFamily(ctx, &DeleteFamilyRequest{Id: "id"})
		require.Equal(t, errToReturn, err)
		require.Nil(t, resp)
	})

	t.Run("should return response with another id", func(t *testing.T) {
		famRepo.EXPECT().
			DeleteFamilyById(gomock.Eq(ctx), gomock.Eq("id")).
			Times(1).
			Return(nil)

		resp, err := service.DeleteFamily(ctx, &DeleteFamilyRequest{Id: "id"})

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, "request_id", resp.Id)
	})
}
