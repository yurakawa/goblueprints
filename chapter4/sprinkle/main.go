package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

func main() {
	f, err := os.Open("transforms.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// TODO: ファイルの内容を配列に入れる標準関数はない?
	lines := make([]string, 0, 100)
	transforms := bufio.NewScanner(f)
	for transforms.Scan() {
		lines = append(lines, transforms.Text())
	}
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := lines[rand.Intn(len(lines))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
