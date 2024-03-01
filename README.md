# README for Word, Line, and Character Counter

This program counts words, lines, and characters in a text file. It's a command-line tool that lets you specify what to count using flags.

## Features
- **Line count**: Counts lines in a file.
- **Word count**: Counts words in a file.
- **Character count**: Counts characters in a file.
- **Flexible input**: Choose counts via command-line flags.

## How to Use

### Compile
First, compile the Go source code:
```sh
go build
```

### Run
Execute the program with your desired flags:
- Count lines: `-l`
- Count words: `-w`
- Count characters: `-c`

### Examples
Count lines in `file.txt`:
```sh
./program -l file.txt
```
Count words in `file.txt`:
```sh
./program -w file.txt
```
Count characters in `file.txt`:
```sh
./program -c file.txt
```
Get all counts for `file.txt`:
```sh
./program file.txt
```

## Error Handling
Errors, such as file not found, will terminate the program with a message.

## Requirements
- Go 1.13 or newer.

## License
Free and open-source.