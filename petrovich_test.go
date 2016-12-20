package petrovich

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	gender   Gender
	infNames []string
}

func TestFirstName(t *testing.T) {
	cases := []testCase{
		{"Анна-Мария", Female, []string{
			"Анны-Марии",
			"Анне-Марии",
			"Анну-Марию",
			"Анной-Марией",
			"Анне-Марии"}},
		{"Василий", Male, []string{
			"Василия",
			"Василию",
			"Василия",
			"Василием",
			"Василии"}},
		{"Кочубей", Male, []string{
			"Кочубея",
			"Кочубею",
			"Кочубея",
			"Кочубеем",
			"Кочубее"}},
		{"Лев", Male, []string{
			"Льва",
			"Льву",
			"Льва",
			"Львом",
			"Льве"}},
		{"Маша", Female, []string{
			"Маши",
			"Маше",
			"Машу",
			"Машей",
			"Маше"}},
	}

	for _, c := range cases {
		assert.Equal(t, c.name, FirstName(c.name, c.gender, Nomenative))
		for i, in := range c.infNames {
			assert.Equal(t, in, FirstName(c.name, c.gender, Case(i)))
		}
	}
}

func TestMiddleName(t *testing.T) {
	cases := []testCase{
		{"Георгиевна-Авраамовна", Female, []string{
			"Георгиевны-Авраамовны",
			"Георгиевне-Авраамовне",
			"Георгиевну-Авраамовну",
			"Георгиевной-Авраамовной",
			"Георгиевне-Авраамовне"}},
		{"Петрович", Male, []string{
			"Петровича",
			"Петровичу",
			"Петровича",
			"Петровичем",
			"Петровиче"}},
		{"Петровна", Female, []string{
			"Петровны",
			"Петровне",
			"Петровну",
			"Петровной",
			"Петровне"}},
	}

	for _, c := range cases {
		assert.Equal(t, c.name, MiddleName(c.name, c.gender, Nomenative))
		for i, in := range c.infNames {
			assert.Equal(t, in, MiddleName(c.name, c.gender, Case(i)))
		}
	}
}

func TestLastName(t *testing.T) {
	cases := []testCase{
		{"Андрейчук", Male, []string{
			"Андрейчука",
			"Андрейчуку",
			"Андрейчука",
			"Андрейчуком",
			"Андрейчуке"}},
		{"Андрейчук", Female, []string{
			"Андрейчук",
			"Андрейчук",
			"Андрейчук",
			"Андрейчук",
			"Андрейчук"}},
		{"Дюма", Male, []string{
			"Дюма",
			"Дюма",
			"Дюма",
			"Дюма",
			"Дюма"}},
		{"Петров-Водкин", Male, []string{
			"Петрова-Водкина",
			"Петрову-Водкину",
			"Петрова-Водкина",
			"Петровым-Водкиным",
			"Петрове-Водкине"}},
		{"Салтыков-Щедрин", Male, []string{
			"Салтыкова-Щедрина",
			"Салтыкову-Щедрину",
			"Салтыкова-Щедрина",
			"Салтыковым-Щедриным",
			"Салтыкове-Щедрине"}},
	}

	for _, c := range cases {
		assert.Equal(t, c.name, LastName(c.name, c.gender, Nomenative))
		for i, in := range c.infNames {
			assert.Equal(t, in, LastName(c.name, c.gender, Case(i)))
		}
	}
}
