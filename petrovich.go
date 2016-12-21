// Package petrovich provides methods to inflect Russian first, last, and middle names.
package petrovich

//go:generate go run rules-generator/rules-generator.go

import (
	"strings"
	"unicode/utf8"
)

type Gender int

const (
	Androgynous Gender = iota
	Male
	Female
)

type Case int

const (
	Nomenative Case = -1 + iota
	Genitive
	Dative
	Accusative
	Instrumental
	Prepositional
)

// FirstName inflects first name depending on the given gender and case.
func FirstName(name string, g Gender, c Case) string {
	return inflect(name, g, c, allRules.firstName)
}

// MiddleName inflects middle name depending on the given gender and case.
func MiddleName(name string, g Gender, c Case) string {
	return inflect(name, g, c, allRules.middleName)
}

// LastName inflects last name depending on the given gender and case.
func LastName(name string, g Gender, c Case) string {
	return inflect(name, g, c, allRules.lastName)
}

func inflect(name string, g Gender, c Case, rg rulesGroup) string {
	if c == Nomenative {
		return name
	}

	ns := strings.Split(name, "-")
	r := make([]string, len(ns))

	for i, n := range ns {
		if e := checkException(n, g, c, rg); e != "" {
			r[i] = e
		} else {
			r[i] = findInRules(n, g, c, rg)
		}
	}

	return strings.Join(r, "-")
}

func checkException(name string, g Gender, c Case, rg rulesGroup) string {
	ln := strings.ToLower(name)

	for _, r := range rg.exceptions {
		if r.gender == Androgynous || r.gender == g {
			for _, t := range r.test {
				if t == ln && len(r.mods) > int(c) {
					return applyRule(name, r.mods[c])
				}
			}
		}
	}

	return ""
}

func findInRules(name string, g Gender, c Case, rg rulesGroup) string {
	for _, r := range rg.suffixes {
		if r.gender == Androgynous || r.gender == g {
			for _, t := range r.test {
				if strings.HasSuffix(name, t) && len(r.mods) > int(c) {
					return applyRule(name, r.mods[c])
				}
			}
		}
	}

	return name
}

func applyRule(name, mod string) string {
	if mod == "." {
		return name
	}

	if i := strings.LastIndex(mod, "-"); i >= 0 {
		l := utf8.RuneCountInString(name) - i - 1
		return string([]rune(name)[:l]) + mod[i+1:]
	}

	return name + mod
}
