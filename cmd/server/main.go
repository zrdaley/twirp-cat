package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	
	"github.com/twitchtv/twirp"
	"github.com/zrdaley/twirp-cat/pkg/cat"
)

type randomCat struct{}

func (h *randomCat) MakeCat(ctx context.Context, catinput *cat.CatInput) (*cat.Cat, error) {
	if catinput.Name == "" {
		return nil, twirp.InvalidArgumentError("Name", "A cat cannot be unnamed!")
	}
	coats := []string{"white", "black", "tabby", "calico"}
	treats := []string{"salmon", "chicken", "liver"}
	return &cat.Cat{
		Name:  catinput.Name,
		Coat: coats[rand.Intn(len(coats))],
		FavoriteTreat:  treats[rand.Intn(len(treats))],
	}, nil
}

func main() {
	server := cat.NewCatGeneratorServer(&randomCat{})
	log.Print("Starting Cat Server ... >.<")
	
	log.Fatal(http.ListenAndServe(":8080", server))
}