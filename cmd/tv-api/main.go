package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tvapi "github.com/keitax/tv-api"
)

func main() {
	port := os.Getenv("PORT")
	app := tvapi.NewApplication()
	log.Printf("Listening on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), app); err != nil {
		log.Fatal(err)
	}
}
