package main

import (
	"log"

	"github.com/ssr0016/gobank/api"
	"github.com/ssr0016/gobank/model"
)

func main() {
	store, err := model.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
