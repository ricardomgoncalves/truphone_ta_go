// +build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
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
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		fam := family.Family{
			Id:          id,
			Name:        "Family Test",
			CountryCode: "ES",
		}

		err := repo.CreateFamily(context.Background(), fam)
		require.Nil(t, err)
	})
	t.Run("should return error on create", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		fam := family.Family{
			Id:          id,
			Name:        "Family Test",
			CountryCode: "ES",
		}

		err := repo.CreateFamily(context.Background(), fam)
		require.NotNil(t, err)
	})
}

func TestRepo_GetFamilyById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should return family by id", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038"
		fam, err := repo.GetFamilyById(context.Background(), id)
		require.Nil(t, err)
		assert.Equal(t, id, fam.Id)
		assert.Equal(t, "Family 1", fam.Name)
		assert.Equal(t, "PT", fam.CountryCode)
	})
	t.Run("should return error", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da037"
		_, err := repo.GetFamilyById(context.Background(), id)
		require.NotNil(t, err)
	})
}

func TestRepo_ListFamilies(t *testing.T) {
	rp := getRepo(t)
	t.Run("should return all families", func(t *testing.T) {
		options := make([]repo.FilterOption, 0)
		fams, err := rp.ListFamilies(context.Background(), options...)
		t.Log(fams)
		require.Nil(t, err)
		assert.Equal(t, 2, len(fams))
	})
	t.Run("should return all families", func(t *testing.T) {
		options := make([]repo.FilterOption, 0)
		options = append(options, repo.WithCountryCode("ES"))

		fams, err := rp.ListFamilies(context.Background(), options...)
		require.Nil(t, err)
		assert.Equal(t, 1, len(fams))
	})
	t.Run("should return all families", func(t *testing.T) {
		options := make([]repo.FilterOption, 0)
		options = append(options, repo.WithLimit(1))

		fams, err := rp.ListFamilies(context.Background(), options...)
		require.Nil(t, err)
		assert.Equal(t, 1, len(fams))
	})
	t.Run("should return all families", func(t *testing.T) {
		options := make([]repo.FilterOption, 0)
		options = append(options, repo.WithOffset(1))

		fams, err := rp.ListFamilies(context.Background(), options...)
		require.Nil(t, err)
		assert.Equal(t, 1, len(fams))
	})
}

func TestRepo_UpdateFamilyById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should update record", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		err := repo.UpdateFamilyById(context.Background(), id, family.Family{
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
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da03c"
		err := repo.UpdateFamilyById(context.Background(), id, family.Family{
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
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		err := repo.DeleteFamilyById(context.Background(), id)
		require.Nil(t, err)

		_, err = repo.GetFamilyById(context.Background(), id)
		require.NotNil(t, err)
	})
	t.Run("should return not found error", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		err := repo.DeleteFamilyById(context.Background(), id)
		require.NotNil(t, err)
		assert.Equal(t, family.ErrorFamilyNotFound, err)
	})
}

func TestRepo_CreateMember(t *testing.T) {
	repo := getRepo(t)
	t.Run("should create member", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
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

		err := repo.CreateMember(context.Background(), member)
		require.Nil(t, err)
	})
	t.Run("should return error on create", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
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

		err := repo.CreateMember(context.Background(), member)
		require.NotNil(t, err)
	})
}

func TestRepo_GetMemberById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should return member by id", func(t *testing.T) {
		id := "11810f35-309a-4836-b7e9-1fee57bed924"
		member, err := repo.GetMemberById(context.Background(), id)
		require.Nil(t, err)
		assert.Equal(t, id, member.Id)
		assert.Equal(t, "Father", member.FirstName)
		assert.Equal(t, "1", member.LastName)
	})
	t.Run("should return error", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da037"
		_, err := repo.GetMemberById(context.Background(), id)
		require.NotNil(t, err)
	})
}

func TestRepo_UpdateMemberById(t *testing.T) {
	repo := getRepo(t)
	t.Run("should update record", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		err := repo.UpdateMemberById(context.Background(), id, family.Member{
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
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da03c"
		err := repo.UpdateMemberById(context.Background(), id, family.Member{
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
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		err := repo.DeleteMemberById(context.Background(), id)
		require.Nil(t, err)

		_, err = repo.GetMemberById(context.Background(), id)
		require.NotNil(t, err)
	})
	t.Run("should return not found error", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da036"
		err := repo.DeleteMemberById(context.Background(), id)
		require.NotNil(t, err)
		assert.Equal(t, family.ErrorMemberNotFound, err)
	})
}
