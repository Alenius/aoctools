package aoctools

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadNewlineSeparatedFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		new_line := scanner.Text()
		lines = append(lines, new_line)
	}

	return lines
}

func ReadCommaSeparatedFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		new_line := scanner.Text()
		new_line_separated := strings.Split(new_line, ",")
		lines = append(lines, new_line_separated...)
	}

	return lines
}

func GetMedian(intSlice []int) int {
	sliceLength := len(intSlice)
	if (sliceLength % 2) == 0 {
		lowerMiddleIndex := int((float64(sliceLength) / 2) - 1)
		upperIndex := int(float64(sliceLength) / 2)

		average := (intSlice[upperIndex] + intSlice[lowerMiddleIndex]) / 2
		return average

	} else {
		middleIndex := math.Floor((float64(sliceLength) / 2))
		return intSlice[int(middleIndex)]
	}

}

func GetAverage(intSlice []int) int {
	sliceLength := len(intSlice)

	total := 0
	for _, value := range intSlice {
		total += value
	}

	average := int(math.Round(float64(total / sliceLength)))

	return average
}
