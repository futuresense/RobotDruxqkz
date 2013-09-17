package main

import (
	//	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//	"log"
	//	"os"
	"os/exec"
)

var ()

type Robot interface {
	initRobot()
	setRobotPrefs()
	showPrefs()
	robotCycle()
}

type RobotStatistics struct {
	Cycle_state    int
	Name           string
	Musical_taste  string
	Startup_phrase string
	Cycle_length   int
	Tag_history    []string
}

func (r *RobotStatistics) showPrefs() {
	p("doh")
}

func (r *RobotStatistics) initRobot() {
	filebyte, _ := ioutil.ReadFile("robotpreferences.json")
	json.Unmarshal(filebyte, &r)
	return
}

func (r *RobotStatistics) setPrefs() {
	var editorPath string = checkForEditor()
	editor := exec.Command(editorPath, "./robotpreferences.json")
	editor.Start()
}

//Check robotpreferences.json exists
func checkForPrefsFile() {
	filebyte, err := ioutil.ReadFile("robotpreferences.json")
	if err != nil {
		//fill with default's
		data, _ := json.Marshal("quirky")
		fmt.Println(data)
		ioutil.WriteFile("robotpreferences.json", data, 774)

	}
	json.Unmarshal(filebyte, &r)
}

//helper function
func checkForEditor() string {
	//HACK
	path, err := exec.LookPath("subl")

	if err != nil {
		path, err = exec.LookPath("EDITOR")
		fmt.Println("Please SET EDITOR global variable. This can be done on the command line. \"subl\", then restart this program. %s", path)
	}
	fmt.Println("editor: ", path)
	return path

}
func (r *RobotStatistics) robotCycle() {
	p := fmt.Println
	go func() {

		musicSessionTimer := time.NewTimer(time.Minute * 2)

		go surfing()
		<-musicSessionTimer.C
		videoSessionTimer := time.NewTimer(time.Minute * 2)

		go trolling()
		<-videoSessionTimer.C
		sleepSessionTimer := time.NewTimer(time.Minute * 2)

		go snoring()
		<-sleepSessionTimer.C
		p("done")

	}()
}
