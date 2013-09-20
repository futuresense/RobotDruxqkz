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
	troll()
	surfing()
	chilling()
}

type RobotAction struct {
	Sneakometer    interface{}
	Clumsoclock    int64
	Activity_level int64
	P/*reciseotron*/ int64
}

func (r *RobotAction) troll() {
	//todo
}

func Question([]string) {
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
	}()
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

func init() {
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
