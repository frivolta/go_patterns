package builder_ex_2

// SiteMapBuilder builder interface
type SiteMapBuilder interface {
	buildHeader() *SiteMapBuilder
	buildFooter() *SiteMapBuilder
	buildPage(url string) *SiteMapBuilder
	buildBody() *SiteMapBuilder
	String() string
}
