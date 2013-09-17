package main

type robotAction interface {
	trolling()
	surfing()
	chilling()
	coverArt()
}

type RobotActivities struct {
	Sneakometer    int64
	Clumsoclock    int64
	Activity_level int64
	P/*reciseotron*/ int64
}

func (r *RobotActivities) trolling() {
	r.P = 17
	r.Activity_level = r.P + r.Sneakometer
}
