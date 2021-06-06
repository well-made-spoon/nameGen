# nameGen
A CLI and library for generating random names in the form of "Adjective Noun"

# Install

1. Download binary

2. tar -xvzf nameGen-v0.2-linux-amd64.tar.gz

3. sudo cp ./nameGen /usr/local/bin/

# Usage

## CLI
```
./nameGen  --help
Usage of ./nameGen:
  -flippy
    	Print random name(s) using Flippy Text
  -n int
    	Number of random names to generate (default 1)
```

## Library
```
package main

import (
  "github.com/well-made-spoon/nameGen/RandomName"
)

func main() {
  fmt.Println(nameGen.RandomName())
}
```
