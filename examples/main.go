package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tainacleal/nyt-go/nytbooks"
	"github.com/tainacleal/nyt-go/nyttop"
)

func main() {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	// NYT Top Stories
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

	fmt.Println("Top stories:")
	for _, a := range articles[:toShow] {
		fmt.Println(a.Title)
	}

	// NYT Best Sellers
	nytBooksClient := nytbooks.New(apiKey)
	ctx = context.Background()

	// fetch all available list options
	bestSellerOptions, err := nytBooksClient.BestSellersListsOptions(ctx)
	if err != nil {
		panic(err)
	}

	if len(bestSellerOptions.Lists) == 0 {
		fmt.Println("no lists returned")
		os.Exit(0)
	}

	// fetch latest best sellers for the first list
	bestSellers, err := nytBooksClient.LatestBestSellers(ctx, bestSellerOptions.Lists[0].ServiceName, 0)
	if err != nil {
		panic(err)
	}

	toShow = 3
	total = len(bestSellers.Books)
	if total < toShow {
		toShow = total
	}

	fmt.Println("Best Sellers:")
	for _, b := range bestSellers.Books[:toShow] {
		if len(b.Details) > 0 {
			fmt.Println(b.Details[0].Title)
		} else {
			fmt.Println("no detail available")
		}
	}

}
