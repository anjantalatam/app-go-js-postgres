package main

import (
	"log"

	"github.com/anjantalatam/myPackageProgram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType

	myVar.TypeName = "Anjan"
	log.Println(myVar.TypeName)
}
