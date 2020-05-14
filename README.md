# audacity-go

This go package lets you control Audacity by using scripting/macros.

## Usage

```go
import audacity "github.com/MarkRosemaker/audacity-go"
```

After you've used the commands, make sure

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
