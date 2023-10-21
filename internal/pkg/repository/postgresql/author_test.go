package postgresql

import (
	"context"
	"database/sql"
	"homework-3/internal/pkg/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name FROM authors WHERE id=$1;", gomock.Any()).Return(nil)
		s.mockDb.EXPECT().Select(gomock.Any(), gomock.Any(), "SELECT id, name, author_id FROM books WHERE author_id=$1;", gomock.Any()).Return(sql.ErrNoRows)
		// act
		author, err := s.repo.GetByID(ctx, int64(id))
		// assert

		require.NoError(t, err)
		assert.Equal(t, int64(0), author.Id)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name FROM authors WHERE id=$1;", gomock.Any()).Return(sql.ErrNoRows)

			// act
			author, err := s.repo.GetByID(ctx, int64(id))
			// assert
			require.EqualError(t, err, "object not found")

			assert.Nil(t, author)
		})
	})
}

func TestAdd(t *testing.T) {
	t.Parallel()
	var (
		ctx    = context.Background()
		author = &models.Author{Name: "Name"}
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "INSERT INTO authors(name) VALUES($1) RETURNING id;", gomock.Any()).Return(nil)

		_, err := s.repo.Add(ctx, author)
		require.NoError(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "WITH q AS (DELETE FROM books WHERE author_id = $1) DELETE FROM authors WHERE id = $1 RETURNING id;", gomock.Any()).Return(nil)

		// act
		err := s.repo.DeleteById(ctx, int64(id))
		// assert

		require.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "WITH q AS (DELETE FROM books WHERE author_id = $1) DELETE FROM authors WHERE id = $1 RETURNING id;", gomock.Any()).Return(sql.ErrNoRows)

			// act
			err := s.repo.DeleteById(ctx, int64(id))
			// assert
			require.EqualError(t, err, "object not found")

		})
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	var (
		ctx    = context.Background()
		author = &models.Author{Name: "Name", Id: 1}
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "UPDATE authors SET name = $1 WHERE id = $2 RETURNING id;", gomock.Any(), gomock.Any()).Return(nil)

		err := s.repo.Update(ctx, author)
		require.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "UPDATE authors SET name = $1 WHERE id = $2 RETURNING id;", gomock.Any(), gomock.Any()).Return(sql.ErrNoRows)
			// act
			err := s.repo.Update(ctx, author)
			// assert
			require.EqualError(t, err, "object not found")

		})
	})
}
