package main

import (
	"fmt"
	"reflect"
)

func Main() {
	for _, v := range []any{"hi", 42, func() {}, struct{}{}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

	a := SlicesIndex([]string{"12"}, "12")

	fmt.Println("\nCOUNT", a)
}

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}
