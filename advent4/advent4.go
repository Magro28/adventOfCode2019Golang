package main

import (
	"fmt"
	"strconv"
)

type Range struct {
	from, to int
}



func checkDoubleDigits(code int)bool{
	codeStr := strconv.Itoa(code)
	for i:=0; i< len(codeStr)-1;i++{
		if (codeStr[i]==codeStr[i+1]){
			return true
		}
	}
	//fmt.Println(code, "no duplicate!")
	return false
}

func checkNeverDecrease(code int)bool{
	codeStr := strconv.Itoa(code)
	for i:=0; i< len(codeStr)-1;i++{
		if (codeStr[i]>codeStr[i+1]){
			//fmt.Println(code, "decreases!")
			return false
		}
	}
	return true
}

func checkNoGroupOfMatchingDigits(code int)bool {
	codeStr := strconv.Itoa(code)
	isNotPartOfLargerGroup := false;
	for i:=0; i< len(codeStr)-1;i++{

		count := checkGroup(codeStr, codeStr[i])
		if count==2 {
			isNotPartOfLargerGroup = true
		}
	}

	return isNotPartOfLargerGroup
}

func checkGroup(codeStr string, value uint8) int {
	count:=1;
	for i:=0; i< len(codeStr)-1;i++{
		if codeStr[i] == codeStr[i+1] && codeStr[i]==value {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("Advent 4")
	var r = Range{183564,657474}
	count := 0
	fmt.Println("Part 1")
	for i:=r.from;i<=r.to;i++{
		if (checkDoubleDigits(i)&&checkNeverDecrease(i)){
			fmt.Println("Found valid code:", i)
			count++
		}
	}
	fmt.Println("Found", count,"valid codes")
	fmt.Println("Part 2")
	count = 0
	fmt.Println("Part 1")
	for i:=r.from;i<=r.to;i++{
		if (checkDoubleDigits(i)&&checkNeverDecrease(i)&&checkNoGroupOfMatchingDigits(i)){
			count++
			fmt.Println("Found valid code:", i)
		}
	}
	fmt.Println("Found", count,"valid codes")
}


