package engadget

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("engadget", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for engadget.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.engadget.com/search/?search-terms=%s", url.QueryEscape(q))
}
