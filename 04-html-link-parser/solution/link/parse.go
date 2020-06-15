package link

import "io"

// Link represetns a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// parse will take in a HTML document and will returnd a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {

	return nil, nil
}

