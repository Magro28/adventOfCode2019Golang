package main

import (
	"fmt"
)
var code = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 5, 19, 23, 2, 10, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 35, 5, 39, 2, 6, 39, 43, 1, 43, 5, 47, 2, 47, 10, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 1, 71, 5, 75, 1, 13, 75, 79, 1, 6, 79, 83, 2, 83, 13, 87, 1, 87, 6, 91, 1, 10, 91, 95, 1, 95, 9, 99, 2, 99, 13, 103, 1, 103, 6, 107, 2, 107, 6, 111, 1, 111, 2, 115, 1, 115, 13, 0, 99, 2, 0, 14, 0}

func compute(op int, valAd1 int, valAd2 int) int {

	switch op {
	case 1:
		return add(valAd1,valAd2);
	case 2:
		return mul(valAd1,valAd2);
	case 99:
		return -1;
	default:
		fmt.Println("Error unregcognized value:",op)

	}

	return 1
}

func add(valAd1 int, valAd2 int)int {
	return read(valAd1)+read(valAd2);
}

func mul(valAd1 int, valAd2 int)int {
	return read(valAd1)*read(valAd2);
}

func read(valAd int)int {
	return code[valAd]
}

func write(addr int, val int){
	code[addr]=val
}

func main() {
	fmt.Println("Advent 2")
	for i := 0; i<len(code)-4;i+=4{
		tempSlice := code[i:i+4];
		fmt.Println(i, i+4,"Next computation: ", tempSlice);
		val := compute(tempSlice[0],tempSlice[1],tempSlice[2])
		if val == -1{
			fmt.Println("Programm finished!. Result on index 0 is: ", code[0])
			return
		}
		fmt.Println("Writing result to address:",val,tempSlice[3]);
		write(tempSlice[3], val);
		fmt.Println("code:",code);

	}
}
