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

func TestListFamiliesRequest_GetOffset(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListFamiliesRequest
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return nil limit", func(t *testing.T) {
		req := &ListFamiliesRequest{Limit: nil}
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return limit", func(t *testing.T) {
		limit := uint32(2)
		req := &ListFamiliesRequest{Limit: &limit}
		assert.Equal(t, &limit, req.GetLimit())
	})
}

func TestListFamiliesRequest_GetLimit(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListFamiliesRequest
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return nil offset", func(t *testing.T) {
		req := &ListFamiliesRequest{Offset: nil}
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return offset", func(t *testing.T) {
		offset := uint32(2)
		req := &ListFamiliesRequest{Offset: &offset}
		assert.Equal(t, &offset, req.GetOffset())
	})
}

func TestListFamiliesRequest_GetCountryCode(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListFamiliesRequest
		assert.Equal(t, (*string)(nil), req.GetCountryCode())
	})
	t.Run("should return nil country", func(t *testing.T) {
		req := &ListFamiliesRequest{CountryCode: nil}
		assert.Equal(t, (*string)(nil), req.GetCountryCode())
	})
	t.Run("should return country", func(t *testing.T) {
		countryCode := "PT"
		req := &ListFamiliesRequest{CountryCode: &countryCode}
		assert.Equal(t, &countryCode, req.GetCountryCode())
	})
}

func TestGetFamilyRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *GetFamilyRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return id", func(t *testing.T) {
		req := &GetFamilyRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}

func TestUpdateFamilyRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *UpdateFamilyRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return id", func(t *testing.T) {
		req := &UpdateFamilyRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}

func TestUpdateFamilyRequest_GetFamily(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *UpdateFamilyRequest
		assert.Equal(t, family.Family{}, req.GetFamily())
	})
	t.Run("should return family", func(t *testing.T) {
		fam := family.Family{Id: "id"}
		req := &UpdateFamilyRequest{Family: fam}
		assert.Equal(t, fam, req.GetFamily())
	})
}

func TestDeleteFamilyRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *DeleteFamilyRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return id", func(t *testing.T) {
		req := &DeleteFamilyRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}

func TestCreateMemberRequest_GetMember(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *CreateMemberRequest
		assert.Equal(t, family.Member{}, req.GetMember())
	})
	t.Run("should return member", func(t *testing.T) {
		req := &CreateMemberRequest{Member: family.Member{Id: "id"}}
		assert.Equal(t, family.Member{Id: "id"}, req.GetMember())
	})
}

func TestGetMemberRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *GetMemberRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return id", func(t *testing.T) {
		req := &GetMemberRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}

func TestUpdateMemberRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *UpdateMemberRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return id", func(t *testing.T) {
		req := &UpdateMemberRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}

func TestListMembersRequest_GetOffset(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListMembersRequest
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return nil limit", func(t *testing.T) {
		req := &ListMembersRequest{Limit: nil}
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return limit", func(t *testing.T) {
		limit := uint32(2)
		req := &ListMembersRequest{Limit: &limit}
		assert.Equal(t, &limit, req.GetLimit())
	})
}

func TestListMembersRequest_GetLimit(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListMembersRequest
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return nil offset", func(t *testing.T) {
		req := &ListMembersRequest{Offset: nil}
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return offset", func(t *testing.T) {
		offset := uint32(2)
		req := &ListMembersRequest{Offset: &offset}
		assert.Equal(t, &offset, req.GetOffset())
	})
}

func TestListMembersRequest_GetFamilyId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListMembersRequest
		assert.Equal(t, (*string)(nil), req.GetFamilyId())
	})
	t.Run("should return nil country", func(t *testing.T) {
		req := &ListMembersRequest{FamilyId: nil}
		assert.Equal(t, (*string)(nil), req.GetFamilyId())
	})
	t.Run("should return country", func(t *testing.T) {
		countryCode := "PT"
		req := &ListMembersRequest{FamilyId: &countryCode}
		assert.Equal(t, &countryCode, req.GetFamilyId())
	})
}

func TestListMembersRequest_GetParentId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListMembersRequest
		assert.Equal(t, (*string)(nil), req.GetParentId())
	})
	t.Run("should return nil country", func(t *testing.T) {
		req := &ListMembersRequest{ParentId: nil}
		assert.Equal(t, (*string)(nil), req.GetParentId())
	})
	t.Run("should return country", func(t *testing.T) {
		countryCode := "PT"
		req := &ListMembersRequest{ParentId: &countryCode}
		assert.Equal(t, &countryCode, req.GetParentId())
	})
}

func TestUpdateMemberRequest_GetMember(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *UpdateMemberRequest
		assert.Equal(t, family.Member{}, req.GetMember())
	})
	t.Run("should return member", func(t *testing.T) {
		member := family.Member{Id: "id"}
		req := &UpdateMemberRequest{Member: member}
		assert.Equal(t, member, req.GetMember())
	})
}

func TestDeleteMemberRequest_GetId(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *DeleteMemberRequest
		assert.Equal(t, "", req.GetId())
	})
	t.Run("should return id", func(t *testing.T) {
		req := &DeleteMemberRequest{Id: "id"}
		assert.Equal(t, "id", req.GetId())
	})
}

func TestListAccumulatedFamiliesRequest_GetLimit(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListAccumulatedFamiliesRequest
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return nil limit", func(t *testing.T) {
		req := &ListAccumulatedFamiliesRequest{Limit: nil}
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return limit", func(t *testing.T) {
		limit := uint32(2)
		req := &ListAccumulatedFamiliesRequest{Limit: &limit}
		assert.Equal(t, &limit, req.GetLimit())
	})
}

func TestListAccumulatedFamiliesRequest_GetOffset(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListAccumulatedFamiliesRequest
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return nil offset", func(t *testing.T) {
		req := &ListAccumulatedFamiliesRequest{Offset: nil}
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return offset", func(t *testing.T) {
		offset := uint32(2)
		req := &ListAccumulatedFamiliesRequest{Offset: &offset}
		assert.Equal(t, &offset, req.GetOffset())
	})
}

func TestListFastestGrowingFamiliesRequest_GetLimit(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListFastestGrowingFamiliesRequest
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return nil limit", func(t *testing.T) {
		req := &ListFastestGrowingFamiliesRequest{Limit: nil}
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return limit", func(t *testing.T) {
		limit := uint32(2)
		req := &ListFastestGrowingFamiliesRequest{Limit: &limit}
		assert.Equal(t, &limit, req.GetLimit())
	})
}

func TestListFastestGrowingFamiliesRequest_GetOffset(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListFastestGrowingFamiliesRequest
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return nil offset", func(t *testing.T) {
		req := &ListFastestGrowingFamiliesRequest{Offset: nil}
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return offset", func(t *testing.T) {
		offset := uint32(2)
		req := &ListFastestGrowingFamiliesRequest{Offset: &offset}
		assert.Equal(t, &offset, req.GetOffset())
	})
}

func TestListPossibleDuplicatesMembersRequest_GetLimit(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListPossibleDuplicatesMembersRequest
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return nil limit", func(t *testing.T) {
		req := &ListPossibleDuplicatesMembersRequest{Limit: nil}
		assert.Equal(t, (*uint32)(nil), req.GetLimit())
	})
	t.Run("should return limit", func(t *testing.T) {
		limit := uint32(2)
		req := &ListPossibleDuplicatesMembersRequest{Limit: &limit}
		assert.Equal(t, &limit, req.GetLimit())
	})
}

func TestListPossibleDuplicatesMembersRequest_GetOffset(t *testing.T) {
	t.Run("should return nil on nil request", func(t *testing.T) {
		var req *ListPossibleDuplicatesMembersRequest
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return nil offset", func(t *testing.T) {
		req := &ListPossibleDuplicatesMembersRequest{Offset: nil}
		assert.Equal(t, (*uint32)(nil), req.GetOffset())
	})
	t.Run("should return offset", func(t *testing.T) {
		offset := uint32(2)
		req := &ListPossibleDuplicatesMembersRequest{Offset: &offset}
		assert.Equal(t, &offset, req.GetOffset())
	})
}