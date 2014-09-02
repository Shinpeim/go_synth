package main

import (
	"bufio"
	"code.google.com/p/portaudio-go/portaudio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const sampleRate = 44100

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	scanner := bufio.NewScanner(os.Stdin)

	chanMap := map[string]chan interface{}{}

	for true {
		fmt.Print("command?: ")
		command, arg, err := readCommand(scanner)

		if err != nil {
			fmt.Println("invalid format")
			continue
		}

		switch command {
		case "play":
			err := playOsc(arg, chanMap)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "stop":
			err := stopOsc(arg, chanMap)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func readCommand(scanner *bufio.Scanner) (string, string, error) {
	if !scanner.Scan() {
		return "", "", scanner.Err()
	}

	input := scanner.Text()
	words := strings.Split(input, " ")

	if len(words) < 2 {
		return "", "", errors.New("invalid size")
	}

	command := words[0]
	arg := words[1]

	return command, arg, nil
}

func playOsc(arg string, chanMap map[string]chan interface{}) error {
	args := strings.Split(arg, ":")

	if len(args) != 2 {
		return errors.New("invalid argument in play")
	}

	key := args[0]

	_, present := chanMap[key]
	if present {
		return errors.New(fmt.Sprintf("key: %s is already played", key))
	}

	freq, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("usage > play key:freq")
		return err
	}

	osc := NewSquareOsc(freq, sampleRate)
	c := make(chan interface{})
	go play(osc, c)

	chanMap[key] = c

	return nil
}

func stopOsc(arg string, chanMap map[string]chan interface{}) error {
	c, present := chanMap[arg]
	if present {
		c <- nil
		return nil
	} else {
		return errors.New("no such key")
	}
}

func panicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
