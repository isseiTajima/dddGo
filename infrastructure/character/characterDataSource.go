package infrastructure

import "dddgo/domain/model"

type CharacterDataSource struct {
	characterList []model.Character
}

func (c *CharacterDataSource) InsertCharacter(name string) {
	c.characterList = append(c.characterList, *model.NewCharacter(name))

}

func (c *CharacterDataSource) FindById(userId string) *model.Character {

	for _, chara := range c.characterList {
		if chara.GetId() == userId {
			return &chara
		}
	}

	return nil
}
