// Package gallery renders an HTML page that shows a set of images.
package gallery

import (
	"html/template"
	"io"
)

// Image is one picture on the page.
type Image struct {
	Name          string
	Width, Height int
}

var page = template.Must(template.New("gallery").Parse(`<!doctype html>
<title>Gallery</title>
<body style="background:#0b0d0f;color:#eee;font-family:monospace">
<h1>{{len .}} cards, built by Bazel</h1>
{{range .}}<p><a href="/{{.Name}}"><img src="/{{.Name}}" width="{{.Width}}" height="{{.Height}}" alt="{{.Name}}"></a></p>
{{end}}</body>
`))

// Render writes the page for images to w.
func Render(w io.Writer, images []Image) error {
	return page.Execute(w, images)
}
