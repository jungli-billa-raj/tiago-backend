package main

import (
	"log"
	"tiago/internal/env"
	"tiago/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8880"),
	}

	store := store.NewStorage(nil)

	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
