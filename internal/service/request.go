package service

import "github.com/ricardomgoncalves/truphone_ta_go/pkg/family"

// CreateFamilyRequest
//
// Response by CreateFamily Service.
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

// DeleteFamilyRequest
//
// Response by DeleteFamily Service.
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
