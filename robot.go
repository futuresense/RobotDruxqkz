package robotdruxqkz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/argusdusty/Ferret"
	si "github.com/bitly/go-simplejson"
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

type Robot struct {
	name string
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
	basic []byte
	crazy string
}

func turnOnTheRobot() {

	config := jconfig.LoadConfig("/robot/robot.conf")

	if data, ok := js.Get("about").CheckGet("i_like"); ok {
		fmt.Println(data)
	}
	if ok {
		p("yay")
	}
	cycleMode()

}

func (r *robotThoughts) question(a robotThoughts) *robotThoughts {
	robot := &Robot{}
	p("What is", robot.name)
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

	p(qse.Query(robot.name, 4))
	r.crazy = "woahhh"

	return r
}
