package asteroids

import (
	"io/fs"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Sound struct {
	file   fs.File
	stream *wav.Stream
	player *audio.Player
}

func NewSoundFromFile(audioContext *audio.Context, path string) (*Sound, error) {
	sound := &Sound{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	sound.file = file

	ws, err := wav.DecodeWithSampleRate(AUDIO_SAMPLE_RATE, file)
	if err != nil {
		return nil, err
	}
	sound.stream = ws

	player, err := audioContext.NewPlayer(ws)
	if err != nil {
		return nil, err
	}
	sound.player = player

	return sound, nil
}

func (sound *Sound) Play() {
	sound.player.Rewind()
	sound.player.Play()
}

func (sound *Sound) Close() error {
	if err := sound.player.Close(); err != nil {
		return err
	}
	if err := sound.file.Close(); err != nil {
		return err
	}
	return nil
}
