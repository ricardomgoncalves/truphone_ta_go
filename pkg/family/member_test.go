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

		member, err := rawMember.parse()
		date, _ := time.Parse(time.RFC3339, "2012-02-03T00:04:05Z")

		require.Nil(t, err, "error must be nil")
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

	t.Run("should return an error", func(t *testing.T) {
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
			Birthday:   "2012-02-03T0",
		}

		_, err := rawMember.parse()
		require.NotNil(t, err, "error must not be nil")
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
