package main

import (
	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os/exec"
	"regexp"
	"time"
)

// Only enqueue the root and paths beginning with an "a"
var (
	rxOk = regexp.MustCompile(`http://duckduckgo\.com(/a.*)?$`)
)

func surfing(tag string) {
	p("surfing")
	script, err := exec.LookPath("youtube-dl")
	if err != nil {
		p("no youtube-dl, please install this python script")
	}
	scriptexe := exec.Command("bash -c cd music && " + script + " -x " + tag)
	scriptexe.Start()
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

func ExampleCrawl() {
	// Set custom options
	opts := gocrawl.NewOptions(new(ExampleExtender))
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll

	// Play nice with ddgo when running the test!
	opts.MaxVisits = 2

	// Create crawler and start at root of duckduckgo
	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("https://duckduckgo.com/")

	// Remove "x" before Output: to activate the example (will run on go test)

	// xOutput: voluntarily fail to see log output
}

func getTag() string {
	return "-r0NfT1e4DM"
}
