// Package reddit implements a basic client for the Reddit API
package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

//	"os"
)

// Item describes a Reddit item
type Item struct {
	Title    string
	URL      string
	Author   string
	Comments int `json:"num_comments"`
}

func (i Item) String() string {
	com := ""

	switch i.Comments {
	case 0:
	// nothing
	case 1:
		com = " (1 comment) "
	default:
		com = fmt.Sprintf(" (%d comments) ", i.Comments)
	}

	return fmt.Sprintf("%s: %q\n%s %s\n", i.Author, i.Title, i.URL, com)
}

type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func Init(s string) {
	var subreddit = s

	/*
		if len(os.Args) > 1 {
			fmt.Printf("%s is the subreddit requested.\n\n\n", os.Args[1])
			subreddit = os.Args[1]
		} else {
			subreddit = "nsfw"
		}*/

	items, err := Get(subreddit)
	if err != nil {
		log.Fatal(err)
	}

	var number_choice int
	var ever = true
	for ever {
		fmt.Print("number?")
		fmt.Scan(&number_choice)
		fmt.Printf("choice is: %v \n", number_choice)
		fmt.Printf("format %v", items[number_choice])
		if number_choice == 0 {
			ever = false
		}
	}

}

// Get fetches the most recent items posted to the subreddit
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://www.reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))

	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
	
}
