# Rudimentary Drum Machine

This is an exercise in building a rudimentary yet functional drum machine in Go.

## Generate Song

The generate package will generate a song json file representation with the command below. (Four On The Floor in this example)

```bash
    cd ./generate
    go run generate.go
```

The generate module has been built in a way that allows it to be easily extendable. If someone wanted to generate a different beat they would just need to setup the array of tracks(each with their name, id, and steps) and pass them into the `buildSong` method. It could also be extended in other ways such as outputting different file formats. JSON representation was chosen for songs because of its ease of use, flexibility, and readability.

## Testing Generate Song

This runs some song generations and checks that the steps are what would be expected in the output.

```bash
    cd ./generate
    go test
```

## Play Song

This package will play a song json file with the command below.

```bash
    cd ./play
    go run play.go FourOnTheFloor
```
