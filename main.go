//go:generate go run ./cli/gen-meta/main.go -output meta_gen.go ./package.json
package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	body, err := HTML(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
