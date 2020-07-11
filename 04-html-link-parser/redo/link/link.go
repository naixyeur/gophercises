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
		panic(err)
	}

	var ret []Link
	nodes := getNodes(doc)
	// dfs(n)

	for _, n := range nodes {
		ret = append(ret, buildLink(n))
	}

	for _, r := range ret {
		fmt.Printf("%+v \n\n", r)
		_ = r
	}

	return ret, nil

}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = getText(n)

	return ret
}

// func getText(n *html.Node) string {
//   var ret string
//   if n.Type == html.TextNode {
//     ret = n.Data
//   }
//   for c := n.FirstChild; c != nil; c = c.NextSibling {
//     ret += getText(c) + " "
//   }
//   return strings.Join(strings.Fields(ret), " ")
// }

func getText(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func getNodes(n *html.Node) []*html.Node {
	var ret []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		ret = []*html.Node{n}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, getNodes(c)...)
	}
	return ret
}

// func dfs(n *html.Node) {
//   fmt.Printf("type: %+v  data: %+v \n\n", n.Type, n.Data)
//   for c := n.FirstChild; c != nil; c = c.NextSibling {
//     dfs(c)
//     fmt.Printf("next sibling \n\n")
//   }
//   fmt.Printf("up \n\n")
// }
