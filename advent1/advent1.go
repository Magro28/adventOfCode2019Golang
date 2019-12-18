package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []string {
	var modules = make([]string, 1, 1)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		modules = append(modules, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return modules
}

func calculateFuel(mass int) int {
	return int(float64(mass)/float64(3)) - 2
}

func part1(modules []string){
	var sum int = 0
	for i := 0; i < len(modules); i++ {
		if modules[i] == "" || len(modules[i]) == 0 {
			continue
		}
		val, err := strconv.Atoi(modules[i])
		if err != nil {
			fmt.Println(" Error while converting string to int", err)
		}
		fuel := calculateFuel(val)
		sum += fuel

		fmt.Println("Actual fuel requirements for module ", i, "with mass", modules[i], "FUEL: ", fuel);
	}
	fmt.Println("Part1: Required total fuel", sum);
}

func part2(modules []string){
	var sum int = 0
	for i := 0; i < len(modules); i++ {
		if modules[i] == "" || len(modules[i]) == 0 {
			continue
		}
		val, err := strconv.Atoi(modules[i])
		if err != nil {
			fmt.Println(" Error while converting string to int", err)
		}
		fuel := calculateFuel(val)
		fmt.Println("Actual fuel requirements for module ", i, "with mass", modules[i], "FUEL: ", fuel);
		sum += fuel
		fuelForFuel := 0
		recursiv_counter := 0
	    for {
			recursiv_counter++;
			fuel = calculateFuel(fuel)
			if(fuel<=0){
				break;
			}
			fmt.Println("Fuel requirements of fuel for module:",modules[i], "recursion", recursiv_counter, "fuel:", fuel);
			fuelForFuel+=fuel;
		}
		fmt.Println("Fuel requirements for extra fuel weight:", fuelForFuel);
		sum += fuelForFuel


	}
	fmt.Println("Part2: Required total fuel", sum);
}

func main() {
	fmt.Println("Advent 1")
	var modules = readFile("advent1/input")
	fmt.Println(modules, len(modules))
	part1(modules)
	part2(modules)

}
