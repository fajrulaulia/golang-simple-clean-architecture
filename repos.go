package main

type EvenOddInterface interface {
	ParsePlatRawToString(plat string) int
	PlatTypeCheck(numb int) string
}
