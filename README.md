# Go Token Counter for OpenAI

This repository contains a Go implementation for counting tokens used in OpenAI API requests. It leverages the [tiktoken-rs](https://github.com/zurawiki/tiktoken-rs) Rust library, which is utilized as a static library and imported into the Go program via cgo.

It's important to note that the current implementation only supports the BPE based on cl100k_base.

## Overview

The Go Token Counter provides an efficient way to count tokens in a text string without making an actual API call to OpenAI. By using the Rust library [tiktoken-rs](https://github.com/zurawiki/tiktoken-rs), it ensures high performance and accurate token counting, which is crucial for managing API usage and costs.

## Generating and compiling the Rust library

The package comes with a precompiled libtiktoken.a static library. However, if you need to recompile the library for a different architecture like amd64 or arm64, follow the instructions below:

### Install Rust
```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

### Generate
```
go generate
```

This command will trigger the necessary steps to generate and compile the Rust library for your specific architecture.

## How to use it

```
package main

import (
	"fmt"
	"github.com/devfullcycle/tiktokengo"
)

func main() {
	prompt := "This is a sample text to count tokens."
	tokenCount := tiktokengo.CountTokens(prompt)
	fmt.Printf("Number of tokens: %d\n", tokenCount)
}
```
