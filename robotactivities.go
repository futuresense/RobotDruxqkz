package robotdruxqkz

import (
	//"fmt"
	"os"
	"os/exec"
	//	"github.com/futuresense/linkerator"
	"syscall"
	"time"
)

type robotActivities interface {
	trolling()
	surfing()
	chilling()
}

type RobotAction struct {
	Sneakometer    interface{}
	Clumsoclock    int64
	Activity_level int64
	P/*reciseotron*/ int64
}

func (r *RobotAction) trolling() {
	//todo
}

func (r *RobotAction) coverArt() {
	//
}

func cycle() {
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

type SurfingRequest chan struct{}

func (r *RobotAction) surfing() {
	go func() {
		for {
			responseChan := make(SurfingRequest)
		}
	}
	//check for scripts
	var scripts = make
	for _, tag := range sr.Tag {
		args := []string{"youtube-dl", "-x", sr.Tag}
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			panic(execErr)
		}
		p(tag)
	}

}

func surfingInit() {
	title := make[]
	scripts := make([][]string, 3)
for i := 0; i < count; i++ {
	scripts[i]
}
	
	env := os.Environ()
		ydl, err := exec.LookPath(script)

		if err != nil {
			p("no, please install this python script", script)
		}
}

twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}