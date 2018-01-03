package main

import (
	"fmt"

	"github.com/ByteArena/go-bytearena/client"
)

func main() {
	client, err := client.New()
	if err != nil {
		panic(err)
	}
	err = client.Start(&Agent{})
	if err != nil {
		panic(err)
	}
	fmt.Println("End")
}
