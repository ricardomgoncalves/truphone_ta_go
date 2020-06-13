package duplicates

import (
	"testing"
	"time"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestFindPossibleDuplicates(t *testing.T) {
	t.Run("should return 2 duplicates with only roots", func(t *testing.T) {
		input := []family.Member{
			{
				Id: "id",
			},
			{
				Id: "id2",
			},
			{
				Id:         "id3",
				FirstName:  "First",
				MiddleName: "middle",
				LastName:   "last",
			},
		}
		assert.Equal(t, []family.Member{{Id: "id"}, {Id: "id2"}}, FindPossibleDuplicates(input))
	})
	t.Run("should return 4 duplicates", func(t *testing.T) {
		fatherId := "id"
		motherId := "id2"
		input := []family.Member{
			{
				Id: "id",
			},
			{
				Id: "id2",
			},
			{
				Id:         "id3",
				FirstName:  "First",
				MiddleName: "middle",
				LastName:   "last",
			},
			{
				Id:       "id4",
				FatherId: &fatherId,
				MotherId: &motherId,
			},
			{
				Id:       "id5",
				FatherId: &fatherId,
				MotherId: &motherId,
			},
		}
		assert.Equal(t, []family.Member{
			{Id: "id"},
			{Id: "id2"},
			{Id: "id4", FatherId: &fatherId, MotherId: &motherId},
			{Id: "id5", FatherId: &fatherId, MotherId: &motherId},
		}, FindPossibleDuplicates(input))
	})
}

func TestIsPossibleDuplicates(t *testing.T) {
	t.Run("should get false when no name is on common", func(t *testing.T) {
		assert.False(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{}, family.Member{
			FirstName:  "First",
			MiddleName: "Middle",
			LastName:   "Last",
		}))
	})
	t.Run("should return false with different birthdays", func(t *testing.T) {
		assert.False(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{
			Birthday: time.Now().Add(time.Second * 30),
		}, family.Member{
			Birthday: time.Now(),
		}))
	})
	t.Run("should return true when one of parents is missing", func(t *testing.T) {
		motherId := "id"
		fatherId := "id"
		assert.True(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{}, family.Member{}))
		assert.True(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{
			MotherId: &motherId,
			FatherId: &fatherId,
		}, family.Member{}))
		assert.True(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			MotherId: &motherId,
		}))
	})
	t.Run("should return true with the same parents", func(t *testing.T) {
		motherId := "id"
		fatherId := "id"
		assert.True(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{
			MotherId: &motherId,
			FatherId: &fatherId,
		}, family.Member{
			MotherId: &motherId,
			FatherId: &fatherId,
		}))
	})
	t.Run("should return false with the different parents", func(t *testing.T) {
		motherId := "id"
		fatherId := "id"
		motherId1 := "id1"
		fatherId1 := "id1"
		assert.False(t, isPossibleDuplicates(map[string]struct{}{}, family.Member{
			MotherId: &motherId,
			FatherId: &fatherId,
		}, family.Member{
			MotherId: &motherId1,
			FatherId: &fatherId1,
		}))
	})
}

func TestCheckMembersMothers(t *testing.T) {
	t.Run("should return true on same father", func(t *testing.T) {
		assert.True(t, checkMembersFathers(map[string]struct{}{}, family.Member{
			FatherId: nil,
		}, family.Member{
			FatherId: nil,
		}))
		motherId := "id"
		assert.True(t, checkMembersMothers(map[string]struct{}{}, family.Member{
			MotherId: &motherId,
		}, family.Member{
			MotherId: &motherId,
		}))
		motherId2 := "id"
		assert.True(t, checkMembersMothers(map[string]struct{}{}, family.Member{
			MotherId: &motherId,
		}, family.Member{
			MotherId: &motherId2,
		}))
	})
	t.Run("should return true when one is missing", func(t *testing.T) {
		motherId := "id"
		assert.True(t, checkMembersMothers(map[string]struct{}{}, family.Member{
			MotherId: nil,
		}, family.Member{
			MotherId: &motherId,
		}))
		assert.True(t, checkMembersMothers(map[string]struct{}{}, family.Member{
			MotherId: &motherId,
		}, family.Member{
			MotherId: nil,
		}))
	})
	t.Run("should return true when both are possible duplicated", func(t *testing.T) {
		possibleDuplicatesIds := map[string]struct{}{
			"id":  {},
			"id2": {},
		}
		motherId := "id"
		motherId2 := "id2"
		assert.True(t, checkMembersMothers(possibleDuplicatesIds, family.Member{
			MotherId: &motherId,
		}, family.Member{
			MotherId: &motherId2,
		}))
		assert.True(t, checkMembersMothers(possibleDuplicatesIds, family.Member{
			MotherId: &motherId2,
		}, family.Member{
			MotherId: &motherId,
		}))
	})
	t.Run("should return false when only one is possible duplicated", func(t *testing.T) {
		possibleDuplicatesIds := map[string]struct{}{
			"id": {},
		}
		possibleDuplicatesIds2 := map[string]struct{}{
			"id2": {},
		}
		motherId := "id"
		motherId2 := "id2"
		assert.False(t, checkMembersMothers(possibleDuplicatesIds, family.Member{
			MotherId: &motherId,
		}, family.Member{
			MotherId: &motherId2,
		}))
		assert.False(t, checkMembersMothers(possibleDuplicatesIds, family.Member{
			MotherId: &motherId2,
		}, family.Member{
			MotherId: &motherId,
		}))
		assert.False(t, checkMembersMothers(possibleDuplicatesIds2, family.Member{
			MotherId: &motherId,
		}, family.Member{
			MotherId: &motherId2,
		}))
		assert.False(t, checkMembersMothers(possibleDuplicatesIds2, family.Member{
			MotherId: &motherId2,
		}, family.Member{
			MotherId: &motherId,
		}))
	})
}

func TestCheckMembersFathers(t *testing.T) {
	t.Run("should return true on same father", func(t *testing.T) {
		assert.True(t, checkMembersFathers(map[string]struct{}{}, family.Member{
			FatherId: nil,
		}, family.Member{
			FatherId: nil,
		}))
		fatherId := "id"
		assert.True(t, checkMembersFathers(map[string]struct{}{}, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			FatherId: &fatherId,
		}))
		fatherId2 := "id"
		assert.True(t, checkMembersFathers(map[string]struct{}{}, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			FatherId: &fatherId2,
		}))
	})
	t.Run("should return true when one is missing", func(t *testing.T) {
		fatherId := "id"
		assert.True(t, checkMembersFathers(map[string]struct{}{}, family.Member{
			FatherId: nil,
		}, family.Member{
			FatherId: &fatherId,
		}))
		assert.True(t, checkMembersFathers(map[string]struct{}{}, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			FatherId: nil,
		}))
	})
	t.Run("should return true when both are possible duplicated", func(t *testing.T) {
		possibleDuplicatesIds := map[string]struct{}{
			"id":  {},
			"id2": {},
		}
		fatherId := "id"
		fatherId2 := "id2"
		assert.True(t, checkMembersFathers(possibleDuplicatesIds, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			FatherId: &fatherId2,
		}))
		assert.True(t, checkMembersFathers(possibleDuplicatesIds, family.Member{
			FatherId: &fatherId2,
		}, family.Member{
			FatherId: &fatherId,
		}))
	})
	t.Run("should return false when only one is possible duplicated", func(t *testing.T) {
		possibleDuplicatesIds := map[string]struct{}{
			"id": {},
		}
		possibleDuplicatesIds2 := map[string]struct{}{
			"id2": {},
		}
		fatherId := "id"
		fatherId2 := "id2"
		assert.False(t, checkMembersFathers(possibleDuplicatesIds, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			FatherId: &fatherId2,
		}))
		assert.False(t, checkMembersFathers(possibleDuplicatesIds, family.Member{
			FatherId: &fatherId2,
		}, family.Member{
			FatherId: &fatherId,
		}))
		assert.False(t, checkMembersFathers(possibleDuplicatesIds2, family.Member{
			FatherId: &fatherId,
		}, family.Member{
			FatherId: &fatherId2,
		}))
		assert.False(t, checkMembersFathers(possibleDuplicatesIds2, family.Member{
			FatherId: &fatherId2,
		}, family.Member{
			FatherId: &fatherId,
		}))
	})
}
