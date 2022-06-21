package main

import (
	"fmt"

	route2 "github.com/LeonardoGregoriocs/code_delivery/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "2",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()

	for i, j := range stringjson {
		fmt.Println(i, j)
	}
}
