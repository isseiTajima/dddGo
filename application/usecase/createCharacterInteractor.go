package usecase

import (
	"dddgo/domain/model"
	"dddgo/domain/repository"
	"dddgo/usecase"
)

// CreateCharacterUsecaseの構造体
type CreateCharacterUsecase struct {
	r repository.CharacterRepository
}

func NewCreateCharacterUsecase(r repository.CharacterRepository) usecase.CreateCharacterUsecase {
	return &CreateCharacterUsecase{r}
}

func (c *CreateCharacterUsecase) Create(name string) (*model.Character, error) {
	return c.r.Insert(name)
}
