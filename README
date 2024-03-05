# Cron Parser

This repo contains a very simple parser that displays cron-tab information in human understandable format.

## Requirements

- [Go](https://golang.org/doc/install) >= Go 1.22


## Getting started

To run the command, first build the project, then run the command with a single argument.
```shell
go build
~$ cron-parser ＂argument＂
```

Or simply run the command without building it first
```shell
~$ go run main.go ＂argument＂
```


## Usage

Run the command after the project has been built
```shell
~$ cron-parser ＂*/15 0 1,15 * 1-5 /usr/bin/cmd＂
```

Or run the command directly
```shell
~$ go run main.go ＂*/15 0 1,15 * 1-5 /usr/bin/cmd＂
```

This should dispplay the following output
```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```
