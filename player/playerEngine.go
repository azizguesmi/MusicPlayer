package player

import (
	"errors"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Player struct {
	file     *os.File
	done     chan bool
	Ctrl     *beep.Ctrl
	Streamer beep.StreamSeekCloser
	Format   beep.Format
}

// create a new Player
func NewPlaeyr(f string) (*Player, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, errors.New("error reading the file")
	}
	streamer, format, err := mp3.Decode(file)
	InitSpeaker(format.SampleRate)
	if err != nil {
		return nil, errors.New("error creating the stream")
	}
	c := beep.Ctrl{Streamer: streamer}
	return &Player{
		file:     file,
		Streamer: streamer,
		Format:   format,
		Ctrl:     &c,
		done:     make(chan bool, 1),
	}, nil
}

func (p *Player) Play() error {
	p.Ctrl.Paused = false
	speaker.Play(beep.Seq(p.Ctrl, beep.Callback(func() { p.done <- true })))
	return nil
}

// pause the truck
func (p *Player) Pause() {
	speaker.Lock()
	defer speaker.Unlock()
	p.Ctrl.Paused = true
}

func (p *Player) IsPaused() bool {
	return p.Ctrl.Paused
}

// resume the truck
func (p *Player) Resume() {
	speaker.Lock()
	defer speaker.Unlock()
	p.Ctrl.Paused = false
}

func (p *Player) Wait() {
	<-p.done
}

func (p *Player) Close() {
	p.Streamer.Close()
	p.file.Close()
}

func InitSpeaker(sampleRate beep.SampleRate) error {
	return speaker.Init(sampleRate, sampleRate.N(time.Second/10))
}

func (p *Player) LaunchSong() {
	p.Play()
	p.Wait()
	p.Close()
}

func (p *Player) IsClosed() bool {

	select {
	case <-p.done:
		return true
	default:
		return false
	}
}
