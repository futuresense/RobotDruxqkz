package main

import (
	//	"bufio"
	"encoding/json"

	"fmt"
	"io/ioutil"
	//	"log"
	"os"
	"os/exec"
)

type Robot struct {
	Name           string
	Musical_taste  string
	Startup_phrase string
}

var (
	quirky = Robot{Name: "default", Musical_taste: "rock"}
)

//Check robotpreferences.json exists

func getRobotPrefs() {
	checkForPrefsFile()
	var r Robot
	json.Unmarshal(filebyte, &r)
	// create file, populate JSON ARRAY
	//b, _ := json.Marshal(quirky)
	//os.File.Write(b)

	fmt.Printf("Name: %s \nMusical Tastes: %s", r.Name, r.Musical_taste)
}

func setRobotPrefs() {
	var editorPath string = checkForEditor()
	subl := exec.Command(editorPath, "./robotpreferences.json")
	subl.Start()
}

func checkForPrefsFile() bool {
	filebyte, err := ioutil.ReadFile("robotpreferences.json")
	if err != nil {
		f, _ := os.Create("./robotpreferences.json")
		defer f.Close()
	}
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
func getLocalSongsAndProcess() {
	//Todo

}

func robotCycle() {
	//youtubeDlPlaylist := exec.Command(youtube-dl-playlist, ...)
}
