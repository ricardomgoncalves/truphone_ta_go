package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/errors"
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

		err := &pq.Error{Code:"40002"}
		outErr := repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorFamilyAlreadyExists))

		err = &pq.Error{Code:"42710"}
		outErr = repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorFamilyAlreadyExists))

		err = &pq.Error{Code:"23505"}
		outErr = repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorFamilyAlreadyExists))
	})
	t.Run("should return locked error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code:"55006"}
		outErr := repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorFamilyLocked))
	})
	t.Run("should return unknown error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code:"123123"}
		outErr := repo.checkFamilyError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorFamilyUnknown))

		randomErr := errors.New("random error")
		outErr = repo.checkFamilyError(randomErr)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(randomErr, family.ErrorFamilyUnknown))
	})
	t.Run("should return not found error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		outErr := repo.checkFamilyError(gorm.ErrRecordNotFound)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(gorm.ErrRecordNotFound, family.ErrorFamilyNotFound))
	})
}

func TestRepo_CheckMemberError(t *testing.T) {
	t.Run("should return nil on nil error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)
		assert.Nil(t, repo.checkMemberError(nil))
	})
	t.Run("should return already exists error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code:"40002"}
		outErr := repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorMemberAlreadyExists))

		err = &pq.Error{Code:"42710"}
		outErr = repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorMemberAlreadyExists))

		err = &pq.Error{Code:"23505"}
		outErr = repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorMemberAlreadyExists))
	})
	t.Run("should return locked error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code:"55006"}
		outErr := repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorMemberLocked))
	})
	t.Run("should return unknown error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		err := &pq.Error{Code:"123123"}
		outErr := repo.checkMemberError(err)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(err, family.ErrorMemberUnknown))

		randomErr := errors.New("random error")
		outErr = repo.checkMemberError(randomErr)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(randomErr, family.ErrorMemberUnknown))
	})
	t.Run("should return not found error", func(t *testing.T) {
		repo := NewPostgresRepo(nil)

		outErr := repo.checkMemberError(gorm.ErrRecordNotFound)
		require.NotNil(t, outErr)
		assert.True(t, outErr == errors.Wrap(gorm.ErrRecordNotFound, family.ErrorMemberNotFound))
	})
}
