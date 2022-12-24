package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() {
	// Anonymous struct
	anotherDoctor := struct{name string}{name: "John"}
	fmt.Println(anotherDoctor)
	doc := anotherDoctor
	doc.name = "John Doe"
	fmt.Println(doc)

	aDoctor := Doctor{
		// don't need to add names
		number: 3,
		actorName: "Joe",
		companions: []string {
			"Jared", 
			"Steph",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[0])

	t := reflect.TypeOf(Doctor{})
	field, _ := t.FieldByName("actorName")
	fmt.Println(field.Tag)

	// Go uses embedding instead of inheritance, where it can take a struct type as a variable of another struct
}

// Summary
// Collections of disparate data types that describe a single concept
// Keyed by name fields
// Can create anonymous structs (short lived structs)
// They are value types
// Can use composition via embedding (instead of inheritance)
// Tags can be added to struct fields to describe fields