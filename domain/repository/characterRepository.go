package repository

import "dddgo/domain/model"

type CharacterRepository interface {
	InsertCharacter(name string) error
	FindById(userId string) (*model.Character, error)
}
