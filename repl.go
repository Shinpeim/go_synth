package main

import (
	"./osc"
	"bufio"
	"code.google.com/p/portaudio-go/portaudio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const sampleRate = 44100

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("freq sec: ")
		freq, sec, err := readFreqAndSecFromScanner(scanner)

		if err != nil {
			fmt.Println("invalid format")
			continue
		}

		osc := osc.NewSquareOsc(freq, sampleRate)
		go play(osc, sec)
	}
}

func readFreqAndSecFromScanner(scanner *bufio.Scanner) (freq float64, sec time.Duration, err error) {
	if scanner.Scan() {
		input := scanner.Text()
		words := strings.Split(input, " ")

		if len(words) < 2 {
			return 0, 0, errors.New("invalid size")
		}

		freq, freqErr := strconv.Atoi(words[0])
		if freqErr != nil {
			return 0, 0, freqErr
		}

		sec, secErr := strconv.Atoi(words[1])
		if secErr != nil {
			return 0, 0, secErr
		}

		return float64(freq), time.Duration(sec), nil
	} else {
		return 0, 0, scanner.Err()
	}
}

func play(osc osc.Osc, sec time.Duration) {
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, 0, osc.Process)
	panicIfError(err)
	defer stream.Close()

	panicIfError(stream.Start())
	time.Sleep(sec * time.Second)
	panicIfError(stream.Stop())
}

func panicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
