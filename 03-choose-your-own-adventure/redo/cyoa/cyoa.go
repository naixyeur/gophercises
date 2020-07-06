package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func init() {
	defaultTemplate = template.Must(template.New("").Parse(defaultTemplateString))
}

var defaultTemplate *template.Template

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(*http.Request) string
}

type option func(*handler)

func TemplateOption(tmpl *template.Template) option {
	return func(h *handler) {
		h.t = tmpl
	}
}

func PathFnOption(pFn func(*http.Request) string) option {
	return func(h *handler) {
		h.pathFn = pFn
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Fatalf("%+v", err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "chapter not found", http.StatusNotFound)

}

func NewHandler(sty Story, opts ...option) http.Handler {
	h := handler{
		s:      sty,
		t:      defaultTemplate,
		pathFn: defaultPathFn,
	}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)
	return story, err
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

var defaultTemplateString = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>
			
			{{range .Paragraphs}}
				<p>{{.}}</p>
			{{end}}

			<ul>
				{{range .Options}}
					<li>
						<a href="/{{.Chapter}}">{{.Text}}</a>
					</li>
				{{end}}
			</ul>
		</section>
	</body>

	<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align: center;
			position: relative;
		}
		.page{
			width: 80%;
			max-width:500px;
			margin: auto;
			margin-top: 40px;
			margin-botton: 40px;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		}
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		li {
			padding-top: 10px;
		}
		a,
		a:visited{
			text-decoration: none;
			color: #6295b5;
		}
		a:active,
		a:hover{
			color: #7792a2;
		}
		p {
			text-indent: 1em;
		}
	</style>

</html>
`
