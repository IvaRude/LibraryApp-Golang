package tests

import (
	"context"
	"testing"

	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/repository/postgresql"
	"homework-3/tests/fixtures"
	"homework-3/tests/states"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)

		//act
		resp, err := repo.Add(ctx, fixtures.Author().Valid().P())

		//assert
		require.NoError(t, err)
		assert.NotZero(t, resp)
	})
}

func TestGetAuthor(t *testing.T) {
	var (
		ctx         = context.Background()
		AuthorValid = fixtures.Author().Valid().P()
	)
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)
		resp, err := repo.Add(ctx, AuthorValid)

		require.NoError(t, err)
		assert.NotZero(t, resp)
		//act
		respGet, err := repo.GetByID(ctx, resp)

		//assert
		require.NoError(t, err)
		assert.Equal(t, AuthorValid.Name, respGet.Name)
	})
	t.Run("not found", func(t *testing.T) {
		db.SetUp(t, AuthorValid)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)
		//act
		respGet, err := repo.GetByID(ctx, states.WrongAuthorID)

		//assert
		require.EqualError(t, err, repository.ErrObjectNotFound.Error())
		assert.Nil(t, respGet)
	})
}

func TestUpdateAuthor(t *testing.T) {
	var (
		ctx        = context.Background()
		auhtorData = fixtures.Author().Valid()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)
		id, err := repo.Add(ctx, auhtorData.P())
		//act
		err = repo.Update(ctx, auhtorData.ID(id).Name(states.AuthorName2).P())
		//assert
		require.NoError(t, err)
		authorUpdated, err := repo.GetByID(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, states.AuthorName2, authorUpdated.Name)
	})
	t.Run("not found", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)
		//act
		err := repo.Update(ctx, auhtorData.ID(states.WrongAuthorID).P())

		//assert
		require.EqualError(t, err, repository.ErrObjectNotFound.Error())
	})
}

func TestDeleteAuthor(t *testing.T) {
	var (
		ctx        = context.Background()
		auhtorData = fixtures.Author().Valid()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)
		id, err := repo.Add(ctx, auhtorData.P())
		//act
		err = repo.DeleteById(ctx, id)
		//assert
		require.NoError(t, err)
		authorDeleted, err := repo.GetByID(ctx, id)
		require.Error(t, err, repository.ErrObjectNotFound)
		assert.Nil(t, authorDeleted)
	})
	t.Run("not found", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewAuthors(db.DB)
		//act
		err := repo.DeleteById(ctx, states.WrongAuthorID)

		//assert
		require.Error(t, err, repository.ErrObjectNotFound)
	})
}
