package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"
	"github.com/zrdaley/twirp-cat/pkg/cat"
)

func main() {
	client := cat.NewCatGeneratorJSONClient("http://localhost:8080", &http.Client{})

	var (
		randomCat *cat.Cat
		err error
	)
	for i := 0; i < 3; i++ {
		randomCat, err = client.MakeCat(context.Background(), &cat.CatInput{Name: "sparkles"})
		if err != nil {
			if twerr, ok := err.(twirp.Error); ok {
				if twerr.Meta("retryable") != "" {
					// Log the error and go again.
					log.Printf("got error %q, retrying", twerr)
					continue
				}
			}
			// This was some fatal error!
			log.Fatal(err)
		}
	}
	fmt.Printf("%+v", randomCat)
}