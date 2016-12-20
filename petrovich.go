package petrovich

//go:generate go-bindata -pkg petrovich rules/rules.json

import (
	"encoding/json"
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

var allRules rules

func init() {
	b := MustAsset("rules/rules.json")

	err := json.Unmarshal(b, &allRules)
	if err != nil {
		panic(err)
	}
}

func FirstName(name string, g Gender, c Case) string {
	return inflect(name, g, c, allRules.FirstName)
}

func MiddleName(name string, g Gender, c Case) string {
	return inflect(name, g, c, allRules.MiddleName)
}

func LastName(name string, g Gender, c Case) string {
	return inflect(name, g, c, allRules.LastName)
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

	for _, r := range rg.Exceptions {
		if rg := r.getGender(); rg == Androgynous || rg == g {
			for _, t := range r.Test {
				if t == ln && len(r.Mods) > int(c) {
					return applyRule(name, r.Mods[c])
				}
			}
		}
	}

	return ""
}

func findInRules(name string, g Gender, c Case, rg rulesGroup) string {
	for _, r := range rg.Suffixes {
		if rg := r.getGender(); rg == Androgynous || rg == g {
			for _, t := range r.Test {
				if strings.HasSuffix(name, t) && len(r.Mods) > int(c) {
					return applyRule(name, r.Mods[c])
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
