package view

//go:generate stringer -type=animationState
type animationState int

const (
	aStopped animationState = iota
	aRunning
	aFinishing
)

type animation struct {
	currFrame int
	numFrames int
	loop      bool
	state     animationState
}

func newAnimation(numFrames int, loop bool) *animation {
	return &animation{
		numFrames: numFrames,
		loop:      loop,
	}
}

func (a *animation) Start() {
	a.state = aRunning
}

func (a *animation) Stop() {
	a.state = aFinishing
}

func (a *animation) Update() (animating bool) {
	switch a.state {
	case aRunning:
		if a.loop {
			a.currFrame = (a.currFrame + 1) % a.numFrames
			return true
		}

		if a.currFrame < a.numFrames-1 {
			a.currFrame++
			return true
		}
		a.state = aStopped
		return false

	case aFinishing:
		if a.currFrame < a.numFrames-1 {
			a.currFrame++
			return true
		}
		a.state = aStopped
		return false

	default:
		return false
	}
}

func (a *animation) Value(fudge float32) float32 {
	switch a.currFrame {
	case 0:
		return 0

	case a.numFrames - 1:
		return 1

	default:
		return (float32(a.currFrame) + fudge) / float32(a.numFrames-1)
	}
}
