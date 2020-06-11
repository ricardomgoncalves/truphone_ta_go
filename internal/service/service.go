package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/countrycode"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/requestid"
	"net/http"
	"strings"
)

type Service interface {
	CreateFamily(ctx context.Context, req *CreateFamilyRequest) (*CreateFamilyResponse, error)
	/*GetFamilyById (ctx context.Context, req *GetFamilyByIdRequest) (*GetFamilyByIdResponse, error)
	ListFamilies (ctx context.Context, req *ListFamiliesRequest) (*ListFamiliesResponse, error)
	UpdateFamily (ctx context.Context, req *UpdateFamilyRequest) (*UpdateFamilyResponse, error)*/
	DeleteFamily(ctx context.Context, req *DeleteFamilyRequest) (*DeleteFamilyResponse, error)

	/*CreateMember (ctx context.Context, req *CreateMemberRequest) (*CreateMemberResponse, error)
	GetMemberById (ctx context.Context, req *GetMemberByIdRequest) (*GetMemberByIdResponse, error)
	ListMembers (ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error)
	UpdateMember (ctx context.Context, req *UpdateMemberRequest) (*UpdateMemberResponse, error)
	DeleteMember (ctx context.Context, req *DeleteMemberRequest) (*DeleteMemberResponse, error)*/
}

type FamilyService struct {
	familyRepo repo.FamilyRepo
	memberRepo repo.MemberRepo
}

func NewFamilyService(familyRepo repo.FamilyRepo, memberRepo repo.MemberRepo) (FamilyService, error) {
	return FamilyService{
		familyRepo: familyRepo,
		memberRepo: memberRepo,
	}, nil
}

func (service FamilyService) CreateFamily(ctx context.Context, req *CreateFamilyRequest) (*CreateFamilyResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	fam := req.GetFamily()
	fam.Id = uuid.New().String()

	if strings.ReplaceAll(fam.Name, " ", "") == "" {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "name should have a character different of a space")
	}

	if len(fam.Name) < 3 {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at least 3 characters")
	}

	if len(fam.Name) > 30 {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at maximum 30 characters")
	}

	if !countrycode.IsValid(fam.CountryCode) {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "invalid country code")
	}

	if err := service.familyRepo.CreateFamily(ctx, fam); err != nil {
		return nil, err
	}

	return &CreateFamilyResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  fam.Id,
	}, nil
}

func (service FamilyService) DeleteFamily(ctx context.Context, req *DeleteFamilyRequest) (*DeleteFamilyResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	id := req.GetId()

	if id == "" {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided")
	}

	if err := service.familyRepo.DeleteFamilyById(ctx, id); err != nil {
		return nil, err
	}

	return &DeleteFamilyResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}
