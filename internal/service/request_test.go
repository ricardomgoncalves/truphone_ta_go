package service

import (
	"testing"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestCreateFamilyRequest_GetFamily(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *CreateFamilyRequest
		assert.Equal(t, family.Family{}, req.GetFamily())
	})
	t.Run("should return family", func(t *testing.T) {
		fam := family.Family{Id: "id"}
		req := &CreateFamilyRequest{Family: fam}
		assert.Equal(t, fam, req.GetFamily())
	})
}

func TestDeleteFamilyRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *DeleteFamilyRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return family", func(t *testing.T) {
		req := &DeleteFamilyRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}
