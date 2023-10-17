package postgresql

import (
	"testing"

	mock_database "homework-3/internal/pkg/db/mocks"

	"homework-3/internal/pkg/repository"

	"github.com/golang/mock/gomock"
)

type authorsRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.AuthorsRepo
	mockDb *mock_database.MockDBops
}

func setUp(t *testing.T) authorsRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDBops(ctrl)
	repo := NewAuthors(mockDb)
	return authorsRepoFixture{
		ctrl:   ctrl,
		repo:   repo,
		mockDb: mockDb,
	}
}

func (a *authorsRepoFixture) tearDown() {
	a.ctrl.Finish()
}
