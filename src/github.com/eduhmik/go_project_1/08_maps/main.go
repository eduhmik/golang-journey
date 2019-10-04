package main

import (
	"fmt"
	"reflect"
)

//function types
type mapf func(interface{}) interface{}

// func(value, memo) interface
type reducef func(interface{}, interface{}) interface{}
type filterf func(interface{}) bool

func main() {
	a := []int{1, 2, 3, 4}

	//Multiply everything by 2
	b := Map(a, func(val interface{}) interface{} {
		return val.(int) * 2
	})

	//Shoud be [2,4,6,8]
	fmt.Println("MAP:", a, b)

	//Summation
	c := Reduce(b, 0, func(val interface{}, memo interface{}) interface{} {
		return memo.(int) + val.(int)
	})

	//Should be 20
	fmt.Println("REDUCE:", b, c)

	//Check if the number is divisble by 4
	d := Filter(b, func(val interface{}) bool {
		return val.(int)%4 == 0
	})

	//Should be [4,8]
	fmt.Println("FILTER:", b, d)
}

//Map(slice, func)
func Map(in interface{}, fn mapf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		out[i] = fn(val.Index(i).Interface())
	}

	return out
}

//Reduce(slice, starting value, func)
func Reduce(in interface{}, memo interface{}, fn reducef) interface{} {
	val := reflect.ValueOf(in)

	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}

	return memo
}

//Filter(slice, predicate func)
func Filter(in interface{}, fn filterf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()

		if fn(current) {
			out = append(out, current)
		}
	}

	return out
}