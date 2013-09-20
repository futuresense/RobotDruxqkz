package robotdruxqkz

import (
	"fmt"
	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
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

type SubscriptionRequest chan 
	
	var TickerSubscriptionChan chan Request
}

func Ticker() {
	go func() {
		tick := 0
		var subscriptions []SubscriptionRequest
		for {
			select {
				case request := <- TickerSubscriptionChan;
					subscriptions = append(subscriptions, request)
				default:
					for _, subscription := range subscriptions {
						subscription <- tick
					}
					tick++
			}
		}
	}()
}

type URLBuild struct {
	Full_url string
}

type SearchResults []*SearchResult

func (s SearchResults) Len() int      { return len(s) }
func (s SearchResults) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

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

func Surfer() {
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

func Crawler() {
}

func getArtistILike() string {
	var a []string
	if isUser && isOnline {
		a = []string{}
	} else if isUser && !isOnline {
		//hack, readin file here
		a = []string{"shlohmo", "bjork", "sadf"}
	} else {
		a = []string{}
	}
	return a[0]
}
