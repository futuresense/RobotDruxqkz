package main

import (
	//	"fmt"
	"github.com/codegangsta/cli"
	//"github.com/futuresense/druxqkz/reddit""
	"github.com/futuresense/RobotDruxqkz/talk"
	"os"
)

var (
	on bool = false
)

func main() {
	app := cli.NewApp()
	app.Name = "Robot Druxqkz\n"
	app.Usage = "This robot, while probably drunk, and mostly just a lazy web surfing monkey, reading reddit, watching movies, chillin to great music.. Is actually just about the coolest Robot ever...."
	app.Version = "0.1.awesome"
	app.Commands = []cli.Command{
		{
			Name:      "preferences",
			ShortName: "p",
			Usage:     "Show your robot preferences",
			Action: func(c *cli.Context) {
				getRobotPrefs()
			},
		},
		{
			Name:      "edit preferences",
			ShortName: "e",
			Usage:     "Edit your robot preferences",
			Action: func(c *cli.Context) {
				setRobotPrefs()
			},
		},
		{
			Name:      "help",
			ShortName: "h",
			Usage:     "Help at your service.",
			Action: func(c *cli.Context) {
				println("Help can be found at http://dankspliff.herokuapp.com")
			},
		},
		{
			Name:      "cycling",
			ShortName: "c",
			Usage:     "Put the Robot in Cycling mode. 8 hrs sleeping, 8 hrs music and surfing, 8hrs trolling and reading.",
			Action: func(c *cli.Context) {
				robotCycle()
			},
		},
		{
			Name:      "trolling",
			ShortName: "t",
			Usage:     "Definitely not polite dinner conversation.",
			Action: func(c *cli.Context) {
				trolling()
			},
		},
		{
			Name:      "surfing",
			ShortName: "s",
			Usage:     "Surfing the Web.",
			Action: func(c *cli.Context) {
				talk.Init()
			},
		},
		{
			Name:      "music",
			ShortName: "m",
			Usage:     "Listening to Tunes.",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args()[0])
			},
		},
	}
	app.Run(os.Args)
}

func init() { //fmt.println("on")
}

// helper functions
func onSwitch() bool {
	if on {
		on = false
	} else {
		on = true
	}
	return on
}
