package petrovich

type rules struct {
	firstName  rulesGroup
	middleName rulesGroup
	lastName   rulesGroup
}

type rulesGroup struct {
	exceptions []rule
	suffixes   []rule
}

type rule struct {
	gender Gender
	test   []string
	mods   []string
}
