package repository

import "dddgo/domain/model"

type CharacterRepository interface {
	Insert(userId int, name string) error
	FindById(userId string) (*model.Character, error)
}
