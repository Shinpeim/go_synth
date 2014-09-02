package main

import (
	"errors"
	"math"
)

func NewOsc(oscType string, freq float64, sampleRate float64) (playable, error) {
	switch oscType {
	case "sine":
		return NewSineOsc(freq, sampleRate), nil
	case "square":
		return NewSquareOsc(freq, sampleRate), nil
	default:
		return nil, errors.New("unknown osc type")
	}
}

type squareOsc struct {
	freq       float64
	sampleRate float64
}

func NewSquareOsc(freq float64, sampleRate float64) playable {
	return &squareOsc{freq, sampleRate}
}

func (osc *squareOsc) Process(out [][]float32) {
	for i := range out[0] {
		out[0][i] = osc.powerInNthSample(i)
	}
}

func (osc *squareOsc) powerInNthSample(n int) float32 {
	leng := int(osc.sampleRate / osc.freq)

	if n%int(leng) < int(leng)/2 {
		return 0.1
	} else {
		return -0.1
	}
}

type sineOsc struct {
	step  float64
	phase float64
}

func NewSineOsc(freq float64, sampleRate float64) playable {
	return &sineOsc{freq / sampleRate, 0}
}

func (osc *sineOsc) Process(out [][]float32) {
	for i := range out[0] {
		out[0][i] = osc.powerInNthSample(i)
	}
}

func (osc *sineOsc) powerInNthSample(n int) float32 {
	currentPhase := osc.phase
	_, osc.phase = math.Modf(osc.phase + osc.step)
	return float32(math.Sin(2 * math.Pi * currentPhase))
}
