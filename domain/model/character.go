package model

import (
	"github.com/google/uuid"
)

type Character struct {
	id   string
	name string

	// private final Level level;
	// private final TribeClassification tribe;

	// private final GenderClassification gender;

	// private final Equipment equipment;

	// private final Job job;
}

func NewCharacter(name string) *Character {
	c := new(Character)
	u := uuid.New()
	c.id = u.String()
	c.name = name
	return c
}

func (c Character) GetId() string {
	return c.id
}
func (c Character) GetName() string {
	return c.name
}
