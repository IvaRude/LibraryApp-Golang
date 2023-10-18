package routers

import (
	"context"
	"net/http"
	"testing"

	"homework-3/internal/pkg/repository"
	mock_repository "homework-3/internal/pkg/repository/mocks"
	"homework-3/internal/pkg/server"
	"homework-3/tests/fixtures"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetAuthor(t *testing.T) {
	var (
		ctx context.Context
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().GetByID(gomock.Any(), int64(id)).Return(fixtures.Author().Valid().P(), nil)
		//act
		author, code := GetAuthor(ctx, s, int64(id))
		// assert
		require.Equal(t, http.StatusOK, int(code))
		assert.Equal(t, "{\"Id\":500001,\"Name\":\"Author 1\",\"Books\":null}", string(author))
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
			mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
			s := server.NewServer(mockAuthorsRepo, mockBookRepo)
			mockAuthorsRepo.EXPECT().GetByID(gomock.Any(), int64(id)).Return(nil, repository.ErrObjectNotFound)
			//act
			author, code := GetAuthor(ctx, s, int64(id))
			// assert
			require.Equal(t, http.StatusNotFound, int(code))
			assert.Nil(t, author)
		})
		t.Run("interal error", func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
			mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
			s := server.NewServer(mockAuthorsRepo, mockBookRepo)
			mockAuthorsRepo.EXPECT().GetByID(gomock.Any(), int64(id)).Return(nil, assert.AnError)
			//act
			author, code := GetAuthor(ctx, s, int64(id))
			// assert
			require.Equal(t, http.StatusInternalServerError, int(code))
			assert.Nil(t, author)
		})
	})
}

func Test_CreateAuthor(t *testing.T) {
	var (
		ctx        context.Context
		authorData = &updateAuthorRequest{}
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(0), nil)
		//act
		code := CreateAuthor(ctx, s, authorData)
		// assert
		require.Equal(t, http.StatusOK, int(code))
	})
	t.Run("interal error", func(t *testing.T) {
		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)
		//act
		code := CreateAuthor(ctx, s, authorData)
		// assert
		require.Equal(t, http.StatusInternalServerError, int(code))
	})
}

func Test_UpdateAuthor(t *testing.T) {
	var (
		ctx        context.Context
		authorData = &updateAuthorRequest{}
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		//act
		code := UpdateAuthor(ctx, s, authorData)
		// assert
		require.Equal(t, http.StatusOK, int(code))
	})
	t.Run("not found", func(t *testing.T) {
		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(repository.ErrObjectNotFound)
		//act
		code := UpdateAuthor(ctx, s, authorData)
		// assert
		require.Equal(t, http.StatusNotFound, int(code))
	})
	t.Run("interal error", func(t *testing.T) {
		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(assert.AnError)
		//act
		code := UpdateAuthor(ctx, s, authorData)
		// assert
		require.Equal(t, http.StatusInternalServerError, int(code))
	})
}

func Test_DeleteAuthor(t *testing.T) {
	var (
		ctx context.Context
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
		mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
		s := server.NewServer(mockAuthorsRepo, mockBookRepo)
		mockAuthorsRepo.EXPECT().DeleteById(gomock.Any(), int64(id)).Return(nil)
		//act
		code := DeleteAuthor(ctx, s, int64(id))
		// assert
		require.Equal(t, http.StatusOK, int(code))
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
			mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
			s := server.NewServer(mockAuthorsRepo, mockBookRepo)
			mockAuthorsRepo.EXPECT().DeleteById(gomock.Any(), int64(id)).Return(repository.ErrObjectNotFound)
			//act
			code := DeleteAuthor(ctx, s, int64(id))
			// assert
			require.Equal(t, http.StatusNotFound, int(code))
		})
		t.Run("interal error", func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthorsRepo := mock_repository.NewMockAuthorsRepo(ctrl)
			mockBookRepo := mock_repository.NewMockBooksRepo(ctrl)
			s := server.NewServer(mockAuthorsRepo, mockBookRepo)
			mockAuthorsRepo.EXPECT().DeleteById(gomock.Any(), int64(id)).Return(assert.AnError)
			//act
			code := DeleteAuthor(ctx, s, int64(id))
			// assert
			require.Equal(t, http.StatusInternalServerError, int(code))
		})
	})
}
