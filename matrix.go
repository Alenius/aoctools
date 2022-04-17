package aoctools

import "fmt"

func PrettyPrintMatrix[K any](matrix [][]K) {
	noOfRows := len(matrix)
	for i := 0; i < noOfRows; i++ {
		noOfColumns := len(matrix[i])
		for j := 0; j < noOfColumns; j++ {
			fmt.Print(matrix[j][i])
		}
		fmt.Println("")
	}
}
