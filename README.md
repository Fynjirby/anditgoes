# anditgoes

all your daily notes can be here :)

#### Feel free to contribute!

## Building

-   Install [Go](https://go.dev/) and make sure it's working with `go version`
-   Clone repo
-   Run `go build` in repo directory, then move it to `/usr/local/bin/`

## How does it work?

When you run `anditgoes write <text>` it will automatically create a file `~/Documents/anditgoes/notes.toml` with toml-like syntax like this 👇 but with your <tеxt>
```
[13.05.2025]
should make some cool things tomorrow

[14.05.2025]
today i made some cool things!
```
When you run `anditgoes read <date>` it will show you all notes of the <dаtе> <br>
When you run `anditgoes clear <date>` it will clear all notes of the <dаtе>

## Usage

`anditgoes`

-   `write <text>` to add a note of today with <text>
-   `read <date>` to read all notes of the <date>
-   `clear <date>` to clear all notes of the <date>
-   `help` or `man` to see help message

Note: date format is 02.01.2006
