package model

type Character struct {
	id   int
	name string

	// private final Level level;
	// private final TribeClassification tribe;

	// private final GenderClassification gender;

	// private final Equipment equipment;

	// private final Job job;
}

func NewCharacter(id int, name string) *Character {
	c := new(Character)
	c.id = id
	c.name = name
	return c
}

func (c Character) GetId() int {
	return c.id
}
func (c Character) GetName() string {
	return c.name
}
