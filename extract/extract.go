package extract

import (
	"net/http"
	"slices"
	"sort"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/dyatlov/go-htmlinfo/htmlinfo"
	"github.com/dyatlov/go-oembed/oembed"
	"golang.org/x/xerrors"
	"mvdan.cc/xurls/v2"
)

var (
	domains = []string{"bandcamp.com", "soundcloud.com", "spotify.com", "tidal.com", "youtube.com"}
	strict  = xurls.Strict()
)

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

	// use canonical og:url if present
	i := *info.GenerateOembedFor(u)
	if info.OGInfo.URL != "" {
		i.URL = info.OGInfo.URL
	}

	return i, nil
}

func MediaInfos(content string) ([]oembed.Info, error) {
	links := Extract(content)
	infos := map[string]oembed.Info{}
	for _, link := range links {
		info, err := Info(link)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		if !slices.Contains(domains, domainutil.Domain(info.URL)) {
			continue
		}

		infos[info.URL] = info
	}

	var result []oembed.Info
	var urls []string
	for url := range infos {
		urls = append(urls, url)
	}
	sort.Strings(urls)

	for _, url := range urls {
		result = append(result, infos[url])
	}

	return result, nil
}
