package main

import (
	"fmt"
	"reflect"
)

//function types
type mapf func(interface{}, interface{}) interface{}

// func(value, memo) interface
type reducef func(interface{}, interface{}) interface{}
type filterf func(interface{}) interface{}


func main() {

	// define your map

	emails := make(map[string] string)
	// alternative of writing a map function

	// emails := map[string]string{"email": "bob@gmail.com"}

	// assign key value pairs

	emails["email"] = "bob@gmail.com"

	emails["email1"] = "sharon@gmail.com"

	emails["email2"] = "eduh@gmail.com"

	fmt.Println("Map:", emails)

	// delate an email
	delete(emails, "email")
	fmt.Println("new list array after deleting an item ", emails)

	filteredEmailsList := Filter(emails, func(val interface{}) interface{} {
		fmt.Println(val)
		return val
	})

	fmt.Println("FILTER", emails, filteredEmailsList)

}

func Map(in interface{}, fn mapf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		out[i] = fn(val.Index(i).Interface(), val.Index(i).Interface())
	}

	return out
}

// filter a map function

func Filter(in interface{}, fn filterf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()

		out = append(out, current)

	}

	return out
}

// reduce a map function
