package family

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMemberWithId(t *testing.T) {
	t.Run("should return member with id", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038"
		member := NewMemberWithId(id)
		assert.Equal(t, id, member.Id)
	})
}

func TestMember_UnmarshalJSON(t *testing.T) {
	t.Run("should return valid member", func(t *testing.T) {
		date, _ := time.Parse(time.RFC3339, "2012-02-03T00:04:05Z")
		id := "146adc77-912c-4bd2-a70e-4b136e6b0791"

		sourceMembers := []struct {
			input  []byte
			output Member
		}{
			{
				input: []byte("{\"id\":\"146adc77-912c-4bd2-a70e-4b136e6b0791\",\"family_id\":\"146adc77-912c-4bd2-a70e-4b136e6b0791\",\"first_name\":\"ri\",\"middle_name\":\"mi\",\"last_name\":\"go\",\"father_id\":\"146adc77-912c-4bd2-a70e-4b136e6b0791\",\"mother_id\":\"146adc77-912c-4bd2-a70e-4b136e6b0791\",\"spouse_id\":\"146adc77-912c-4bd2-a70e-4b136e6b0791\",\"birthday\":\"2012-02-03T00:04:05Z\"}"),
				output: Member{
					Id:         id,
					FamilyId:   id,
					FirstName:  "ri",
					MiddleName: "mi",
					LastName:   "go",
					FatherId:   &id,
					MotherId:   &id,
					SpouseId:   &id,
					Birthday:   date,
				},
			},
		}

		for i := 0; i < len(sourceMembers); i++ {
			outMember := Member{}
			err := json.Unmarshal(sourceMembers[i].input, &outMember)
			require.Nil(t, err, "should not get a error on unmarshal", i)

			assert.Equal(t, sourceMembers[i].output.Id, outMember.Id, "Id must be equal")
			assert.Equal(t, sourceMembers[i].output.FamilyId, outMember.FamilyId, "FamilyId must be equal")
			assert.Equal(t, sourceMembers[i].output.FirstName, outMember.FirstName, "FirstName must be equal")
			assert.Equal(t, sourceMembers[i].output.MiddleName, outMember.MiddleName, "MiddleName must be equal")
			assert.Equal(t, sourceMembers[i].output.LastName, outMember.LastName, "LastName must be equal")
			assert.Equal(t, sourceMembers[i].output.FatherId, outMember.FatherId, "FatherId must be equal")
			assert.Equal(t, sourceMembers[i].output.MotherId, outMember.MotherId, "MotherId must be equal")
			assert.Equal(t, sourceMembers[i].output.SpouseId, outMember.SpouseId, "SpouseId must be equal")
			assert.Equal(t, sourceMembers[i].output.Birthday, outMember.Birthday, "Birthday must be equal")
		}
	})
	t.Run("should return error due to invalid date", func(t *testing.T) {
		sourceMembers := [][]byte{
			[]byte("{\"id\":1,\"family_id\":1,\"first_name\":\"ri\",\"middle_name\":\"mi\",\"last_name\":\"go\",\"father_id\":1,\"mother_id\":1,\"spouse_id\":1,\"last_name\":\"2012-02-03\"}"),
		}

		for i := 0; i < len(sourceMembers); i++ {
			outMember := Member{}
			err := json.Unmarshal(sourceMembers[i], outMember)
			require.NotNil(t, err, "should get a error on unmarshal", i)
		}
	})
	t.Run("should return error on json", func(t *testing.T) {
		sourceMembers := [][]byte{
			[]byte(""),
			[]byte("{"),
		}

		for i := 0; i < len(sourceMembers); i++ {
			outMember := Member{}
			err := json.Unmarshal(sourceMembers[i], &outMember)
			require.NotNil(t, err, "should get a error on unmarshal", i)
		}
	})
}

func TestMemberRaw_Parse(t *testing.T) {
	t.Run("should return a valid member", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038"
		rawMember := memberRaw{
			Id:         id,
			FamilyId:   id,
			FirstName:  "Ricardo",
			MiddleName: "Miguel",
			LastName:   "Goncalves",
			FatherId:   &id,
			MotherId:   &id,
			SpouseId:   &id,
			Birthday:   "2012-02-03T00:04:05Z",
		}

		member := rawMember.parse()
		date, _ := time.Parse(time.RFC3339, "2012-02-03T00:04:05Z")

		assert.Equal(t, id, member.Id, "Id must be equal")
		assert.Equal(t, id, member.FamilyId, "FamilyId must be equal")
		assert.Equal(t, "Ricardo", member.FirstName, "FirstName must be equal")
		assert.Equal(t, "Miguel", member.MiddleName, "MiddleName must be equal")
		assert.Equal(t, "Goncalves", member.LastName, "LastName must be equal")
		assert.Equal(t, id, *member.FatherId, "FatherId must be equal")
		assert.Equal(t, id, *member.MotherId, "MotherId must be equal")
		assert.Equal(t, id, *member.SpouseId, "SpouseId must be equal")
		assert.Equal(t, date, member.Birthday, "Birthday must be equal")
	})
}

func TestMember_ToRaw(t *testing.T) {
	t.Run("should return correctly the raw member", func(t *testing.T) {
		date, err := time.Parse(time.RFC3339, "2012-02-03T00:04:05Z")
		require.Nil(t, err, "should be able to parse date")

		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038"
		member := Member{
			Id:         id,
			FamilyId:   id,
			FirstName:  "Ricardo",
			MiddleName: "Miguel",
			LastName:   "Goncalves",
			FatherId:   &id,
			MotherId:   &id,
			SpouseId:   &id,
			Birthday:   date,
		}

		rawMember := member.toRaw()

		assert.Equal(t, id, rawMember.Id, "Id must be equal")
		assert.Equal(t, id, rawMember.FamilyId, "FamilyId must be equal")
		assert.Equal(t, "Ricardo", rawMember.FirstName, "FirstName must be equal")
		assert.Equal(t, "Miguel", rawMember.MiddleName, "MiddleName must be equal")
		assert.Equal(t, "Goncalves", rawMember.LastName, "LastName must be equal")
		assert.Equal(t, id, *rawMember.FatherId, "FatherId must be equal")
		assert.Equal(t, id, *rawMember.MotherId, "MotherId must be equal")
		assert.Equal(t, id, *rawMember.SpouseId, "SpouseId must be equal")
		assert.Equal(t, "2012-02-03T00:04:05Z", rawMember.Birthday, "Birthday must be equal")
	})
}

func TestMember_Patch(t *testing.T) {
	t.Run("should not do anything on nil member", func(t *testing.T) {
		var fam *Member
		fam.Patch(Member{})
	})
	t.Run("should only update family_id", func(t *testing.T) {
		fam := &Member{
			Id:       "id",
			FamilyId: "family",
		}
		fam.Patch(Member{FamilyId: "updated"})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", fam.FamilyId)
	})
	t.Run("should only update first_name", func(t *testing.T) {
		fam := &Member{
			Id:        "id",
			FirstName: "Member",
		}
		fam.Patch(Member{FirstName: "updated"})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", fam.FirstName)
	})
	t.Run("should only update middle_name", func(t *testing.T) {
		fam := &Member{
			Id:         "id",
			MiddleName: "Member",
		}
		fam.Patch(Member{MiddleName: "updated"})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", fam.MiddleName)
	})
	t.Run("should only update last_name", func(t *testing.T) {
		fam := &Member{
			Id:       "id",
			LastName: "Member",
		}
		fam.Patch(Member{LastName: "updated"})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", fam.LastName)
	})
	t.Run("should only update father_id", func(t *testing.T) {
		fatherId1 := "Father"
		fatherId2 := "updated"
		fam := &Member{
			Id:       "id",
			FatherId: &fatherId1,
		}
		fam.Patch(Member{FatherId: &fatherId2})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", *fam.FatherId)
	})
	t.Run("should only update mother_id", func(t *testing.T) {
		motherId1 := "Mother"
		motherId2 := "updated"
		fam := &Member{
			Id:       "id",
			MotherId: &motherId1,
		}
		fam.Patch(Member{MotherId: &motherId2})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", *fam.MotherId)
	})
	t.Run("should only update spouse_id", func(t *testing.T) {
		spouseId1 := "Spouse"
		spouseId2 := "updated"
		fam := &Member{
			Id:       "id",
			SpouseId: &spouseId1,
		}
		fam.Patch(Member{SpouseId: &spouseId2})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", *fam.SpouseId)
	})
	t.Run("should only updated birthday", func(t *testing.T) {
		fam := &Member{
			Id:       "id",
			Birthday: time.Now(),
		}
		tm := time.Now()
		fam.Patch(Member{Birthday: tm})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, tm, fam.Birthday)
	})
}

func TestMember_HasCommonName(t *testing.T) {
	t.Run("should return true on first_name", func(t *testing.T) {
		member := Member{
			FirstName: "First",
		}
		assert.True(t, member.HasCommonName(Member{FirstName: "First"}))
	})
	t.Run("should return true on middle_name", func(t *testing.T) {
		member := Member{
			FirstName:  "First",
			MiddleName: "Middle",
		}
		assert.True(t, member.HasCommonName(Member{MiddleName: "Middle"}))
	})
	t.Run("should return true on first_name", func(t *testing.T) {
		member := Member{
			FirstName:  "First",
			MiddleName: "Middle",
			LastName:   "Last",
		}
		assert.True(t, member.HasCommonName(Member{LastName: "Last"}))
	})
	t.Run("should return false", func(t *testing.T) {
		member := Member{
			FirstName:  "First",
			MiddleName: "Middle",
			LastName:   "Last",
		}
		assert.False(t, member.HasCommonName(Member{}))
	})
}

func TestMember_HasSimilarBirthday(t *testing.T) {
	t.Run("should return true on equal birthdays", func(t *testing.T) {
		birthday := time.Now()
		member := Member{
			Birthday: birthday,
		}
		assert.True(t, member.HasSimilarBirthday(Member{Birthday: birthday}))
	})
	t.Run("should return false", func(t *testing.T) {
		member := Member{
			Birthday: time.Now(),
		}
		assert.False(t, member.HasSimilarBirthday(Member{
			Birthday: time.Now().Add(time.Second * 1),
		}))
	})
}

func TestMember_IsMissingMother(t *testing.T) {
	t.Run("should return true on nil mother_id", func(t *testing.T) {
		assert.True(t, Member{}.IsMissingMother())
	})
	t.Run("should return false", func(t *testing.T) {
		motherId := "id"
		assert.False(t, Member{
			MotherId: &motherId,
		}.IsMissingMother())
	})
}

func TestMember_IsMissingFather(t *testing.T) {
	t.Run("should return true on nil father_id", func(t *testing.T) {
		assert.True(t, Member{}.IsMissingFather())
	})
	t.Run("should return false", func(t *testing.T) {
		fatherId := "id"
		assert.False(t, Member{
			FatherId: &fatherId,
		}.IsMissingFather())
	})
}

func TestMember_HasSameMother(t *testing.T) {
	t.Run("should return true on both nil", func(t *testing.T) {
		member := Member{}
		assert.True(t, member.HasSameMother(Member{}))
	})
	t.Run("should return false on one nil", func(t *testing.T) {
		motherId := "id"
		member := Member{
			MotherId: &motherId,
		}
		assert.False(t, member.HasSameMother(Member{}))
	})
	t.Run("should return true", func(t *testing.T) {
		motherId1 := "id"
		motherId2 := "id"
		member := Member{
			MotherId: &motherId1,
		}
		assert.True(t, member.HasSameMother(Member{MotherId: &motherId2}))
	})
	t.Run("should return false", func(t *testing.T) {
		motherId1 := "id"
		motherId2 := "id2"
		member := Member{
			MotherId: &motherId1,
		}
		assert.False(t, member.HasSameMother(Member{MotherId: &motherId2}))
	})
}

func TestMember_HasSameFather(t *testing.T) {
	t.Run("should return true on both nil", func(t *testing.T) {
		member := Member{}
		assert.True(t, member.HasSameFather(Member{}))
	})
	t.Run("should return false on one nil", func(t *testing.T) {
		fatherId := "id"
		member := Member{
			FatherId: &fatherId,
		}
		assert.False(t, member.HasSameFather(Member{}))
	})
	t.Run("should return true", func(t *testing.T) {
		fatherId1 := "id"
		fatherId2 := "id"
		member := Member{
			FatherId: &fatherId1,
		}
		assert.True(t, member.HasSameFather(Member{FatherId: &fatherId2}))
	})
	t.Run("should return false", func(t *testing.T) {
		fatherId1 := "id"
		fatherId2 := "id2"
		member := Member{
			FatherId: &fatherId1,
		}
		assert.False(t, member.HasSameFather(Member{FatherId: &fatherId2}))
	})
}