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

//Check robotpreferences.json exists

func robotPrefs() {

	filebyte, err := ioutil.ReadFile("robotpreferences.json")
	if err != nil {
		fmt.Println("Hold on, need to do some setup to use this Robot.")
		setRobotPrefs()
	}
	var r Robot
	json.Unmarshal(filebyte, &r)
	fmt.Printf("Name: %s \nMusical Tastes: %s", r.Name, r.Musical_tastes)
}

func setRobotPrefs() {
	f, _ := os.Create("./robotpreferences.json")
	defer f.Close()
	r := Robot{"Bob", "Rock", "Hi there"}
	b, _ := json.Marshal(r)
	subl, err := exec.LookPath("subl")
	if err != nil {
		fmt.Println("Please symlink sublime text to command line, \"subl\", then restart this program. %s", subl)
	}

	fmt.Println(b)

}
