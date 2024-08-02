package main

type ControlMusic struct {
	shouldStop chan bool
}

func NewControlMusic() *ControlMusic {
	return &ControlMusic{
		shouldStop: make(chan bool),
	}
}
