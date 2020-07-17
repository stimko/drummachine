package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const STEPS = 16

type Song struct {
	Name  string
	BPM   uint
	Steps []string
}

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		os.Exit(2)
	}

	song := readSongFile(args[0])

	fmt.Println(song)
}

func playSong(song Song) {

	stepTimeNanos := 60000000000 / song.BPM

	go func() {
		currentStep := 0
		for currentStep < STEPS {
			currentStep++
			fmt.Print("|" + song.Steps[currentStep] + "|")
			time.Sleep(time.Duration(stepTimeNanos) * time.Nanosecond)
		}
	}()

}

func readSongFile(path string) Song {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var song Song

	json.Unmarshal(byteValue, &song)

	return song
}
