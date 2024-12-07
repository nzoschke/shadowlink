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

func TestOG(t *testing.T) {
	type test struct {
		in    string
		title string
		url   string
	}
	tests := []test{
		// bandcamp
		{in: "https://jexopolis.bandcamp.com/", title: "Jex Opolis", url: "https://jexopolis.bandcamp.com"},
		{in: "https://jexopolis.bandcamp.com/album/apocalypse-2", title: "Apocalypse, by Jex Opolis", url: "https://jexopolis.bandcamp.com/album/apocalypse-2"},
		{in: "https://jexopolis.bandcamp.com/track/apocalypse-at-the-acropolis", title: "Apocalypse (At the Acropolis), by Jex Opolis", url: "https://jexopolis.bandcamp.com/track/apocalypse-at-the-acropolis"},

		// soundcloud
		{in: "https://soundcloud.com/avalonemerson/", title: "Avalon Emerson", url: "https://soundcloud.com/avalonemerson"},
		{in: "https://soundcloud.com/avalonemerson?utm_source=clipboard&utm_medium=text&utm_campaign=social_sharing", title: "Avalon Emerson", url: "https://soundcloud.com/avalonemerson"},
		{in: "https://on.soundcloud.com/Hmr5s9pu8ppM6zbF8", title: "Avalon Emerson", url: "https://soundcloud.com/avalonemerson"},
		{in: "https://soundcloud.com/avalonemerson/sets/dj-kicks-ep-2", title: "DJ-Kicks EP", url: "https://soundcloud.com/avalonemerson/sets/dj-kicks-ep-2"},
		{in: "https://on.soundcloud.com/F5krqbb1LZn5nVsr7", title: "DJ-Kicks EP", url: "https://soundcloud.com/avalonemerson/sets/dj-kicks-ep-2"},
		{in: "https://soundcloud.com/avalonemerson/avalon-emerson-poodle-power-dj?in=avalonemerson/sets/dj-kicks-ep-2", title: "Poodle Power (DJ-Kicks)", url: "https://soundcloud.com/avalonemerson/avalon-emerson-poodle-power-dj"},

		// spotify
		{in: "https://open.spotify.com/artist/73A3bLnfnz5BoQjb4gNCga?si=eBY8A942So-FEYrkT7I9vA", title: "BICEP", url: "https://open.spotify.com/artist/73A3bLnfnz5BoQjb4gNCga"},
		{in: "https://open.spotify.com/album/0EdtTRCl3J22AnWrNpH1w9?si=BzL76ZmATTaW-Gafs3HXeg", title: "Isles", url: "https://open.spotify.com/album/0EdtTRCl3J22AnWrNpH1w9"},
		{in: "https://open.spotify.com/track/0pORLCI6Ep1eyqHJXbUPKG?si=fa2666649cf14d8d", title: "X", url: "https://open.spotify.com/track/0pORLCI6Ep1eyqHJXbUPKG"},

		// tidal
		{in: "https://tidal.com/browse/artist/5023272?u", title: "Nathan Micay", url: "https://tidal.com/browse/artist/5023272"},
		{in: "https://tidal.com/browse/album/302217867?u", title: "Nathan Micay - Fangs (Club Mix)", url: "https://tidal.com/browse/album/302217867"},
		{in: "https://tidal.com/browse/track/302217868?u", title: "Nathan Micay - Fangs", url: "https://tidal.com/browse/track/302217868"},

		// youtube
		{in: "https://www.youtube.com/channel/UCvyAsTkXVoCQtFG-rGhwCKw", title: "The Chemical Brothers", url: "https://www.youtube.com/channel/UCvyAsTkXVoCQtFG-rGhwCKw"},
		{in: "https://youtube.com/@thechemicalbrothers?si=x33MN7XY9-QCp7W3", title: "The Chemical Brothers", url: "https://www.youtube.com/channel/UCvyAsTkXVoCQtFG-rGhwCKw"},
		{in: "https://youtube.com/playlist?list=PLOnz33WtAqQPR3GGZG3xvrDpbeCFSimz8&si=8CtK_DfX-whgQLqF", title: `"Born In The Echoes" - Remixes`, url: "http://www.youtube.com/playlist?list=PLOnz33WtAqQPR3GGZG3xvrDpbeCFSimz8"},
		{in: "https://www.youtube.com/watch?v=LO2RPDZkY88", title: "The Chemical Brothers - Go (Official Music Video)", url: "https://www.youtube.com/watch?v=LO2RPDZkY88"},
	}

	for _, test := range tests {
		info, err := extract.Info(test.in)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, test.title, info.Title)
		assert.Equal(t, test.url, info.URL)
	}
}
