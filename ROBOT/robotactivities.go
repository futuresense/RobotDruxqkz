package main

type robotAction interface {
	trolling()
	surfing()
	chilling()
}

type RobotActivities struct {
	Activity_level int
}

func (r *RobotActivities) trolling() {
	p("haha") //TODO
}
