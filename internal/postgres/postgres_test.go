package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepo_CheckFamilyError(t *testing.T) {
	t.Run("should return nil on nil error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)
		assert.Nil(t, repo.checkFamilyError(nil))
	})
	t.Run("should return already exists error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code: "40002"}
		outErr := repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyAlreadyExists, err))

		err = &pq.Error{Code: "42710"}
		outErr = repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyAlreadyExists, err))

		err = &pq.Error{Code: "23505"}
		outErr = repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyAlreadyExists, err))
	})
	t.Run("should return locked error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code: "55006"}
		outErr := repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyLocked, err))
	})
	t.Run("should return unknown error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code: "123123"}
		outErr := repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyUnknown, err))

		randomErr := errors.New("random error")
		outErr = repo.checkFamilyError(randomErr)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyUnknown, randomErr))
	})
	t.Run("should return itself", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		outErr := repo.checkFamilyError(family.ErrorFamilyNotFound)
		require.NotNil(t, outErr)
		assert.True(t, outErr == family.ErrorFamilyNotFound)
	})
	t.Run("should return not found error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		outErr := repo.checkFamilyError(gorm.ErrRecordNotFound)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorFamilyNotFound, gorm.ErrRecordNotFound))
	})
}

func TestRepo_CheckMemberError(t *testing.T) {
	t.Run("should return nil on nil error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)
		assert.Nil(t, repo.checkMemberError(nil))
	})
	t.Run("should return already exists error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code: "40002"}
		outErr := repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberAlreadyExists, err))

		err = &pq.Error{Code: "42710"}
		outErr = repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberAlreadyExists, err))

		err = &pq.Error{Code: "23505"}
		outErr = repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberAlreadyExists, err))
	})
	t.Run("should return locked error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code: "55006"}
		outErr := repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberLocked, err))
	})
	t.Run("should return unknown error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code: "123123"}
		outErr := repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberUnknown, err))

		randomErr := errors.New("random error")
		outErr = repo.checkMemberError(randomErr)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberUnknown, randomErr))
	})
	t.Run("should return itself", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		outErr := repo.checkMemberError(family.ErrorMemberNotFound)
		require.NotNil(t, outErr)
		assert.True(t, outErr == family.ErrorMemberNotFound)
	})
	t.Run("should return not found error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		outErr := repo.checkMemberError(gorm.ErrRecordNotFound)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(family.ErrorMemberNotFound, gorm.ErrRecordNotFound))
	})
}
