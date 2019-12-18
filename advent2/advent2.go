package main

import (
	"fmt"
)

var codeOriginal = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 5, 19, 23, 2, 10, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 35, 5, 39, 2, 6, 39, 43, 1, 43, 5, 47, 2, 47, 10, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 1, 71, 5, 75, 1, 13, 75, 79, 1, 6, 79, 83, 2, 83, 13, 87, 1, 87, 6, 91, 1, 10, 91, 95, 1, 95, 9, 99, 2, 99, 13, 103, 1, 103, 6, 107, 2, 107, 6, 111, 1, 111, 2, 115, 1, 115, 13, 0, 99, 2, 0, 14, 0}

func compute(op int, valAd1 int, valAd2 int, code []int) int {

	switch op {
	case 1:
		return add(valAd1, valAd2, code);
	case 2:
		return mul(valAd1, valAd2, code);
	case 99:
		return -1;
	default:
		fmt.Println("Error unregcognized value:", op)

	}

	return 1
}

func add(valAd1 int, valAd2 int, code []int) int {
	return read(valAd1, code) + read(valAd2, code);
}

func mul(valAd1 int, valAd2 int, code []int) int {
	return read(valAd1, code) * read(valAd2, code);
}

func read(valAd int, code []int) int {
	return code[valAd]
}

func write(addr int, val int, code []int) {
	code[addr] = val
}

func part1(code []int) int {
	for i := 0; i < len(code)-4; i += 4 {
		tempSlice := code[i : i+4];
		fmt.Println(i, i+4, "Next computation: ", tempSlice);
		val := compute(tempSlice[0], tempSlice[1], tempSlice[2], code)
		if val == -1 {
			return code[0]
		}
		fmt.Println("Writing result to address:", val, tempSlice[3]);
		write(tempSlice[3], val, code);
		fmt.Println("code:", code);
	}
	return 0;
}

func part2(code []int)int {
	noun := 0
	verb := 0
	searchOutput := 19690720

	for i := 0; i <= 99; i++ {
		fmt.Println("Noun:", noun, "Verb:", verb)
		output := 0
		for j := 0; j <= 99; j++ {
			noun = i
			verb = j
			//reset code slice
			code = make([]int, len(codeOriginal))
			copy(code, codeOriginal)
			code[1] = noun
			code[2] = verb
			output = part1(code)
			if (output == searchOutput) {
				break;
			}
		}
		if (output == searchOutput) {
			fmt.Println("SearchOutput",searchOutput, "found with noun ", noun, "and verb", verb)
			fmt.Println("The solution is noun", noun, "and verb", verb)
			fmt.Println("100 * noun + verb = ", 100*noun+verb)
			break;
		}
	}
	return 100*noun+verb
}

func main() {
	fmt.Println("Advent 2")
	fmt.Println("Part 1")
	//copy code slice from original
	code := make([]int, len(codeOriginal))
	copy(code, codeOriginal)
	fmt.Println(code)
	part1Solution := part1(code)
	fmt.Println("Part 2")
	fmt.Println(code)
	part2solution := part2(code)

	fmt.Println("Part 1 Solution: Result on index 0 is: ", part1Solution)
	fmt.Println("Part 2 Solution: 100 * noun + verb: ", part2solution)
}
