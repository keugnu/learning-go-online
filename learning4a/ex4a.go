package main

import (
	"fmt"
	"strings"
)

type String string
type Int int
type Strings []String
type Ints []Int

type UpperCaser interface {
	UpperCase() String
}

func (s String) UpperCase() String {
	return String(strings.ToUpper(string(s)))
}

type Exponentiater interface {
	Exponentiate() Int
}

func (i Int) Exponentiate() Int {
	return Int(int(i) * int(i))
}

type Mapper interface {
	Map(func(interface{}) interface{}) []interface{}
}

func (i Ints) Map(f func(Int) Int) Ints {
	j := make(Ints, len(i))
	for k, v := range i {
		j[k] = f(Int(v))
	}
	return j
}

func (s Strings) Map(f func(String) String) Strings {
	j := make(Strings, len(s))
	for k, v := range s {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := Ints{1, 4, 9, 0}
	fmt.Printf("%v", m.Map(Int.Exponentiate))
	l := Strings{"these", "are", "strings"}
	fmt.Printf("%s", l.Map(String.UpperCase))
}
