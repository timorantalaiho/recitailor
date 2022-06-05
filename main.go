package main

import (
	"bufio"
	"fmt"
	perm "github.com/gitchander/permutation"
	"os"
	"strconv"
	"strings"
)

type Song struct {
	name            string
	durationSeconds int
}

type SongSlice []Song

func (ss SongSlice) Len() int      { return len(ss) }
func (ss SongSlice) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }

func main() {
	fileContents := readFile()

	songs := parseSongs(fileContents)

	permutations := perm.New(SongSlice(songs))
	for permutations.Next() {
		fmt.Println(songs)
	}
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
