package osc

type Osc interface {
	powerInNthSample(n int) float32
	Process(out [][]float32)
}

type squareOsc struct {
	freq       float64
	sampleRate float64
}

func NewSquareOsc(freq float64, sampleRate float64) Osc {
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
