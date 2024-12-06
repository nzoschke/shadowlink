package extract

import (
	"net/http"

	"github.com/dyatlov/go-htmlinfo/htmlinfo"
	"github.com/dyatlov/go-oembed/oembed"
	"golang.org/x/xerrors"
	"mvdan.cc/xurls/v2"
)

var strict = xurls.Strict()

func Extract(s string) []string {
	return strict.FindAllString(s, -1)
}

func Info(u string) (oembed.Info, error) {
	resp, err := http.Get(u)
	if err != nil {
		return oembed.Info{}, xerrors.Errorf(": %w", err)
	}
	defer resp.Body.Close()

	info := htmlinfo.NewHTMLInfo()
	err = info.Parse(resp.Body, &u, nil)
	if err != nil {
		return oembed.Info{}, xerrors.Errorf(": %w", err)
	}

	return *info.GenerateOembedFor(u), nil
}

func InfoLite(i oembed.Info) oembed.Info {
	i.AuthorURL = ""
	i.HTML = ""
	i.ProviderURL = ""
	i.ThumbnailURL = ""
	i.URL = ""
	return i
}
