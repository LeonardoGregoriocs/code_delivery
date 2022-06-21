package main

import (
	"fmt"
	"log"

	route2 "github.com/LeonardoGregoriocs/code_delivery/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

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
