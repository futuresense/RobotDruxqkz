package main

import (
	"os/exec"
)

func surfing() {
	p("surfing")
	script, err := exec.LookPath("youtube-dl")
	if err != nil {
		p("no youtube-dl, please install this python script")
		scriptexe := exec.Command(script, "sdffd")
		scriptexe.Start()
	}

}
