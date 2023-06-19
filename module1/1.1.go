package main

import (
	"fmt"
)

func main() {
	Question1()
}

func Question1() {
	a := "I"
	b := "am"
	c := "stupid"
	d := "and"
	e := "weak"
	array := []*string{&a, &b, &c, &d, &e}
	fmt.Printf("%v \n", array)
	UpdateArray(array)
	fmt.Printf("%v \n", array)
}
func UpdateArray(array []*string) []*string {
	for _, s := range array {
		switch *s {
		case "stupid":
			f := "smart"
			s = &f
		case "weak":
			g := "strong"
			s = &g
		}
	}
	return array
}
