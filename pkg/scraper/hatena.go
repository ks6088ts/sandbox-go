package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

// HatenaScraper ...
type HatenaScraper struct {
	url       string
	collector *colly.Collector
}

// NewHatenaScraper creates a HatenaScraper
func NewHatenaScraper(url string, debug bool) *HatenaScraper {
	allowedDomains := colly.AllowedDomains("b.hatena.ne.jp")
	if debug {
		return &HatenaScraper{
			url: url,
			collector: colly.NewCollector(
				allowedDomains,
				colly.CacheDir("outputs/cache"), // for development
			),
		}
	}

	return &HatenaScraper{
		url: url,
		collector: colly.NewCollector(
			allowedDomains,
		),
	}
}

// Scrape ...
func (s *HatenaScraper) Scrape() ([]string, error) {
	results := []string{}

	// On every a element which has href attribute call callback
	s.collector.OnHTML(".entrylist-contents-main", func(e *colly.HTMLElement) {
		results = append(results, fmt.Sprintf("[%v: %v](%v)",
			e.ChildText(".entrylist-contents-users > a > span"),   // bookmark
			e.ChildAttr(".entrylist-contents-title > a", "title"), // title
			e.ChildAttr(".entrylist-contents-title > a", "href"),  // link
		))
	})

	// Before making a request print "Visiting ..."
	s.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("# ", r.URL.String())
	})

	// Start scraping on url
	err := s.collector.Visit(s.url)
	if err != nil {
		return nil, err
	}

	return results, nil
}
