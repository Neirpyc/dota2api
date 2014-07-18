package main

import (
	"fmt"
	"github.com/l2x/dota2api"
)

func main() {
	dota2, err := dota2api.LoadConfig("../config.ini")

	if err != nil {
		fmt.Println(err)
	}

	steamId, err := dota2.ResolveVanityUrl("specode")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(steamId)

}
