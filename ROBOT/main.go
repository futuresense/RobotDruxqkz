package main

import (
	//"github.com/codegangsta/cli"
	//"github.com/futuresense/druxqkz/reddit""
	"github.com/benmanns/goworker"

	"bufio"
	"fmt"
	"os"
)

var (
	rp robotPreferences
	ra robotAction
	on bool = false
	p       = fmt.Println
)

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}
func helloWorker(queue string, args ...interface{}) error {
	fmt.Println("Hello, world!")
	return nil
}

func init() {
	goworker.Register("Hello", helloWorker)

	ra.coverArt()
	go func() {

		r := bufio.NewReader(os.Stdin)
		var inputString string
		//always watch for ESC
		for {

			inputBytes, _, _ := r.ReadLine()
			inputString = string(inputBytes)
			if (inputString == "x") || (inputString == "q") {
				os.Exit(0)
			}
		}

	}()
}
