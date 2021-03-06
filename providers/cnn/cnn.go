package cnn

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("cnn", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for CNN.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.cnn.com/search/?text=%s", url.QueryEscape(q))
}
