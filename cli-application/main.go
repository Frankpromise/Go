package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strings"
	"os"
)

func main() {
	entry := flag.String("Entry", "1", "log entry to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() 

	bufReader := bufio.NewReader(f)
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *entry) {
			fmt.Println(line)
		}
	}
}