// +build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getRepo(t *testing.T) *Repo {
	t.Helper()

	postgresConnectionUrl, err := BuildConnectionString(
		"//localhost:5432",
		"disable",
		"postgres",
		"postgres",
		"postgres",
	)
	require.Nil(t, err)

	db, err := gorm.Open("postgres", postgresConnectionUrl)
	require.Nil(t, err)

	if err := CreateTables(db); err != nil {
		t.Log(err)
	}

	if err := Populate(db); err != nil {
		t.Log(err)
	}

	return NewPostgresRepo(db)
}

func TestRepo_CreateFamily(t *testing.T) {
	repo := getRepo(t)
	t.Run("should create family without any error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		fam := family.Family{
			Id:          id,
			Name:        "Family Test",
			CountryCode: "ES",
		}

		err = repo.CreateFamily(context.Background(), fam)
		require.Nil(t, err)
	})
	t.Run("should return error on create", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		fam := family.Family{
			Id:          id,
			Name:        "Family Test",
			CountryCode: "ES",
		}

		err = repo.CreateFamily(context.Background(), fam)
		require.NotNil(t, err)
	})
}

func TestRepo_GetFamilyById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should return family by id", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038")
		require.Nil(t, err)

		fam, err := repo.GetFamilyById(context.Background(), id)
		require.Nil(t, err)
		assert.Equal(t, id, fam.Id)
		assert.Equal(t, "Family 1", fam.Name)
		assert.Equal(t, "PT", fam.CountryCode)
	})
	t.Run("should return error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da037")
		require.Nil(t, err)

		_, err = repo.GetFamilyById(context.Background(), id)
		require.NotNil(t, err)
	})
}

func TestRepo_UpdateFamilyById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should update record", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		err = repo.UpdateFamilyById(context.Background(), id, family.Family{
			Id:          id,
			Name:        "Family Updated",
			CountryCode: "UK",
		})
		require.Nil(t, err)

		rFam, err := repo.GetFamilyById(context.Background(), id)
		require.Nil(t, err)

		assert.Equal(t, id, rFam.Id)
		assert.Equal(t, "Family Updated", rFam.Name)
		assert.Equal(t, "UK", rFam.CountryCode)
	})
	t.Run("should return not found error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da03c")
		require.Nil(t, err)

		err = repo.UpdateFamilyById(context.Background(), id, family.Family{
			Id:          id,
			Name:        "Family Updated",
			CountryCode: "UK",
		})
		require.NotNil(t, err)
		assert.Equal(t, family.ErrorFamilyNotFound, err)
	})
}

func TestRepo_DeleteFamilyById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should delete record", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		err = repo.DeleteFamilyById(context.Background(), id)
		require.Nil(t, err)

		_, err = repo.GetFamilyById(context.Background(), id)
		require.NotNil(t, err)
	})
	t.Run("should return not found error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		err = repo.DeleteFamilyById(context.Background(), id)
		require.NotNil(t, err)
		assert.Equal(t, family.ErrorFamilyNotFound, err)
	})
}

func TestRepo_CreateMember(t *testing.T) {
	repo := getRepo(t)
	t.Run("should create member", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		member := family.Member{
			Id:         id,
			FamilyId:   id,
			FirstName:  "Member",
			MiddleName: "-",
			LastName:   "Test",
			FatherId:   &id,
			MotherId:   &id,
			SpouseId:   &id,
			Birthday:   time.Now(),
		}

		err = repo.CreateMember(context.Background(), member)
		require.Nil(t, err)
	})
	t.Run("should return error on create", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		member := family.Member{
			Id:         id,
			FamilyId:   id,
			FirstName:  "Member",
			MiddleName: "-",
			LastName:   "Test",
			FatherId:   &id,
			MotherId:   &id,
			SpouseId:   &id,
			Birthday:   time.Now(),
		}

		err = repo.CreateMember(context.Background(), member)
		require.NotNil(t, err)
	})
}

func TestRepo_GetMemberById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should return member by id", func(t *testing.T) {
		id, err := uuid.Parse("11810f35-309a-4836-b7e9-1fee57bed924")
		require.Nil(t, err)

		member, err := repo.GetMemberById(context.Background(), id)
		require.Nil(t, err)
		assert.Equal(t, id, member.Id)
		assert.Equal(t, "Father", member.FirstName)
		assert.Equal(t, "1", member.LastName)
	})
	t.Run("should return error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da037")
		require.Nil(t, err)

		_, err = repo.GetMemberById(context.Background(), id)
		require.NotNil(t, err)
	})
}

func TestRepo_GetMembersByFamilyId(t *testing.T) {
	repo := getRepo(t)

	id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038")
	require.Nil(t, err)

	t.Run("should return list of family member", func(t *testing.T) {
		fam, err := repo.GetMembersByFamilyId(context.Background(), id, nil, nil)
		require.Nil(t, err)
		assert.Equal(t, 3, len(fam))
	})
	t.Run("should return list with correct limit", func(t *testing.T) {
		limit := 2
		fam, err := repo.GetMembersByFamilyId(context.Background(), id, nil, &limit)
		require.Nil(t, err)
		assert.Equal(t, 2, len(fam))
	})
	t.Run("should return list with correct offset", func(t *testing.T) {
		offset := 1
		fam, err := repo.GetMembersByFamilyId(context.Background(), id, &offset, nil)
		require.Nil(t, err)
		assert.Equal(t, 2, len(fam))
	})
	t.Run("should return list and ignore limit", func(t *testing.T) {
		limit := 0
		fam, err := repo.GetMembersByFamilyId(context.Background(), id, nil, &limit)
		require.Nil(t, err)
		assert.Equal(t, 3, len(fam))
	})
	t.Run("should return list and ignore offset", func(t *testing.T) {
		offset := 0
		fam, err := repo.GetMembersByFamilyId(context.Background(), id, &offset, nil)
		require.Nil(t, err)
		assert.Equal(t, 3, len(fam))
	})
	t.Run("should return list with offset and limit", func(t *testing.T) {
		offset := 1
		limit := 1
		fam, err := repo.GetMembersByFamilyId(context.Background(), id, &offset, &limit)
		require.Nil(t, err)
		assert.Equal(t, 1, len(fam))
	})
}

func TestRepo_UpdateMemberById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should update record", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		err = repo.UpdateMemberById(context.Background(), id, family.Member{
			Id:       id,
			LastName: "Updated",
		})
		require.Nil(t, err)

		rFam, err := repo.GetMemberById(context.Background(), id)
		require.Nil(t, err)

		assert.Equal(t, id, rFam.Id)
		assert.Equal(t, "Updated", rFam.LastName)
	})
	t.Run("should return not found error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da03c")
		require.Nil(t, err)

		err = repo.UpdateMemberById(context.Background(), id, family.Member{
			Id:       id,
			LastName: "Updated",
		})
		require.NotNil(t, err)
		assert.Equal(t, family.ErrorMemberNotFound, err)
	})
}

func TestRepo_DeleteMemberById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should delete record", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		err = repo.DeleteMemberById(context.Background(), id)
		require.Nil(t, err)

		_, err = repo.GetMemberById(context.Background(), id)
		require.NotNil(t, err)
	})
	t.Run("should return not found error", func(t *testing.T) {
		id, err := uuid.Parse("9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036")
		require.Nil(t, err)

		err = repo.DeleteMemberById(context.Background(), id)
		require.NotNil(t, err)
		assert.Equal(t, family.ErrorMemberNotFound, err)
	})
}
