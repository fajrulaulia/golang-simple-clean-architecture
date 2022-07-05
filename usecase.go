package main

type MainService struct {
	svc EvenOddInterface
}

func (c *MainService) Validate(bl string) string {
	num := c.svc.ParsePlatRawToString(bl)
	result := c.svc.PlatTypeCheck(num)
	if result == "GENAP" {
		return "Plat anda Genap"
	}
	return "Plat anda Ganjil"
}
