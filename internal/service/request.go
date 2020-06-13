package service

import "github.com/ricardomgoncalves/truphone_ta_go/pkg/family"

// CreateFamilyRequest
//
// Request by CreateFamily Service.
//
// swagger:model
type CreateFamilyRequest struct {
	// the family for the request
	//
	// required: true
	Family family.Family `json:"family"`
}

func (req *CreateFamilyRequest) GetFamily() family.Family {
	if req == nil {
		return family.Family{}
	}

	return req.Family
}

// ListFamiliesRequest
//
// Request by ListFamilies Service.
//
// swagger:model
type ListFamiliesRequest struct {
	// the offset for the request
	//
	// required: false
	// example: 1
	Offset *uint32 `json:"offset"`

	// the limit for the request
	//
	// required: false
	// example: 1
	Limit *uint32 `json:"limit"`

	// the country code for the request
	//
	// required: false
	// example: PT
	CountryCode *string `json:"country_code"`
}

func (req *ListFamiliesRequest) GetOffset() *uint32 {
	if req == nil {
		return nil
	}

	return req.Offset
}

func (req *ListFamiliesRequest) GetLimit() *uint32 {
	if req == nil {
		return nil
	}

	return req.Limit
}

func (req *ListFamiliesRequest) GetCountryCode() *string {
	if req == nil {
		return nil
	}

	return req.CountryCode
}

// GetFamilyRequest
//
// Request by GetFamily Service.
//
// swagger:model
type GetFamilyRequest struct {
	// the family for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`
}

func (req *GetFamilyRequest) GetId() string {
	if req == nil {
		return ""
	}

	return req.Id
}

// UpdateFamilyRequest
//
// Request by UpdateFamily Service.
//
// swagger:model
type UpdateFamilyRequest struct {
	// the family for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the family for the request
	//
	// required: true
	Family family.Family `json:"family"`
}

func (req *UpdateFamilyRequest) GetId() string {
	if req == nil {
		return ""
	}

	return req.Id
}

func (req *UpdateFamilyRequest) GetFamily() family.Family {
	if req == nil {
		return family.Family{}
	}

	return req.Family
}

// DeleteFamilyRequest
//
// Request by DeleteFamily Service.
//
// swagger:model
type DeleteFamilyRequest struct {
	// the family for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`
}

func (req *DeleteFamilyRequest) GetId() string {
	if req == nil {
		return ""
	}

	return req.Id
}

// CreateMemberRequest
//
// Request by CreateMember Service.
//
// swagger:model
type CreateMemberRequest struct {
	// the family for the request
	//
	// required: true
	Member family.Member `json:"member"`
}

func (req *CreateMemberRequest) GetMember() family.Member {
	if req == nil {
		return family.Member{}
	}

	return req.Member
}

// GetMemberRequest
//
// Request by GetMember Service.
//
// swagger:model
type GetMemberRequest struct {
	// the family for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`
}

func (req *GetMemberRequest) GetId() string {
	if req == nil {
		return ""
	}

	return req.Id
}

// ListMembersRequest
//
// Request by ListMembers Service.
//
// swagger:model
type ListMembersRequest struct {
	// the offset for the request
	//
	// required: false
	// example: 1
	Offset *uint32 `json:"offset"`

	// the limit for the request
	//
	// required: false
	// example: 1
	Limit *uint32 `json:"limit"`

	// the family id for the request
	//
	// required: false
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	FamilyId *string `json:"family_id"`

	// the parent id for the request
	//
	// required: false
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	ParentId *string `json:"parent_id"`
}

func (req *ListMembersRequest) GetOffset() *uint32 {
	if req == nil {
		return nil
	}

	return req.Offset
}

func (req *ListMembersRequest) GetLimit() *uint32 {
	if req == nil {
		return nil
	}

	return req.Limit
}

func (req *ListMembersRequest) GetFamilyId() *string {
	if req == nil {
		return nil
	}

	return req.FamilyId
}

func (req *ListMembersRequest) GetParentId() *string {
	if req == nil {
		return nil
	}

	return req.ParentId
}

// UpdateMemberRequest
//
// Request by UpdateMember Service.
//
// swagger:model
type UpdateMemberRequest struct {
	// the family for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the family for the request
	//
	// required: true
	Member family.Member `json:"member"`
}

func (req *UpdateMemberRequest) GetId() string {
	if req == nil {
		return ""
	}

	return req.Id
}

func (req *UpdateMemberRequest) GetMember() family.Member {
	if req == nil {
		return family.Member{}
	}

	return req.Member
}

// DeleteMemberRequest
//
// Request by DeleteMember Service.
//
// swagger:model
type DeleteMemberRequest struct {
	// the family for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`
}

func (req *DeleteMemberRequest) GetId() string {
	if req == nil {
		return ""
	}

	return req.Id
}

// ListAccumulatedFamiliesRequest
//
// Request by ListAccumulatedFamilies Service.
//
// swagger:model
type ListAccumulatedFamiliesRequest struct {
	// the offset for the request
	//
	// required: false
	// example: 1
	Offset *uint32 `json:"offset"`

	// the limit for the request
	//
	// required: false
	// example: 1
	Limit *uint32 `json:"limit"`
}

func (req *ListAccumulatedFamiliesRequest) GetOffset() *uint32 {
	if req == nil {
		return nil
	}

	return req.Offset
}

func (req *ListAccumulatedFamiliesRequest) GetLimit() *uint32 {
	if req == nil {
		return nil
	}

	return req.Limit
}

// ListFastestGrowingFamiliesRequest
//
// Request by ListFastestGrowingFamilies Service.
//
// swagger:model
type ListFastestGrowingFamiliesRequest struct {
	// the offset for the request
	//
	// required: false
	// example: 1
	Offset *uint32 `json:"offset"`

	// the limit for the request
	//
	// required: false
	// example: 1
	Limit *uint32 `json:"limit"`
}

func (req *ListFastestGrowingFamiliesRequest) GetOffset() *uint32 {
	if req == nil {
		return nil
	}

	return req.Offset
}

func (req *ListFastestGrowingFamiliesRequest) GetLimit() *uint32 {
	if req == nil {
		return nil
	}

	return req.Limit
}

// ListPossibleDuplicatesMembersRequest
//
// Request by ListPossibleDuplicatesMembers Service.
//
// swagger:model
type ListPossibleDuplicatesMembersRequest struct {
	// the offset for the request
	//
	// required: false
	// example: 1
	Offset *uint32 `json:"offset"`

	// the limit for the request
	//
	// required: false
	// example: 1
	Limit *uint32 `json:"limit"`
}

func (req *ListPossibleDuplicatesMembersRequest) GetOffset() *uint32 {
	if req == nil {
		return nil
	}

	return req.Offset
}

func (req *ListPossibleDuplicatesMembersRequest) GetLimit() *uint32 {
	if req == nil {
		return nil
	}

	return req.Limit
}
