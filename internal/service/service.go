package service

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/age"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/countrycode"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/duplicates"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/limiter"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/requestid"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/verify"
)

type Service interface {
	CreateFamily(ctx context.Context, req *CreateFamilyRequest) (*CreateFamilyResponse, error)
	GetFamily(ctx context.Context, req *GetFamilyRequest) (*GetFamilyResponse, error)
	ListFamilies(ctx context.Context, req *ListFamiliesRequest) (*ListFamiliesResponse, error)
	UpdateFamily(ctx context.Context, req *UpdateFamilyRequest) (*UpdateFamilyResponse, error)
	DeleteFamily(ctx context.Context, req *DeleteFamilyRequest) (*DeleteFamilyResponse, error)

	CreateMember(ctx context.Context, req *CreateMemberRequest) (*CreateMemberResponse, error)
	GetMember(ctx context.Context, req *GetMemberRequest) (*GetMemberResponse, error)
	ListMembers(ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error)
	UpdateMember(ctx context.Context, req *UpdateMemberRequest) (*UpdateMemberResponse, error)
	DeleteMember(ctx context.Context, req *DeleteMemberRequest) (*DeleteMemberResponse, error)

	ListAccumulatedFamilies(ctx context.Context, req *ListAccumulatedFamiliesRequest) (*ListAccumulatedFamiliesResponse, error)
	ListFastestGrowingFamilies(ctx context.Context, req *ListFastestGrowingFamiliesRequest) (*ListFastestGrowingFamiliesResponse, error)
	ListPossibleDuplicatesMembers(ctx context.Context, req *ListPossibleDuplicatesMembersRequest) (*ListPossibleDuplicatesMembersResponse, error)
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
		if err := verify.StringLength(newFam.Name, 3, 30); err != nil {
			return nil, errors.Wrap(family.ErrorFamilyBadRequest, err)
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

	if members, err := service.memberRepo.ListMembers(ctx, repo.WithFamilyId(id)); err == nil && len(members) > 0 {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "family is not empty")
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

func (service FamilyService) ListMembers(ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	options := make([]repo.FilterOption, 0)
	if offset := req.GetOffset(); offset != nil {
		options = append(options, repo.WithOffset(*offset))
	}

	if limit := req.GetLimit(); limit != nil {
		options = append(options, repo.WithLimit(*limit))
	}

	if familyId := req.GetFamilyId(); familyId != nil {
		options = append(options, repo.WithFamilyId(*familyId))
	}

	if parentId := req.GetParentId(); parentId != nil {
		options = append(options, repo.WithParentId(*parentId))
	}

	members, err := service.memberRepo.ListMembers(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &ListMembersResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  members,
	}, nil
}

func (service FamilyService) UpdateMember(ctx context.Context, req *UpdateMemberRequest) (*UpdateMemberResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	id := req.GetId()
	if id == "" {
		return nil, errors.Annotate(family.ErrorFamilyBadRequest, "family id should be provided")
	}

	newMember := req.GetMember()

	if newMember.FatherId != nil && *newMember.FatherId == id {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "father_id should not be the same as the id")
	}

	if newMember.MotherId != nil && *newMember.MotherId == id {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "mother_id should not be the same as the id")
	}

	if newMember.SpouseId != nil && *newMember.SpouseId == id {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "spouse_id should not be the same as the id")
	}

	if len(newMember.FirstName) > 0 {
		if err := verify.StringLength(newMember.FirstName, 3, 30); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "first_name"))
		}
	}

	if len(newMember.MiddleName) > 0 {
		if err := verify.StringLength(newMember.MiddleName, 3, 30); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "middle_name"))
		}
	}

	if len(newMember.LastName) > 0 {
		if err := verify.StringLength(newMember.LastName, 3, 30); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "last_name"))
		}
	}

	if len(newMember.FamilyId) > 0 {
		if _, err := service.familyRepo.GetFamilyById(ctx, newMember.FamilyId); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "family_id"))
		}
	}

	if newMember.FatherId != nil && *newMember.FatherId != "" {
		if _, err := service.memberRepo.GetMemberById(ctx, *newMember.FatherId); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "father_id"))
		}
	}

	if newMember.MotherId != nil && *newMember.MotherId != "" {
		if _, err := service.memberRepo.GetMemberById(ctx, *newMember.MotherId); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "mother_id"))
		}
	}

	if newMember.SpouseId != nil && *newMember.SpouseId != "" {
		if _, err := service.memberRepo.GetMemberById(ctx, *newMember.SpouseId); err != nil {
			return nil, errors.Wrap(family.ErrorMemberBadRequest, errors.Annotate(err, "spouse_id"))
		}
	}

	member, err := service.memberRepo.GetMemberById(ctx, id)
	if err != nil {
		return nil, err
	}

	member.Patch(newMember)
	if err := service.memberRepo.UpdateMemberById(ctx, id, *member); err != nil {
		return nil, err
	}

	return &UpdateMemberResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  *member,
	}, nil
}

func (service FamilyService) DeleteMember(ctx context.Context, req *DeleteMemberRequest) (*DeleteMemberResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	id := req.GetId()
	if id == "" {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "member id should be provided")
	}

	if members, err := service.memberRepo.ListMembers(ctx, repo.WithParentId(id)); err == nil && len(members) > 0 {
		return nil, errors.Annotate(family.ErrorMemberBadRequest, "this member currently has children")
	}

	if err := service.memberRepo.DeleteMemberById(ctx, id); err != nil {
		return nil, err
	}

	return &DeleteMemberResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (service FamilyService) ListAccumulatedFamilies(ctx context.Context, req *ListAccumulatedFamiliesRequest) (*ListAccumulatedFamiliesResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	offset := uint32(0)
	limit := uint32(0)

	if val := req.GetOffset(); val != nil {
		offset = *val
	}

	if val := req.GetLimit(); val != nil {
		limit = *val
	}

	families, err := service.familyRepo.ListFamilies(ctx)
	if err != nil {
		return nil, err
	}

	members := make(map[string][]family.Member, 0)
	for _, fam := range families {
		familyMembers, err := service.memberRepo.ListMembers(ctx, repo.WithFamilyId(fam.Id))
		if err != nil {
			log.Printf("Family %s got error: %v", fam.Id, err.Error())
			continue
		}

		members[fam.Id] = familyMembers
	}

	agedFamilies := age.FindAccumulatedAge(families, members)
	agedFamilies = limiter.LimitAgedFamilies(agedFamilies, offset, limit)

	return &ListAccumulatedFamiliesResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  agedFamilies,
	}, nil
}

func (service FamilyService) ListFastestGrowingFamilies(ctx context.Context, req *ListFastestGrowingFamiliesRequest) (*ListFastestGrowingFamiliesResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	offset := uint32(0)
	limit := uint32(0)

	if val := req.GetOffset(); val != nil {
		offset = *val
	}

	if val := req.GetLimit(); val != nil {
		limit = *val
	}

	families, err := service.familyRepo.ListFamilies(ctx)
	if err != nil {
		return nil, err
	}

	members := make(map[string][]family.Member, 0)
	for _, fam := range families {
		familyMembers, err := service.memberRepo.ListMembers(ctx, repo.WithFamilyId(fam.Id))
		if err != nil {
			log.Printf("Family %s got error: %v", fam.Id, err.Error())
			continue
		}

		members[fam.Id] = familyMembers
	}

	agedFamilies := age.FindFastestGrowingAge(families, members)
	agedFamilies = limiter.LimitAgedFamilies(agedFamilies, offset, limit)

	return &ListFastestGrowingFamiliesResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  agedFamilies,
	}, nil
}

func (service FamilyService) ListPossibleDuplicatesMembers(ctx context.Context, req *ListPossibleDuplicatesMembersRequest) (*ListPossibleDuplicatesMembersResponse, error) {
	requestId, _ := requestid.GetRequestId(ctx)

	offset := uint32(0)
	limit := uint32(0)

	if val := req.GetOffset(); val != nil {
		offset = *val
	}

	if val := req.GetLimit(); val != nil {
		limit = *val
	}

	families, err := service.familyRepo.ListFamilies(ctx)
	if err != nil {
		return nil, err
	}

	members := make([]family.Member, 0)
	for _, fam := range families {
		familyMembers, err := service.memberRepo.ListMembers(ctx, repo.WithFamilyId(fam.Id))
		if err != nil {
			continue
		}

		members = append(members, duplicates.FindPossibleDuplicates(familyMembers)...)
	}

	members = limiter.LimitMembers(members, offset, limit)
	return &ListPossibleDuplicatesMembersResponse{
		Id:      requestId,
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Result:  members,
	}, nil
}
