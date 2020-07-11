package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sandbox/gophercises/04-html-link-parser/solution/link"
	"strings"
)

/*
	1. GET the webpage
	2. parse all the link on the page
	3. build proper url with our links
	4. filter out any links w/ a diff domain
	5. find all pages (BFS)
	6. print out XML
*/

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"value"`
}

type urlSet struct {
	Url   []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {

	urlFlag := flag.String("url", "http://gophercises.com/", "the url that you want to build a sitemap for")
	maxDepthFlag := flag.Int("depth", 3, "traversal depth")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepthFlag)

	toXml := urlSet{
		Url:   make([]loc, len(pages)),
		Xmlns: xmlns,
	}

	for i, page := range pages {
		toXml.Url[i].Value = page
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	fmt.Print(xml.Header)
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

	fmt.Println()
	// for _, h := range pages {
	//   fmt.Println(h)
	// }

}

func bfs(urlStr string, maxDepth int) []string {
	visited := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, map[string]struct{}{}

		if len(q) == 0 {
			break
		}

		for url, _ := range q {
			if _, ok := visited[url]; ok {
				continue
			}
			visited[url] = struct{}{}
			for _, l := range get(url) {
				nq[l] = struct{}{}
			}

		}
	}

	ret := make([]string, 0, len(visited))

	for l, _ := range visited {
		ret = append(ret, l)
	}

	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		// return struct{}{}
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(getHrefs(resp.Body, base), withPrefix(base))
	// return filter(getHrefs(resp.Body, base), withPrefix("https://twitter"))
}

func getHrefs(r io.Reader, base string) []string {
	var ret []string
	links, _ := link.Parse(r)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, l := range links {
		switch {
		case keepFn(l):
			ret = append(ret, l)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(str string) bool {
		return strings.HasPrefix(str, pfx)
	}
}
