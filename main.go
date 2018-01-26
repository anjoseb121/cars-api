package main

import (
	"log"

	"github.com/anjoseb121/cars-api/api"
)

func main() {
	a := api.New(":9111")
	log.Fatal(a.Start())
}
