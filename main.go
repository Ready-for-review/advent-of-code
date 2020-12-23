package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func CalculateFuel(moduleMass int) int {
	return moduleMass/3 - 2
}

func SumOfFuels(moduleMasses []int) (fuel int) {
	for _, moduleMass := range moduleMasses {
		fuel += CalculateFuel(moduleMass)
	}
	return
}

// ImportNumberList opens the given file and parses its contents (comma separated numbers) to an int array
func ImportNumberList(file string) ([]int, error) {
	handle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer handle.Close()
	return ToNumberList(handle)
}

// ToNumberList takes a reader and transform the results of it to a []int
func ToNumberList(handle io.Reader) (listOfNumbers []int, err error) {
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		listOfNumbers = append(listOfNumbers, number)
	}
	return listOfNumbers, nil
}

func CalculateFuelFromInputFile(input string) (int, error) {
	moduleMasses, err := ImportNumberList(input)
	if err != nil {
		return 0, err
	}

	return SumOfFuels(moduleMasses), nil
}

func main() {
	solution, err := CalculateFuelFromInputFile("./input.txt")
	if err != nil {
		log.Fatal("Couldn't parse input file")
	}
	fmt.Printf("needed fuel: %d", solution)
}
