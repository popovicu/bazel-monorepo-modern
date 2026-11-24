package gallery

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRender(t *testing.T) {
	var b strings.Builder
	if err := Render(&b, []Image{{Name: "foo.png", Width: 600, Height: 315}}); err != nil {
		t.Fatal(err)
	}

	var imgs []string
	for _, line := range strings.Split(b.String(), "\n") {
		if strings.Contains(line, "<img") {
			imgs = append(imgs, line)
		}
	}
	want := []string{
		`<p><a href="/foo.png"><img src="/foo.png" width="600" height="315" alt="foo.png"></a></p>`,
	}
	if diff := cmp.Diff(want, imgs); diff != "" {
		t.Errorf("Render() images mismatch (-want +got):\n%s", diff)
	}
}
