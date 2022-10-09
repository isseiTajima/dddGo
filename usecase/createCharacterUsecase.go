package usecase

import "dddgo/domain/model"

type CreateCharacterUsecase interface {
	Create(name string) (*model.Character, error)
}
