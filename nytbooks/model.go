package nytbooks

const (
	bestSellersListNameURL   = "https://api.nytimes.com/svc/books/v3/lists/names.json"
	bestSellersListLatestURL = "https://api.nytimes.com/svc/books/v3/lists.json"
)

type ListOption struct {
	Name                   string `json:"list_name"`
	DisplayName            string `json:"display_name"`
	ServiceName            string `json:"list_name_encoded"`
	NewestPublishedDateStr string `json:"newest_published_date"`
}

type BestSellersListOptionsResponse struct {
	Status     string       `json:"status"`
	Copyright  string       `json:"copyright"`
	NumResults int          `json:"num_results"`
	Lists      []ListOption `json:"results"`
}

type BestSellersListResponse struct {
	Status     string `json:"status"`
	Copyright  string `json:"copyright"`
	NumResults int    `json:"num_results"`
	Books      []Book `json:"results"`
}

type Book struct {
	ListName           string        `json:"list_name"`
	ListDisplayName    string        `json:"display_name"`
	BestSellersDateStr string        `json:"bestsellers_date"`
	PublishedDateStr   string        `json:"published_date"`
	Rank               int           `json:"rank"`
	RankLasWeek        int           `json:"rank_last_week"`
	WeeksOnList        int           `json:"weeks_on_list"`
	ISBNS              []ISBN        `json:"isbns"`
	Details            []BookDetails `json:"book_details"`
	Reviews            []BookReview  `json:"reviews"`
}

type ISBN struct {
	ISBN10 string `json:"isbn10"`
	ISBN13 string `json:"isbn13"`
}

type BookDetails struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Contributor string `json:"contributor"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
}

type BookReview struct {
	ReviewLink       string `json:"book_review_link"`
	SundayReviewLink string `json:"sunday_review_link"`
}
