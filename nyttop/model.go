package nyttop

import (
	"errors"
	"time"
)

type Section string

const (
	SectionArts         Section = "arts"
	SectionAutomobiles  Section = "automobile"
	SectionBooks        Section = "books"
	SectionBusiness     Section = "business"
	SectionFashion      Section = "fashion"
	SectionFood         Section = "food"
	SectionHealth       Section = "health"
	SectionHome         Section = "home" // homepage
	SectionInsider      Section = "insider"
	SectionMagazine     Section = "magazine"
	SectionMovies       Section = "movies"
	SectionNYregion     Section = "nyregion"
	SectionObituaries   Section = "obituaries"
	SectionOpinion      Section = "opinion"
	SectionPolitics     Section = "politics"
	SectionRealEstate   Section = "realestate"
	SectionScience      Section = "science"
	SectionSports       Section = "sports"
	SectionSundayReview Section = "sundayreview"
	SectionTechnology   Section = "technology"
	SectionTheater      Section = "theater"
	SectionTMagazine    Section = "t-magazine"
	SectionTravel       Section = "travel"
	SectionUpshot       Section = "upshot"
	SectionUS           Section = "us"
	SectionWorld        Section = "world"
)

var Sections = map[Section]struct{}{
	SectionArts:         {},
	SectionAutomobiles:  {},
	SectionBooks:        {},
	SectionBusiness:     {},
	SectionFashion:      {},
	SectionFood:         {},
	SectionHealth:       {},
	SectionHome:         {},
	SectionInsider:      {},
	SectionMagazine:     {},
	SectionMovies:       {},
	SectionNYregion:     {},
	SectionObituaries:   {},
	SectionOpinion:      {},
	SectionPolitics:     {},
	SectionRealEstate:   {},
	SectionScience:      {},
	SectionSports:       {},
	SectionSundayReview: {},
	SectionTechnology:   {},
	SectionTheater:      {},
	SectionTMagazine:    {},
	SectionTravel:       {},
	SectionUpshot:       {},
	SectionUS:           {},
	SectionWorld:        {},
}
var ErrInvalidSection = errors.New("invalid section")

const (
	baseURL string = "https://api.nytimes.com/svc/topstories/v2"
)

type TopStoriesResponse struct {
	Status      string    `json:"status"`
	Copyright   string    `json:"copyright"`
	Section     string    `json:"section"`
	LastUpdated time.Time `json:"last_updated"`
	NumResults  int       `json:"num_results"`
	Articles    []Article `json:"results"`
}

// Article represents a nyt article
// NOTE: only fields I care about for now
type Article struct {
	Section     string    `json:"section"`
	Subsection  string    `json:"subsection"`
	Title       string    `json:"title"`
	Abstract    string    `json:"abstract"`
	URL         string    `json:"url"`
	ShortURL    string    `json:"short_url"`
	Byline      string    `json:"byline"`
	UpdatedAt   time.Time `json:"updated_date"`
	CreatedAt   time.Time `json:"created_date"`
	PublishedAt time.Time `json:"published_date"`
}
