package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Song struct {
	Name  string   `json:"name"`
	BPM   uint     `json:"bpm"`
	Steps []string `json:"steps"`
}

type Track struct {
	Name  string
	Id    uint
	Steps string
}

func main() {
	fourOnTheFloorTracks := generateFourOnTheFloorTracks()
	song := buildSong(fourOnTheFloorTracks)

	fourOnTheFloor := Song{
		Name:  "FourOnTheFloor",
		BPM:   128,
		Steps: song,
	}

	writeSongToFile(fourOnTheFloor)
}

func generateFourOnTheFloorTracks() []Track {
	bass := Track{
		Name:  "bassdrum",
		Id:    0,
		Steps: "1000100010001000",
	}
	hihat := Track{
		Name:  "hi-hat",
		Id:    1,
		Steps: "0010001000100010",
	}
	snare := Track{
		Name:  "snare",
		Id:    2,
		Steps: "0000100000001000",
	}

	return []Track{bass, snare, hihat}
}

// this should be able to be used to compile any 16 step drum beat
func buildSong(tracks []Track) []string {
	song := []string{"_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "_"}

	for _, track := range tracks {
		//should be 16 for this purpose
		numberOfSteps := len(track.Steps)
		for i := 0; i < numberOfSteps; i++ {
			if string(track.Steps[i]) == "1" {
				if song[i] != "_" {
					song[i] += ("+" + track.Name)
				} else {
					song[i] = track.Name
				}
			}
		}
	}

	return song
}

func writeSongToFile(song Song) {
	var jsonData []byte
	jsonData, _ = json.Marshal(song)
	// file permissions for the output file: equivalent to drw-rw-rw
	const permissions os.FileMode = 0666
	ioutil.WriteFile("../../songs/"+song.Name+".json", jsonData, permissions)
}
