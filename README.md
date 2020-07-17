# Rudimentary Drum Machine

This is an exercise in building a rudimentary yet functional drum machine in Go.

## Generate Song Module

The generate module will generate a song json file representation with the command below. (Four On The Floor in this example)

```bash
    cd ./cmd/generate
    go run generate.go
```

The generate module has been built in a way that allows it to be easily extendable. If someone wanted to generate a different beat they would just need to setup an array of tracks(each with their name, id, and steps) and pass them into the `buildSong` method. It could also be extended in other ways such as outputting different file formats. JSON representation was chosen for generated songs because of its ease of use, flexibility, and readability.

## Compiling Songs

In the generate module a song is initially represented by tracks. Each track has an associated name, id, and steps string. The steps string is a representation of when the instrument should be played on certain steps. In this example all tracks should have 16 steps therefore all steps strings should be 16 characters long. I chose to indicate the instrument being played with a `1`
and not played with a `0`.

Once the tracks are "compiled" together, with the `buildSong` method, the final representation of the "playable" song has been created in a `[]string`. This is then written to the json file ready to be played by the play song module.

## Testing Generate Song

This runs some song generations and checks that the steps are what would be expected in the output.

```bash
    cd ./cmd/generate
    go test
```

## Play Song Module

This module will play a song json file with the command below.

```bash
    cd ./cmd/play
    go run play.go <Name of Song>
```

### Example:

```bash
    go run play.go FourOnTheFloor
```
