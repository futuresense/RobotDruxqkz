package main

import (
	"fmt"
	"os"
	"os/exec"

	"syscall"
	"time"
)

type robotActivities interface {
	trolling()
	surfing()
	chilling()
	robotCycle()
}

type RobotAction struct {
	Sneakometer    int64
	Clumsoclock    int64
	Activity_level int64
	P/*reciseotron*/ int64
}

func (r *RobotAction) trolling() {
	r.P = 17
	r.Activity_level = r.P + r.Sneakometer
}

func (r *RobotAction) coverArt() {
	r.P = 17
	r.Activity_level = r.P + r.Sneakometer
}

func (r *RobotAction) robotCycle() {
	p := fmt.Println
	go func() {

		musicSessionTimer := time.NewTimer(time.Minute * 2)

		go ra.surfing()
		<-musicSessionTimer.C
		videoSessionTimer := time.NewTimer(time.Minute * 2)

		go ra.trolling()
		<-videoSessionTimer.C

		p("done")

	}()
}

func (r *RobotAction) surfing() {
	p("surfing")
	binary, err := exec.LookPath("youtube-dl")
	if err != nil {
		p("no youtube-dl, please install this python script")
	}

	env := os.Environ()
	args := []string{"youtube-dl", "-x", sr.Tag}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	for _, tag := range sr.Tag {
		p(tag)
	}

}
