package gunthur

import (
	"time"
)

// Animation represents info needed to render a Sprites animation
type Animation struct {
	Duration    time.Duration
	StartFrameX int
	StartFrameY int
	TotalFrames int
	MirrorX     bool
	MirrorY     bool
}

func NewAnimation(duration int, startFrameX int, startFrameY int, totalFrames int, mirrorX bool, mirrorY bool) *Animation {
	return &Animation{
		Duration:    time.Millisecond * time.Duration(duration),
		StartFrameX: startFrameX,
		StartFrameY: startFrameY,
		TotalFrames: totalFrames,
		MirrorX:     mirrorX,
		MirrorY:     mirrorY,
	}
}
