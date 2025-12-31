package main

import (
	"testing"
)

func TestShortenAndResolve(t *testing.T) {
	url := "http://x.com"

	url_shortener := NewURLShortener()

	hash, err := url_shortener.Shorten(url)

	if err != nil {
		t.Fatal("Failed to shorten URL")
	}

	resolved_url, err := url_shortener.Resolve(hash)

	if err != nil {
		t.Fatal("Failed to resolve hash")
	}

	if resolved_url != url {
		t.Fatal("URL match failed!")
	}

}
