package say

import (
	"fmt"
	"github.com/futuresense/RobotDruxqkz/reddit"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type OS struct {
	x string
}

var (
	say *os.Process
	awe string
)

func Init(s string) {

	items, err := reddit.Get(s)

	if err != nil {
		log.Fatal(err)
	}
	n := len(items)
	//a := make([]string, n)

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(items[randInt(1, n)].Title)
}

func sayCheck() string {
	path, err := exec.LookPath("say")
	if err != nil {
		log.Fatal("Must not be on a Mac :(")
	}
	return path
}

func Spaz(inputString string) *os.Process {
	path := sayCheck()
	sayExe := exec.Command(path, inputString)
	err := sayExe.Start()
	if err != nil {
		log.Fatal(err)
	}
	say = sayExe.Process
	return say
}

// helper functions

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
