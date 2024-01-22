package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ssr0016/gobank/api"
	"github.com/ssr0016/gobank/model"
	"github.com/ssr0016/gobank/types"
)

// Step 5
func seedAccount(store model.Storage, fname, lname, pw string) *types.Account {

	acc, err := types.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account =>", acc.Number)

	return acc
}

func seedAccounts(s model.Storage) {
	seedAccount(s, "sams", "recs", "123456")
} // end

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := model.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	// // seed stuff
	// seedAccounts(store)

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
