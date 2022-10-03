package main

import (
	"log"

	"github.com/luispfcanales/api-phrases/cmd/http/bootstrap"
)

func main() {
	if err := bootstrap.RunHTTP(); err != nil {
		log.Fatal(err)
	}
}
