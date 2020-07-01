## What

Cross-platform utility to strip [ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors) from text via `STDIN` or a named file.

## Usage

![screenshot](http://i.imgur.com/1E9Lcnt.png)

#### Streaming
`cat file_with_ansi_text | deansify`

#### By filename
`deansify file_with_ansi_text`

## Installation

Click the repo's [releases](https://github.com/shelltoys/deansify/releases) link.

#### Build it locally
```sh
go get -u  github.com/shelltoys/deansify
go install github.com/shelltoys/deansify/cmd/deansify
```
