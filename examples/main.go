package main

import (
	"context"
	"fmt"
	"os"

	"../nyttop"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	topClient := nyttop.New(apiKey)

	ctx := context.Background()
	articles, err := topClient.TopStories(ctx, nyttop.SectionMovies)
	if err != nil {
		panic(err)
	}

	toShow := 3
	total := len(articles)
	if total < toShow {
		toShow = total
	}

	fmt.Println("total articles: ", total)
	for _, a := range articles[:toShow] {
		fmt.Println(a.Title)
	}
}
