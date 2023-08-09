package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//fmt.Println(readFile("./input.txt"))
	input, _ := readFile("./input.txt")
	fmt.Println(maxCalories(input))

}
func maxCalories(cals []int) (int) {

	tmp := 0
	for i := 0; i < len(cals); i++ {
		if cals[i] > tmp {
			tmp = cals[i]
			i++
		} 
	}

	return tmp
}

func readFile(file string) ([]int, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {continue}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

/*
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	//Read input file
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	//Search for the maximum sum of calories
	maxCalories := 0
	currentElfCalories := 0

	for sc.Scan(){
		snack, err := strconv.Atoi(sc.Text())
		currentElfCalories += snack

		//If error is different from nil then I found an empty line
		if err != nil{
			if currentElfCalories>maxCalories{
				maxCalories = currentElfCalories
			}
			//I start with a new elf
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories)
}
*/
