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
	"time"
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
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.New("name should have a character different of a space")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should be at least 3 characters", func(t *testing.T) {
		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "aa",
			CountryCode: "PT",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.New("name should be at least 3 characters")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should be at maximum 30 characters", func(t *testing.T) {
		resp, err := service.CreateFamily(ctx, &CreateFamilyRequest{Family: family.Family{
			Id:          "id",
			Name:        "this name is too long to be in a family name",
			CountryCode: "PT",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.New("name should be at maximum 30 characters")), err)
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

func TestFamilyService_ListFamilies(t *testing.T) {
	ctl := gomock.NewController(t)
	famRepo := repo.NewMockFamilyRepo(ctl)
	memRepo := repo.NewMockMemberRepo(ctl)
	service, err := NewFamilyService(famRepo, memRepo)
	require.Nil(t, err)
	ctx := requestid.WithRequestId(context.Background(), "request_id")

	t.Run("should return families correctly ", func(t *testing.T) {
		req := ListFamiliesRequest{
			Offset:      nil,
			Limit:       nil,
			CountryCode: nil,
		}

		opts := []repo.FilterOption{}
		famRepo.EXPECT().
			ListFamilies(gomock.Eq(ctx), gomock.Eq(opts)).
			Times(1).
			Return([]family.Family{}, nil)

		resp, err := service.ListFamilies(ctx, &req)
		require.Nil(t, err)
		require.Equal(t, []family.Family{}, resp.Result)
	})

	t.Run("should return random error", func(t *testing.T) {
		req := ListFamiliesRequest{
			Offset:      nil,
			Limit:       nil,
			CountryCode: nil,
		}

		opts := []repo.FilterOption{}
		famRepo.EXPECT().
			ListFamilies(gomock.Eq(ctx), gomock.Eq(opts)).
			Times(1).
			Return(nil, errors.New("test"))

		resp, err := service.ListFamilies(ctx, &req)
		require.Equal(t, errors.New("test"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error", func(t *testing.T) {
		val := uint32(1)
		country := "PT"
		req := ListFamiliesRequest{
			Offset:      &val,
			Limit:       &val,
			CountryCode: &country,
		}

		famRepo.EXPECT().
			ListFamilies(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(nil, errors.New("test"))

		resp, err := service.ListFamilies(ctx, &req)
		require.Equal(t, errors.New("test"), err)
		require.Nil(t, resp)
	})
}

func TestFamilyService_GetFamily(t *testing.T) {
	ctl := gomock.NewController(t)
	famRepo := repo.NewMockFamilyRepo(ctl)
	memRepo := repo.NewMockMemberRepo(ctl)
	service, err := NewFamilyService(famRepo, memRepo)
	require.Nil(t, err)
	ctx := requestid.WithRequestId(context.Background(), "request_id")

	t.Run("should return error family id should be provided", func(t *testing.T) {
		resp, err := service.GetFamily(ctx, &GetFamilyRequest{Id: ""})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on repository error", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("id")).
			Times(1).
			Return(nil, errToReturn)

		resp, err := service.GetFamily(ctx, &GetFamilyRequest{Id: "id"})
		require.Equal(t, errToReturn, err)
		require.Nil(t, resp)
	})

	t.Run("should return family", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(&fam, nil)

		resp, err := service.GetFamily(ctx, &GetFamilyRequest{Id: "id"})

		require.Nil(t, err)
		require.Equal(t, fam, resp.Result)
		require.Equal(t, "request_id", resp.Id)
	})
}

func TestFamilyService_UpdateFamily(t *testing.T) {
	ctl := gomock.NewController(t)
	famRepo := repo.NewMockFamilyRepo(ctl)
	memRepo := repo.NewMockMemberRepo(ctl)
	service, err := NewFamilyService(famRepo, memRepo)
	require.Nil(t, err)
	ctx := requestid.WithRequestId(context.Background(), "request_id")

	t.Run("should return error family id should be provided", func(t *testing.T) {
		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{Id: ""})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should have a character different of a space", func(t *testing.T) {
		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Id:          "id",
				Name:        "        ",
				CountryCode: "PT",
			},
		})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "name should have a character different of a space"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should be at least 3 characters", func(t *testing.T) {
		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Id:          "id",
				Name:        "aa",
				CountryCode: "PT",
			},
		})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at least 3 characters"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error name should be at maximum 30 characters", func(t *testing.T) {
		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Id:          "id",
				Name:        "this name is too long to be in a family name",
				CountryCode: "PT",
			},
		})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at maximum 30 characters"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error invalid country code", func(t *testing.T) {
		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Id:          "id",
				Name:        "Family 1",
				CountryCode: "NOT_A_COUNTRY_CODE",
			},
		})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "invalid country code"), err)
		require.Nil(t, resp)

		resp, err = service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Id:          "id",
				Name:        "Family 1",
				CountryCode: "AA",
			},
		})
		require.Equal(t, errors.Annotate(family.ErrorFamilyBadRequest, "invalid country code"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on repository error", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("id")).
			Times(1).
			Return(nil, errToReturn)

		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Id:          "id",
				Name:        "this name",
				CountryCode: "PT",
			},
		})
		require.Equal(t, errToReturn, err)
		require.Nil(t, resp)
	})

	t.Run("should return error on repository error", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		fam := family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("id")).
			Times(1).
			Return(&fam, nil)
		famRepo.EXPECT().
			UpdateFamilyById(gomock.Eq(ctx), gomock.Eq("id"), gomock.Any()).
			Times(1).
			Return(errToReturn)

		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id:     "id",
			Family: fam,
		})
		require.Equal(t, errToReturn, err)
		require.Nil(t, resp)
	})

	t.Run("should return family", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(&fam, nil)
		famRepo.EXPECT().
			UpdateFamilyById(gomock.Eq(ctx), gomock.Eq("id"), gomock.Any()).
			Times(1).
			Return(nil)

		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id:     "id",
			Family: fam,
		})

		require.Nil(t, err)
		require.Equal(t, fam, resp.Result)
		require.Equal(t, "request_id", resp.Id)
	})

	t.Run("should only update name", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(&fam, nil)
		famRepo.EXPECT().
			UpdateFamilyById(gomock.Eq(ctx), gomock.Eq("id"), gomock.Any()).
			Times(1).
			Return(nil)

		resp, err := service.UpdateFamily(ctx, &UpdateFamilyRequest{
			Id: "id",
			Family: family.Family{
				Name: "updated",
			},
		})

		require.Nil(t, err)
		require.NotEqual(t, family.Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}, resp.Result)
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

	t.Run("should return error family id should be provided", func(t *testing.T) {
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

	t.Run("should return no error", func(t *testing.T) {
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

func TestFamilyService_CreateMember(t *testing.T) {
	ctl := gomock.NewController(t)
	famRepo := repo.NewMockFamilyRepo(ctl)
	memRepo := repo.NewMockMemberRepo(ctl)
	service, err := NewFamilyService(famRepo, memRepo)
	require.Nil(t, err)
	ctx := requestid.WithRequestId(context.Background(), "request_id")

	t.Run("should return error first name should have a character different of a space", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName: "      ",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should have a character different of a space"), "first_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error first name should be at least 3 characters", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName: "aa",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should be at least 3 characters"), "first_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error first name should be at maximum 30 characters", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName: "this name is soooo long that sould not be a name",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should be at maximum 30 characters"), "first_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error middle name should have a character different of a space", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "      ",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should have a character different of a space"), "middle_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error middle name should be at least 3 characters", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aa",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should be at least 3 characters"), "middle_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error middle name should be at maximum 30 characters", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "this name is soooo long that sould not be a name",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should be at maximum 30 characters"), "middle_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error last name should have a character different of a space", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "      ",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should have a character different of a space"), "last_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error last name should be at least 3 characters", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aa",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should be at least 3 characters"), "last_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error last name should be at maximum 30 characters", func(t *testing.T) {
		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "this name is soooo long that sould not be a name",
		}})
		require.Equal(t, errors.Wrap(family.ErrorFamilyBadRequest, errors.Annotate(
			errors.New("name should be at maximum 30 characters"), "last_name")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on get family", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("family_id")).
			Times(1).
			Return(nil, errToReturn)

		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aaa",
			FamilyId:   "family_id",
		}})
		require.Equal(t, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(errToReturn, "family_id")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on get father id", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		fatherId := "father_id"
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("family_id")).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(fatherId)).
			Times(1).
			Return(nil, errToReturn)

		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aaa",
			FamilyId:   "family_id",
			FatherId:   &fatherId,
		}})
		require.Equal(t, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(errToReturn, "father_id")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on get mother id", func(t *testing.T) {
		errToReturn := errors.New("some random error")
		fatherId := "father_id"
		motherId := "mother_id"
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("family_id")).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(fatherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(motherId)).
			Times(1).
			Return(nil, errToReturn)

		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aaa",
			FamilyId:   "family_id",
			FatherId:   &fatherId,
			MotherId:   &motherId,
		}})
		require.Equal(t, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(errToReturn, "mother_id")), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on get birthday", func(t *testing.T) {
		fatherId := "father_id"
		motherId := "mother_id"
		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("family_id")).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(fatherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(motherId)).
			Times(1).
			Return(nil, nil)

		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aaa",
			FamilyId:   "family_id",
			FatherId:   &fatherId,
			MotherId:   &motherId,
		}})
		require.Equal(t, errors.Annotate(family.ErrorMemberBadRequest, "birthday should be set"), err)
		require.Nil(t, resp)
	})

	t.Run("should return error on create member", func(t *testing.T) {
		fatherId := "father_id"
		motherId := "mother_id"
		errToReturn := errors.New("random error")

		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("family_id")).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(fatherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(motherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(motherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			CreateMember(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(errToReturn)

		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aaa",
			FamilyId:   "family_id",
			FatherId:   &fatherId,
			MotherId:   &motherId,
			Birthday:   time.Now(),
		}})
		require.Equal(t, errToReturn, err)
		require.Nil(t, resp)
	})

	t.Run("should return response with another id", func(t *testing.T) {
		fatherId := "father_id"
		motherId := "mother_id"

		famRepo.EXPECT().
			GetFamilyById(gomock.Eq(ctx), gomock.Eq("family_id")).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(fatherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(motherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			GetMemberById(gomock.Eq(ctx), gomock.Eq(motherId)).
			Times(1).
			Return(nil, nil)
		memRepo.EXPECT().
			CreateMember(gomock.Eq(ctx), gomock.Any()).
			Times(1).
			Return(nil)

		resp, err := service.CreateMember(ctx, &CreateMemberRequest{Member: family.Member{
			Id:         "id",
			FirstName:  "aaa",
			MiddleName: "aaa",
			LastName:   "aaa",
			FamilyId:   "family_id",
			FatherId:   &fatherId,
			MotherId:   &motherId,
			Birthday:   time.Now(),
		}})

		require.Nil(t, err)
		require.NotEqual(t, "id", resp.Result)
		require.Equal(t, "request_id", resp.Id)
	})
}
