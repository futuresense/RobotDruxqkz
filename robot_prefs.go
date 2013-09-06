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

var (
	quirky = Robot{Name: "default", Musical_taste: "rock", Startup_phrase: "I love life"}
)

type Robot struct {
	Name           string
	Musical_taste  string
	Startup_phrase string
}

func getRobotPrefs() {
	r := new(Robot)
	filebyte, _ := ioutil.ReadFile("robotpreferences.json")
	json.Unmarshal(filebyte, &r)
}

func setRobotPrefs() {
	var editorPath string = checkForEditor()
	editor := exec.Command(editorPath, "./robotpreferences.json")
	editor.Start()
}

//Check robotpreferences.json exists
func checkForPrefsFile() {
	filebyte, err := ioutil.ReadFile("robotpreferences.json")
	if err != nil {
		//fill with default's
		data, _ := json.Marshal(quirky)
		fmt.Println(data)
		ioutil.WriteFile("robotpreferences.json", data, 774)

	}
	var r Robot
	json.Unmarshal(filebyte, &r)
	fmt.Printf("Name: %s \nMusical Tastes: %s", r.Name, r.Musical_taste)
}

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
func robotCycle() {
	p := fmt.Println
	p(time.Now())
	//youtubeDlPlaylist := exec.Command(youtube-dl-playlist, ...)
}
