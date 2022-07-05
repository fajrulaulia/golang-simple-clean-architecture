package main

import (
	"log"
)

func main() {
	mainServiceNew := MainService{EvenOddStruct{}}
	res := mainServiceNew.Validate("bl4318lq")
	log.Println(res)
}
