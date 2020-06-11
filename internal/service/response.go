package service

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
