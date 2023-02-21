package main

import (
	"fmt"
	"reflect"
)

func main() {
	var pl = fmt.Println
	var vName = "Mario Haryzal"
	var vFirstName = "Mario"
	var vLastName = "Haryzal"
	age := 24
	phi := 3.14
	pl(vName, vFirstName, vLastName, age, phi)
	pl(phi, reflect.TypeOf(phi))
}
