# Rudimentary Drum Machine

This is an exercise in building a rudimentary yet functional drum machine.

## Generate Song

This package will generate a song json file representation with the command below.

```bash
    cd ./generate
    go run genrateSong.go
```

The package has been built in a way that allows it to be easily extendable. If someone wanted to generate a different beat they would just need to setup the array of tracks(each with their name, id, and steps) and pass them into the `buildSong` method.

## Play Song

This package will play a song json file with the command below.

```bash
    cd ./play
    go run play.go <Name of Song>
```
