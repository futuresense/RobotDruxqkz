package robotdruxqkz

import (
	"bytes"
	//	"encoding/json"
	//"fmt"
	"github.com/argusdusty/Ferret"
	//sj "github.com/bitly/go-simplejson"
	"github.com/stathat/jconfig"
	"io/ioutil"
	"strconv"

	//"log"
)

var (
	ok  bool
	err error
	rt  robotThoughts
	rb  robotBrain
)

type robot struct {
	name          string
	isBored       bool
	isInitialized bool
}
type robotMusic interface {
}

type robotBrain interface {
	organize()
	think()
	iterate()
	question()
}

type robotThoughts struct {
	music [][]byte
	basic [][]byte
	crazy string
}

type RobotThought struct {
	thought *robotThoughts
}

//main thread
func thoughtPool(c chan string, r *robotThoughts) {

}

func turnOnTheRobot() {
	//robot config file
	z := robot{isBored: true, isInitialized: true}
	x := robotThoughts{}
	config := jconfig.LoadConfig("/robot/robot.conf")
	if config.GetString("environment") == "production" {
		continue
		// ...
	}
	thoughts := make(chan string)

	go thoughtPool(thoughts, &x)
	go cycleMode()

}

func (r *robotThoughts) question(a string) *robotThoughts {
	p("What is", a)
	Data, err := ioutil.ReadFile("dictionary.dat")
	if err != nil {
		panic(err)
	}
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	for _, Vals := range bytes.Split(Data, []byte("\n")) {
		Vals = bytes.TrimSpace(Vals)
		WordFreq := bytes.Split(Vals, []byte(" "))
		if len(WordFreq) != 2 {
			continue
		}
		Freq, err := strconv.ParseUint(string(WordFreq[1]), 10, 64)
		if err != nil {
			continue
		}
		Words = append(Words, string(WordFreq[0]))
		Values = append(Values, Freq)
	}
	QuestionSearchEngine := ferret.New(Words, Words, Values, Converter)
	var qse = QuestionSearchEngine

	p(qse.Query(a, 4))
	r.crazy = "woahhh"

	return r
}
