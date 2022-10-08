package infrastructure

import (
	"dddgo/domain/model"
	"dddgo/domain/repository"
)

// CharacterRepositoryの構造体
type CharacterRepository struct {
	characterList []model.Character
}

func NewCharacterRepository() repository.CharacterRepository {
	return &CharacterRepository{}
}

func (c *CharacterRepository) Insert(name string) (*model.Character, error) {
	chara := *model.NewCharacter(name)
	c.characterList = append(c.characterList, chara)
	return &chara, nil
}

func (c *CharacterRepository) FindById(userId string) (*model.Character, error) {

	for _, chara := range c.characterList {
		if chara.GetId() == userId {
			return &chara, nil
		}
	}

	return nil, nil
}
