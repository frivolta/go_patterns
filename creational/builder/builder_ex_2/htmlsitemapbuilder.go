package builder_ex_2

import (
	"bytes"
	"fmt"
)

// HTMLSiteMapBuilder implements SiteMapBuilder
type HTMLSiteMapBuilder struct {
	header, footer, body string
	urls                 []string
}

func (h *HTMLSiteMapBuilder) buildHeader() *HTMLSiteMapBuilder {
	h.header = "<ul class=\"sitemap\">\n"
	return h
}

func (h *HTMLSiteMapBuilder) buildFooter() *HTMLSiteMapBuilder {
	h.footer = "</ul>\n"
	return h
}

func (h *HTMLSiteMapBuilder) buildPage(url string) *HTMLSiteMapBuilder {
	h.urls = append(h.urls, url)
	return h
}

func (h *HTMLSiteMapBuilder) buildBody() *HTMLSiteMapBuilder {
	var body bytes.Buffer
	for _, u := range h.urls {
		body.WriteString(fmt.Sprintf("\"<url>\\n<loc>%s</loc>\\n</url>\"", u))
	}
	h.body = body.String()
	return h
}

func (h *HTMLSiteMapBuilder) String() string {
	var sm bytes.Buffer
	sm.WriteString(h.header)
	sm.WriteString(h.body)
	sm.WriteString(h.footer)
	return sm.String()
}

// BuildSiteMap
// implementation of the SiteMapBuilder interface
func BuildSiteMap() {
	h := HTMLSiteMapBuilder{}
	h.buildHeader().buildFooter().buildPage("https://www.rivoltafilippo.com").buildBody()
	fmt.Println(h.String())
}
