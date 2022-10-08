package repository

import "dddgo/domain/model"

type CharacterRepository interface {
	Insert(name string) (*model.Character, error)
	FindById(userId string) (*model.Character, error)
}
