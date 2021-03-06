package bandcamp

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("bandcamp", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Bandcamp.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://bandcamp.com/search?q=%s", url.QueryEscape(q))
}
