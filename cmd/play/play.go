package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

// STEPS - number of steps
const STEPS = 16

// Song - representation of generated song
type Song struct {
	Name  string
	BPM   uint
	Steps []string
}

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Print("Must provide name of song!")
		os.Exit(1)
	}

	//TODO scan songs directory for names that are a similar match and make a suggestion if exact does not exist
	song := readSongFile(args[0])
	playSong(song)
}

func playSong(song Song) {

	//time per step in nanoseconds
	stepTimeNanos := math.Round(float64(60000000000 / song.BPM))
	currentStep := 0
	ticker := time.NewTicker(time.Duration(stepTimeNanos) * time.Nanosecond)
	done := make(chan bool)

	fmt.Println("Now Playing: " + song.Name)
	fmt.Println(fmt.Sprintf("BPM: %d", song.BPM))

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Print("|" + song.Steps[currentStep] + "|")
				currentStep++
			}
		}
	}()

	time.Sleep(time.Duration(16*stepTimeNanos) * time.Nanosecond)
	ticker.Stop()
	done <- true
	fmt.Println("")
}

func readSongFile(name string) Song {

	jsonFile, err := os.Open("../../songs/" + name + ".json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var song Song
	json.Unmarshal(byteValue, &song)
	return song
}
