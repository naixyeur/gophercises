package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var ret []Link
	nodes := linkNodes(doc)

	for _, node := range nodes {
		ret = append(ret, buildLink(node))
	}

	for _, r := range ret {
		fmt.Printf("%+v \n", r)
	}

	return nil, nil
}

func text(n *html.Node) string {

	// fmt.Printf("%+v \n\n", n)
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = text(n)
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// func dfs(n *html.Node) []*html.Node {

//   if n.FirstChild == nil {
//     return []*html.Node{n}
//   }

//   var ret []*html.Node

//   for c := n.FirstChild; c != nil; c = c.NextSibling {
//     ret = append(ret, dfs(c)...)
//   }
//   return ret
// }
