# EL
⛱A Go log lib.

## Feature
- Simple API.
- Level logging.
- Singleton.

## Install
```bash
$ go get -u -v github.com/BouncyElf/el
```

## Example
```Go
package main

import "github.com/BouncyElf/el"

func main() {
		el.Info("info from el", el.Map{
				"info":"hello, world",
		})
		el.Warn("should not be here")
		el.Panic("panic from el")
}
```

## Doc
[Doc Here](https://godoc.org/github.com/BouncyElf/el)

