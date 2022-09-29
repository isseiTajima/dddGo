package repository

import "github.com/isseiTajima/dddGo/domain/model"

type CharacterRepository interface {
	Insert(userId, name) error
	FindById(userId string) (*domain.character, error)
}

