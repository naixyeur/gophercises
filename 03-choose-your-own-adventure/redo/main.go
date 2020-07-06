package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sandbox/gophercises/03-choose-your-own-adventure/redo/cyoa"
)

func main() {
	port := flag.Int("port", 8000, "port to start CYOA story")
	filename := flag.String("file", "gopher.json", "file that use in CYOA story")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(templateString))
	h := cyoa.NewHandler(story, cyoa.TemplateOption(tpl), cyoa.PathFnOption(pathFn))

	mux := http.NewServeMux()
	// mux.Handle("", h)
	mux.Handle("/", h)
	mux.Handle("/story", h)
	mux.Handle("/story/", h)
	fmt.Printf("start CYOA story at port: %d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathFn(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" || path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[7:]
}

var templateString = `
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
						<a href="/story/{{.Chapter}}">{{.Text}}</a>
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
