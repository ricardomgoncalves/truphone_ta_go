package service

import "github.com/ricardomgoncalves/truphone_ta_go/pkg/family"

// CreateFamilyResponse
//
// Response by CreateFamily Service.
//
// swagger:model
type CreateFamilyResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the message for the request
	//
	// required: true
	Result string `json:"result,omitempty"`
}

// ListFamiliesResponse
//
// Response by ListFamilies Service.
//
// swagger:model
type ListFamiliesResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the families for the request
	//
	// required: true
	Result []family.Family `json:"result,omitempty"`
}

// GetFamilyResponse
//
// Response by GetFamily Service.
//
// swagger:model
type GetFamilyResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the family for the request
	//
	// required: true
	Result family.Family `json:"result,omitempty"`
}

// UpdateFamilyResponse
//
// Response by UpdateFamily Service.
//
// swagger:model
type UpdateFamilyResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the family for the request
	//
	// required: true
	Result family.Family `json:"result,omitempty"`
}

// DeleteFamilyResponse
//
// Response by DeleteFamily Service.
//
// swagger:model
type DeleteFamilyResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`
}

// CreateMemberResponse
//
// Response by CreateMember Service.
//
// swagger:model
type CreateMemberResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the message for the request
	//
	// required: true
	Result string `json:"result,omitempty"`
}

// GetMemberResponse
//
// Response by GetMember Service.
//
// swagger:model
type GetMemberResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the member for the request
	//
	// required: true
	Result family.Member `json:"result,omitempty"`
}

// ListMembersResponse
//
// Response by ListMembers Service.
//
// swagger:model
type ListMembersResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the members for the request
	//
	// required: true
	Result []family.Member `json:"result,omitempty"`
}

// UpdateMemberResponse
//
// Response by UpdateMember Service.
//
// swagger:model
type UpdateMemberResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`

	// the member for the request
	//
	// required: true
	Result family.Member `json:"result,omitempty"`
}

// DeleteMemberResponse
//
// Response by DeleteMember Service.
//
// swagger:model
type DeleteMemberResponse struct {
	// the id for the request
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the response code for the request
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the message for the request
	//
	// required: true
	// example: success
	Message string `json:"message"`
}