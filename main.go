package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Song struct {
	name            string
	durationSeconds int
}

func main() {
	fmt.Println("Hello World!")
	fileContents := readFile()
	fmt.Println(parseSongs(fileContents))
}

func readFile() []string {
	file, err := os.Open("songs_with_durations.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseSongs(lines []string) []Song {
	var songs = make([]Song, len(lines))
	for i := 0; i < len(lines); i++ {
		songs[i] = parseSong(lines[i])
	}
	return songs
}

func parseSong(line string) Song {
	fields := strings.Split(line, "\t")
	durationSeconds, _ := strconv.Atoi(fields[1])
	return Song{fields[0], durationSeconds}
}
