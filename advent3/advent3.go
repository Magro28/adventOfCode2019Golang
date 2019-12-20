package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(path string) ([]string, []string) {
	var wire1 = make([]string, 1, 1)
	var wire2 = make([]string, 1, 1)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	wire1 = append(wire1, strings.Split(scanner.Text(), ",")...)
	scanner.Scan()
	wire2 = append(wire2, strings.Split(scanner.Text(), ",")...)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wire1, wire2
}

func main() {
	fmt.Println("Advent 3")
	var wire1, wire2 = readFile("advent3/input")

	fmt.Println("Wire 1", len(wire1))
	for i := 0; i < len(wire1)-1; i++ {
		if wire1[i] == "" || len(wire1[i]) == 0 {
			continue
		}
		fmt.Print(" ", wire1[i])
	}
	fmt.Println()
	fmt.Println("****")
	fmt.Println("Wire 2", len(wire2))
	for i := 0; i < len(wire2)-1; i++ {
		if wire2[i] == "" || len(wire2[i]) == 0 {
			continue
		}
		fmt.Print( " ", wire2[i])
	}
	//TODO work in progress


}
