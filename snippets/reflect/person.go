package main

import (
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func ShowTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag
		println(tag)
	}
}

func show(i interface{}) {
	switch i.(type) {
	case *Person:
		t := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		tag := t.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
		age := v.Elem().Field(1).Int()
		println(tag)
		println(name)
		println(age)
	}
}

func main() {
	p := new(Person)
	p.name = "test"
	p.age = 18
	ShowTag(p)
	show(p)
}
