package main

import (
	"time"

	"github.com/zic20/pokedex/internal"
)

func main() {
	client := internal.NewPokedex(5*time.Second, 5*time.Minute)
	runrepl(&client)
}
