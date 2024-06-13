package main

import "github.com/PCPedroso/pos-fc-apis/configs"

func main() {
	c, _ := configs.LoadConfig(".")
	println(c.DBDriver)
}
