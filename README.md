# go-uniqid

[![Share on X](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Unique%20ID%20generator%20for%20Go%2C%20inspired%20by%20PHP%27s%20uniqid%28%29.%20Goroutine-safe.&url=https://github.com/mintance/go-uniqid&hashtags=go,golang) [![Share on Reddit](https://img.shields.io/badge/share-reddit-orange?logo=reddit&logoColor=white)](https://www.reddit.com/submit?url=https://github.com/mintance/go-uniqid&title=go-uniqid%20-%20Unique%20ID%20generator%20for%20Go)

[![CI](https://github.com/mintance/go-uniqid/actions/workflows/ci.yml/badge.svg)](https://github.com/mintance/go-uniqid/actions/workflows/ci.yml)
[![Tests](https://img.shields.io/badge/tests-9%20passed-brightgreen)](https://github.com/mintance/go-uniqid/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mintance/go-uniqid.svg)](https://pkg.go.dev/github.com/mintance/go-uniqid)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

Unique ID generator for Go, inspired by PHP's `uniqid()`. Goroutine-safe.

## Install

```
go get github.com/mintance/go-uniqid
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/mintance/go-uniqid"
)

func main() {
    fmt.Println(uniqid.New())                              // "6601a3d800001"
    fmt.Println(uniqid.NewWithPrefix("usr_"))               // "usr_6601a3d800002"
    fmt.Println(uniqid.NewWithEntropy())                    // "6601a3d8000035.82749103"
    fmt.Println(uniqid.NewWithPrefixAndEntropy("sess_"))    // "sess_6601a3d8000042.19384756"
}
```

## API

| Function | Description |
|---|---|
| `New()` | 13-char hex ID (timestamp + counter) |
| `NewWithPrefix(prefix)` | Prefixed hex ID |
| `NewWithEntropy()` | 23-char ID with random decimal suffix |
| `NewWithPrefixAndEntropy(prefix)` | Prefixed ID with entropy |

All functions are safe for concurrent use.

## License

[Apache 2.0](LICENSE)

