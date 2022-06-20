package main

import (
	"book-store-of-xambration-go/api"
	_ "book-store-of-xambration-go/api/controllers"
	"book-store-of-xambration-go/infrastructure"
	"log"
)

func main() {
	if err := infrastructure.Migrate(); err != nil {
		log.Fatal(err.Error())
	}
	infrastructure.RouterConfig(api.Register)
}
