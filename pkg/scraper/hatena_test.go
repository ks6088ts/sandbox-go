package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHatenaScraper(t *testing.T) {
	url := "https://b.hatena.ne.jp/hotentry/it"
	s := NewHatenaScraper(url, false)
	assert.Equal(t, url, s.url)

	// Comment in to test with actual network transaction
	// results, err := s.Scrape()
	// if err != nil {
	// 	t.Error(err)
	// }
	// assert.True(t, len(results) > 0)
}
