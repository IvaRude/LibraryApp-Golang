package fixtures

import (
	"homework-3/internal/pkg/repository"
	"homework-3/tests/states"
)

type AuthorBuilder struct {
	instance *repository.Author
}

func Author() *AuthorBuilder {
	return &AuthorBuilder{instance: &repository.Author{}}
}

func (b *AuthorBuilder) ID(v int64) *AuthorBuilder {
	b.instance.Id = v
	return b
}
func (b *AuthorBuilder) Name(v string) *AuthorBuilder {
	b.instance.Name = v
	return b
}

func (b *AuthorBuilder) P() *repository.Author {
	return b.instance
}

func (b *AuthorBuilder) V() repository.Author {
	return *b.instance
}

func (b *AuthorBuilder) Valid() *AuthorBuilder {
	return Author().ID(states.Author1ID).Name(states.AuthorName1)
}
