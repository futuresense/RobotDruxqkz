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
	Musical_tastes string
	Startup_phrase string
}

var (
	quirky = &Robot{Name: "Diggity Dan", Musical_tastes: "Rock, Opera, Blues"}
)

//Check robotpreferences.json exists

func init() {

	filebyte, err := ioutil.ReadFile("robotpreferences.json")
	if err != nil {
		fmt.Println("Hold on, need to do some setup to use this Robot.")
		checkForEditor()
		decideRobotPersonality()
	}
	var r Robot
	json.Unmarshal(filebyte, &r)
	fmt.Printf("Name: %s \nMusical Tastes: %s", r.Name, r.Musical_tastes)
}

func decideRobotPersonality() {
	fmt.Println(`"(*&(*&()*&@()&*@()*&@()*&@()&*@(MAGICAWESOMEPOWER))))"`)

}

func checkForPrefsFile() {

	f, _ := os.Create("./robotpreferences.json")
	defer f.Close()
	r := Robot{"Bob", "Rock", "Hi there"}
	b, _ := json.Marshal(r)

}
func checkForEditor() {
	subl, err := exec.LookPath("subl")
	if err != nil {
		fmt.Println("Please symlink sublime text to command line, \"subl\", then restart this program. %s", subl)
	}

	fmt.Println(b)

}
func getLocalSongsAndProcess() {
	//Todo

}
