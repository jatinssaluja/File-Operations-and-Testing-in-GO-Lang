package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFromFile(fileName string) []string {

	fileData := []string{}

	inFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fileData = append(fileData, line)

	}

	inFile.Close()

	return fileData

}

func CreateMatrices(fileData1 []string, fileData2 []string) ([][]int, [][]int) {

	matrix1 := [][]int{}
	matrix2 := [][]int{}

	temp1 := []int{}
	temp2 := []int{}

	// Load fileData1 into matrix1
	for _, line := range fileData1 {

		numbers := strings.Split(line, " ")

		for _, number := range numbers {

			intnumber, _ := strconv.Atoi(number)
			temp1 = append(temp1, intnumber)

		}

		matrix1 = append(matrix1, temp1)
		temp1 = nil

	}

	// Load fileData2 into matrix2
	for _, line := range fileData2 {

		numbers := strings.Split(line, " ")

		for _, number := range numbers {

			intnumber, _ := strconv.Atoi(number)
			temp2 = append(temp2, intnumber)

		}

		matrix2 = append(matrix2, temp2)
		temp2 = nil

	}

	return matrix1, matrix2

}

func DotProduct(matrix1 [][]int, matrix2 [][]int) [][]int {

    if len(matrix1[0]) != len(matrix2){
    
    fmt.Println("Error:", "Number of columns of the matrix1 is not equal to the number of rows of the matrix2")
		os.Exit(1)
    
    } 

	matrix3 := make([][]int, len(matrix1))

	for i := range matrix3 {
		matrix3[i] = make([]int, len(matrix2[0]))
	}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix1[0]); k++ {
				matrix3[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return matrix3

}

func main() {

	fileData1 := []string{}
	fileData2 := []string{}
	
	scanner := bufio.NewScanner(os.Stdin)
  
    fmt.Print("Enter the name of file1: ")
    scanner.Scan()
    file1 := scanner.Text()

    fmt.Print("Enter the name of file2: ")
    scanner.Scan()
    file2 := scanner.Text()

	// Read from file1
	fileData1 = ReadFromFile(file1)

	// Read from file2
	fileData2 = ReadFromFile(file2)

	//Create matrices from File Data
	matrix1, matrix2 := CreateMatrices(fileData1, fileData2)

	//Calculate Dot Product of matrix1 and matrix2
	matrix3 := DotProduct(matrix1, matrix2)

	fmt.Println("Dot Product:", matrix3)

}
