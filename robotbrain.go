package main

import (
	"fmt"
	json "github.com/bitly/go-simplejson"
	"io/ioutil"
	//"log"
)

var (
	ok  bool
	err error
)

type robotBrain interface {
	organize()
	think()
	iterate()
	question()
}

type robotThoughts struct {
	name string
}

func turnOnTheRobot() {
	p("troppo")
	filebyte, _ := ioutil.ReadFile("robotbrain.json")
	if err != nil {
		p("no robotbrain file, so creating a new one.")
	}

	js, err := json.NewJson(filebyte)
	if err != nil {
		p("Can't read robotbrain file.")
	}

	if data, ok := js.Get("about").CheckGet("i_like"); ok {
		fmt.Println(data)
	}
	if ok {
		p("yay")
	}
}

func (r *robotThoughts) question() {

}
