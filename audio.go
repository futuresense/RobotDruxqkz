package main

import (
	"fmt"
	"github.com/scottferg/Go-SDL/sdl"
	sdl_audio "github.com/scottferg/Go-SDL/sdl/audio"
	"log"
)

var (
	SampleSize = 2048
)

type Audio struct {
	sample      <-chan int16
	samples     []int16
	sampleIndex int
}

func NewAudio(s <-chan int16) *Audio {

	a := Audio{
		sample: s,
	}
	a.samples = make([]int16, SampleSize)

	return &a
}

func (a *Audio) AppendSample(s int16) {
	a.samples[a.sampleIndex] = s
	a.sampleIndex++

	if a.sampleIndex == SampleSize {
		sdl_audio.SendAudio_int16(a.samples)
		a.sampleIndex = 0
	}
}

func (a *Audio) Run() {
	as := sdl_audio.AudioSpec{
		Freq:        44100,
		Format:      sdl_audio.AUDIO_S16SYS,
		Channels:    1,
		Out_Silence: 0,
		Samples:     uint16(SampleSize),
		Out_Size:    0,
	}

	if sdl_audio.OpenAudio(&as, nil) < 0 {
		log.Fatal(sdl.GetError())
	}

	sdl_audio.PauseAudio(false)

	for {
		select {
		case s := <-a.sample:
			a.AppendSample(s)
		}
	}
}

func (a *Audio) Close() {
	fmt.Println("Closing!")
	sdl_audio.PauseAudio(true)
	sdl_audio.CloseAudio()
}
