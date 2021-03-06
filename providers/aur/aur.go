package aur

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("aur", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for the AUR.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://aur.archlinux.org/packages/?K=%s", url.QueryEscape(q))
}
