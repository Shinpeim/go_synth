package main

import (
	"code.google.com/p/portaudio-go/portaudio"
)

type playable interface {
	Process(out [][]float32)
}

func play(handler playable, c chan interface{}) {
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, 0, handler.Process)
	panicIfError(err)
	defer stream.Close()

	panicIfError(stream.Start())
	defer stream.Stop()

	<-c
}
