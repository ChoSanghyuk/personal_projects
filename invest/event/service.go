package event

type Scraper interface {
	Scrape() (string, error)
}
