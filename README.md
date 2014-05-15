## What
Cross-platform utility to strip ANSI escape codes from text STDIN or a named file.

## Usage

![screenshot](http://i.imgur.com/zMZNREr.png)

#### Streaming
`cat file_with_ansi_text | deansify`

#### By filename
`deansify file_with_ansi_text`

## Installation
#### Binary download
[Linux, OSX, and Windows binaries](https://github.com/dpritchett/deansify/releases) available on GitHub.

#### Go users
```
go get github.com/dpritchett/deansify
go install -u github.com/dpritchett/deansify
```
