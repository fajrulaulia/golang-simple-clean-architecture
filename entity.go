package main

import (
	"regexp"
	"strconv"
)

type EvenOddStruct struct{}

func (c EvenOddStruct) ParsePlatRawToString(plat string) int {
	regex := regexp.MustCompile("[0-9]+")
	i, err := strconv.Atoi(regex.FindAllString(plat, -1)[0])
	if err != nil {
		panic(err)
	}

	return int(i)
}

func (c EvenOddStruct) PlatTypeCheck(numb int) string {
	if numb%2 == 0 {
		return "GENAP"
	}
	return "GANJIL"
}
