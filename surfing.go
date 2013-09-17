package main

import (
	//	"fmt"
	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"syscall"
	"time"
)

// Only enqueue the root and paths beginning with an "a"
var (
	isUser   = true
	isOnline = false
	rxOk     = regexp.MustCompile(`http://duckduckgo\.com(/a.*)?$`)
	sr       = SearchResult{}
)

type SurfingCrawlerStuff struct {
	Tag      string
	Priority int16
}

type SearchResult struct {
	Title []string
	Tag   string
}
type SurferSettings struct {
}

type URLBuild struct {
	Full_url string
}

type SearchResults []*SearchResult

func (s SearchResults) Len() int      { return len(s) }
func (s SearchResults) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func surfing() {
	p("surfing")
	binary, err := exec.LookPath("youtube-dl")
	if err != nil {
		p("no youtube-dl, please install this python script")
	}

	env := os.Environ()
	args := []string{"youtube-dl", "-x", sr.Tag}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	for _, tag := range sr.Tag {
		p(tag)
	}

}

////////////////////////////////////////////////////////////////////////////////////
// WEB CRAWLER
////////////////////////////////////////////////////////////////////////////////////
// Create the Extender implementation, based on the gocrawl-provided DefaultExtender,
// because we don't want/need to override all methods.
type ExampleExtender struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visit() and Filter()
}

// Override Visit for our need.
func (this *ExampleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// Use the goquery document or res.Body to manipulate the data
	// ...

	// Return nil and true - let gocrawl find the links
	return nil, true
}

// Override Filter for our need.
func (this *ExampleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited && rxOk.MatchString(ctx.NormalizedURL().String())
}

func SurfingCrawler() {

	// Set custom options
	opts := gocrawl.NewOptions(new(ExampleExtender))
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll

	// Play nice with ddgo when running the test!
	opts.MaxVisits = 2

	// Create crawler and start at root of duckduckgo
	c := gocrawl.NewCrawlerWithOptions(opts)

	c.Run(urlAddress())

	// Remove "x" before Output: to activate the example (will run on go test)

	// xOutput: voluntarily fail to see log output
}
func urlAddress() string {
	URLHome := "http://www.google.com/"
	var URLFragment = []string{"/search?q=same"}
	var URLFragments string
	for _, v := range URLFragment {
		URLFragments += v
	}
	var fullURL = URLFragments + URLHome
	return fullURL
}
func searchYoutube() {
}

/*
func getArtistILike() string {
	var a []string
	if isUser && isOnline {
		a := []string{}
	} else if isUser && !isOnline {
		//hack, readin file here
		a := []string{"shlohmo", "bjork", "sadf"}
	} else {
		a := []string{}
	}
	return a[0]
}*/
