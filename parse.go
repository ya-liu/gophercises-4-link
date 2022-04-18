package link

import "io"

// Link represents a link in an HTML document
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and will return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	var link []Link
	return link, nil
}
