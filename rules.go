package petrovich

type rules struct {
	FirstName  rulesGroup
	MiddleName rulesGroup
	LastName   rulesGroup
}

type rulesGroup struct {
	Exceptions []rule
	Suffixes   []rule
}

type rule struct {
	Gender string
	Test   []string
	Mods   []string
}

func (r rule) getGender() Gender {
	switch r.Gender {
	case "male":
		return Male
	case "female":
		return Female
	}
	return Androgynous
}
