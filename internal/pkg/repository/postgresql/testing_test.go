package postgresql

import (
	"testing"

	"homework-3/internal/pkg/app"
	mock_database "homework-3/internal/pkg/db/mocks"

	"github.com/golang/mock/gomock"
)

type authorsRepoFixture struct {
	ctrl   *gomock.Controller
	repo   app.AuthorsRepo
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
