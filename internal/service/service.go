package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/countrycode"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/requestid"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/verify"
	"net/http"
	"strings"
)

type Service interface {
	CreateFamily(ctx context.Context, req *CreateFamilyRequest) (*CreateFamilyResponse, error)
	GetFamily(ctx context.Context, req *GetFamilyRequest) (*GetFamilyResponse, error)
	ListFamilies(ctx context.Context, req *ListFamiliesRequest) (*ListFamiliesResponse, error)
	UpdateFamily(ctx context.Context, req *UpdateFamilyRequest) (*UpdateFamilyResponse, error)
	DeleteFamily(ctx context.Context, req *DeleteFamilyRequest) (*DeleteFamilyResponse, error)

	CreateMember(ctx context.Context, req *CreateMemberRequest) (*CreateMemberResponse, error)
	GetMember(ctx context.Context, req *GetMemberRequest) (*GetMemberResponse, error)
	/*ListMembers (ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error)
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

	if err := verify.StringLength(fam.Name, 3, 30); err != nil {
		return nil, errors.Wrap(family.ErrorFamilyBadRequest, err)
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

func (service FamilyService) GetFamily(ctx context.Context, req *GetFamilyRequest) (*GetFamilyResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	id := req.GetId()
	if id == "" {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided")
	}

	fam, err := service.familyRepo.GetFamilyById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetFamilyResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  *fam,
	}, nil
}

func (service FamilyService) ListFamilies(ctx context.Context, req *ListFamiliesRequest) (*ListFamiliesResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	options := make([]repo.FilterOption, 0)

	if offset := req.GetOffset(); offset != nil {
		options = append(options, repo.WithOffset(*offset))
	}

	if limit := req.GetLimit(); limit != nil {
		options = append(options, repo.WithLimit(*limit))
	}

	if countryCode := req.GetCountryCode(); countryCode != nil {
		options = append(options, repo.WithCountryCode(*countryCode))
	}

	families, err := service.familyRepo.ListFamilies(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &ListFamiliesResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  families,
	}, nil
}

func (service FamilyService) UpdateFamily(ctx context.Context, req *UpdateFamilyRequest) (*UpdateFamilyResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	id := req.GetId()
	if id == "" {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided")
	}

	newFam := req.GetFamily()
	if len(newFam.Name) > 0 {
		if strings.ReplaceAll(newFam.Name, " ", "") == "" {
			return nil, errors.Annotate(family.ErrorFamilyBadRequest, "name should have a character different of a space")
		}

		if len(newFam.Name) < 3 {
			return nil, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at least 3 characters")
		}

		if len(newFam.Name) > 30 {
			return nil, errors.Annotate(family.ErrorFamilyBadRequest, "name should be at maximum 30 characters")
		}
	}

	if len(newFam.CountryCode) > 0 && !countrycode.IsValid(newFam.CountryCode) {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "invalid country code")
	}

	fam, err := service.familyRepo.GetFamilyById(ctx, id)
	if err != nil {
		return nil, err
	}

	fam.Patch(req.GetFamily())
	if err := service.familyRepo.UpdateFamilyById(ctx, id, *fam); err != nil {
		return nil, err
	}

	return &UpdateFamilyResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  *fam,
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

func (service FamilyService) CreateMember(ctx context.Context, req *CreateMemberRequest) (*CreateMemberResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	member := req.GetMember()
	member.Id = uuid.New().String()

	if err := verify.StringLength(member.FirstName, 3, 30); err != nil {
		return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "first_name"))
	}

	if err := verify.StringLength(member.MiddleName, 3, 30); err != nil {
		return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "middle_name"))
	}

	if err := verify.StringLength(member.LastName, 3, 30); err != nil {
		return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "last_name"))
	}

	if _, err := service.familyRepo.GetFamilyById(ctx, member.FamilyId); err != nil {
		return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "family_id"))
	}

	if member.FatherId != nil {
		if _, err := service.memberRepo.GetMemberById(ctx, *member.FatherId); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "father_id"))
		}
	}

	if member.MotherId != nil {
		if _, err := service.memberRepo.GetMemberById(ctx, *member.MotherId); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "mother_id"))
		}
	}

	if member.Birthday.IsZero() {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "birthday should be set")
	}

	if err := service.memberRepo.CreateMember(ctx, member); err != nil {
		return nil, err
	}

	return &CreateMemberResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  member.Id,
	}, nil
}

func (service FamilyService) GetMember(ctx context.Context, req *GetMemberRequest) (*GetMemberResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	id := req.GetId()
	if id == "" {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "member id should be provided")
	}

	member, err := service.memberRepo.GetMemberById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetMemberResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  *member,
	}, nil
}
