package main

type robotAction interface {
	trolling()
	surfing()
	chilling()
	coverArt()
}

type RobotActivities struct {
	Activity_level int
}

func (r *RobotActivities) trolling() {
	p("haha") //TODO
}
