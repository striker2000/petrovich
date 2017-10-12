![Petrovich](petrovich.png)

Petrovich is the library which inflects Russian names to given grammatical case.

This is the Go port of https://github.com/petrovich.

[![GoDoc](https://godoc.org/github.com/striker2000/petrovich?status.svg)](https://godoc.org/github.com/striker2000/petrovich)
[![Build Status](https://travis-ci.org/striker2000/petrovich.svg?branch=master)](https://travis-ci.org/striker2000/petrovich)
![cover.run go](https://cover.run/go/github.com/striker2000/petrovich.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/striker2000/petrovich)](https://goreportcard.com/report/github.com/striker2000/petrovich)

## Installation

```
go get github.com/striker2000/petrovich
```

## Usage

All functions takes three arguments: name in nominative case, gender and grammatical case.

```go
import "github.com/striker2000/petrovich"

n := petrovich.FirstName("Кузьма", petrovich.Male, petrovich.Genitive)
fmt.Print(n) // "Кузьмы"

n = petrovich.MiddleName("Сергеевич", petrovich.Male, petrovich.Instrumental)
fmt.Print(n) // "Сергеевичем"

n = petrovich.LastName("Петров-Водкин", petrovich.Male, petrovich.Prepositional)
fmt.Print(n) // "Петрове-Водкине"
```

Valid values for gender are `petrovich.Androgynous`, `petrovich.Male` and `petrovich.Female`.

Full list of grammatical cases is in the table below.

| Case                      | Case (in Russian) | Question (in Russian)  |
|---------------------------|-------------------|------------------------|
| `petrovich.Nominative`    | Именительный      | Кто?                   |
| `petrovich.Genitive`      | Родительный       | Кого?                  |
| `petrovich.Dative`        | Дательный         | Кому?                  |
| `petrovich.Accusative`    | Винительный       | Кого?                  |
| `petrovich.Instrumental`  | Творительный      | Кем?                   |
| `petrovich.Prepositional` | Предложный        | О ком?                 |
