# audacity-go

This go package lets you control Audacity by using scripting/macros.

## Usage

```go
import audacity "github.com/MarkRosemaker/audacity-go"
```

After your program is done using all the Audacity commands, make sure to call `ClosePipes()` to close the pipes to and from Audacity.

## Example

```go
defer audacity.ClosePipes()

// some code, e.g. a loop

path := "D:/audio/myfile.m4a"

audacity.RemoveTracks() // clear the project
audacity.Import(path) // import the file
audacity.SelectAll() // select all for noise reduction to be applied
audacity.NoiseReduction() // apply noise reduction
audacity.Normalize() // normalize the audio

target := strings.Replace(path, ".m4a", ".mp3", 1)
audacity.Export(target) // export the project to mp3 (or whatever extension target has)
```

## Progress

This repository is not fully done nor sufficiently tested.

To be done:

- [x] connecting to pipes on a windows system
- [x] implementation of universal `Command(cmd string, args ...string) Response` function
- [x] implementation of basic commands
- [ ] implementation of more commands
- [ ] automatic starting of Audacity if it's not running/pipes are not found
- [ ] testing and implementation with non-windows system

Pull requests are very welcome!
