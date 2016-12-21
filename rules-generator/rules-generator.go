package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type rules struct {
	FirstName  rulesGroup
	MiddleName rulesGroup
	LastName   rulesGroup
}

type rulesGroup struct {
	Exceptions rulesSet
	Suffixes   rulesSet
}

type rulesSet []rule

type rule struct {
	Gender string
	Test   []string
	Mods   []string
}

func main() {
	b, err := ioutil.ReadFile("rules/rules.json")
	if err != nil {
		panic(err)
	}

	var r rules

	err = json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}

	o, err := os.Create("rules_generated.go")
	if err != nil {
		panic(err)
	}

	defer o.Close()

	fmt.Fprint(o, "// DO NOT EDIT!\n// Code generated from rules/rules.json\n\npackage petrovich\n\n")

	fmt.Fprint(o, "var allRules = rules{\n\tfirstName: rulesGroup{\n")
	printRulesGroup(o, r.FirstName)

	fmt.Fprint(o, "\t},\n\tmiddleName: rulesGroup{\n")
	printRulesGroup(o, r.MiddleName)

	fmt.Fprint(o, "\t},\n\tlastName: rulesGroup{\n")
	printRulesGroup(o, r.LastName)

	fmt.Fprint(o, "\t},\n}\n")
}

func printRulesGroup(o io.Writer, g rulesGroup) {
	fmt.Fprint(o, "\t\texceptions: []rule{\n")
	printRulesSet(o, g.Exceptions)
	fmt.Fprint(o, "\t\t},\n")

	fmt.Fprint(o, "\t\tsuffixes: []rule{\n")
	printRulesSet(o, g.Suffixes)
	fmt.Fprint(o, "\t\t},\n")
}

func printRulesSet(o io.Writer, s rulesSet) {
	for _, r := range s {
		var g string
		switch r.Gender {
		case "male":
			g = "Male"
		case "female":
			g = "Female"
		default:
			g = "Androgynous"
		}

		fmt.Fprintf(o, "\t\t\t{\n\t\t\t\tgender: %s,\n\t\t\t\ttest: []string{\n", g)
		for _, t := range r.Test {
			fmt.Fprintf(o, "\t\t\t\t\t\"%s\",\n", t)
		}
		fmt.Fprint(o, "\t\t\t\t},\n\t\t\t\tmods: []string{\n")
		for _, t := range r.Mods {
			fmt.Fprintf(o, "\t\t\t\t\t\"%s\",\n", t)
		}
		fmt.Fprint(o, "\t\t\t\t},\n\t\t\t},\n")
	}
}
