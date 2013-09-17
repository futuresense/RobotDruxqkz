package main

import (
	"github.com/codegangsta/cli"
	//"github.com/futuresense/druxqkz/reddit""
	"github.com/benmanns/goworker"

	"bufio"
	"fmt"
	"os"
)

var (
	rp robotPreferences
	ra robotActivities
	on bool = false
	p       = fmt.Println
)

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error", err)
	}

	app := cli.NewApp()
	app.Name = "Robot Druxqkz\n"
	app.Usage = "This robot, while probably drunk, and mostly just a lazy web surfing monkey, reading reddit, watching movies, chillin to great music.. Is actually just about the coolest Robot ever....Press Q to quit or X to exit at anytime."
	app.Version = "0.1.awesome"
	app.Commands = []cli.Command{
		{
			Name:      "preferences",
			ShortName: "p",
			Usage:     "Show your robot preferences",
			Action: func(c *cli.Context) {
				rp.showPrefs()
				//Robot.getRobotPrefs()
			},
		},
		{
			Name:      "edit preferences",
			ShortName: "e",
			Usage:     "Edit your robot preferences",
			Action: func(c *cli.Context) {
				rp.showPrefs()
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
				ra.robotCycle()
			},
		},
		{
			Name:      "trolling",
			ShortName: "t",
			Usage:     "Definitely not polite dinner conversation.",
			Action: func(c *cli.Context) {
				ra.trolling()
			},
		},
		{
			Name:      "surfing",
			ShortName: "s",
			Usage:     "Surfing the Web.",
			Action: func(c *cli.Context) {
				ra.surfing()
			},
		},
		{
			Name:      "resting",
			ShortName: "r",
			Usage:     "Listening to Tunes.",
			Action: func(c *cli.Context) {
			},
		},
	}
	app.Run(os.Args)

}

func init() {
	p("clikkk")
	turnOnTheRobot()
	goworker.Register("awesome", dumbFunc)

	go func() {

		r := bufio.NewReader(os.Stdin)
		var inputString string
		//always watch for ESC
		for {

			inputBytes, _, _ := r.ReadLine()
			inputString = string(inputBytes)
			if (inputString == "x") || (inputString == "q") {
				os.Exit(0)
			}
		}

	}()
}
func dumbFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v", queue, args)

	return nil
}
