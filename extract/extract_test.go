package extract_test

import (
	"testing"

	"github.com/dyatlov/go-oembed/oembed"
	"github.com/nzoschke/shadowlink/extract"
	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	u := "https://www.youtube.com/watch?v=_erVOAbz420"
	info, err := extract.Info(u)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, oembed.Info{
		AuthorName:      "Documentary Films Trailer | Cosmos Documentaries",
		AuthorURL:       "https://www.youtube.com/@documentaryfilmstrailercos8331",
		Description:     "â�¶ Watch Free the full Documentary Series Online in Cosmos Documentaries Click Here: http://cosmos-documentaries.blogspot.gr/2014/03/cosmos-spacetime-odyssey-...",
		Height:          113,
		HTML:            "<iframe width=\"200\" height=\"113\" src=\"https://www.youtube.com/embed/_erVOAbz420?feature=oembed\" frameborder=\"0\" allow=\"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share\" referrerpolicy=\"strict-origin-when-cross-origin\" allowfullscreen title=\"Cosmos: A Spacetime Odyssey | Official Trailer\"></iframe>",
		ProviderName:    "YouTube",
		ProviderURL:     "https://www.youtube.com/",
		ThumbnailHeight: 360,
		ThumbnailURL:    "https://i.ytimg.com/vi/_erVOAbz420/hqdefault.jpg",
		ThumbnailWidth:  480,
		Title:           "Cosmos: A Spacetime Odyssey | Official Trailer",
		Type:            "video",
		URL:             "https://www.youtube.com/watch?v=_erVOAbz420",
		Width:           200,
	}, info)
}

func TestExtract(t *testing.T) {
	s := "this trailer https://www.youtube.com/watch?v=_erVOAbz420 is cool"
	urls := extract.Extract(s)
	assert.Equal(t, []string{
		"https://www.youtube.com/watch?v=_erVOAbz420",
	}, urls)
}
