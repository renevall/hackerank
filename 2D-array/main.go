package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var sum []int32
	var max int32
	for x := 0; x < len(arr)-2; x++ {
		for y := 0; y < len(arr)-2; y++ {
			sum = append(sum, arr[x][y]+arr[x][y+1]+arr[x][y+2]+arr[x+1][y+1]+arr[x+2][y]+arr[x+2][y+1]+arr[x+2][y+2])
		}
	}
	for k, e := range sum {
		if k == 0 {
			max = e
			continue
		}
		if e > max {
			max = e
		}
	}
	return max
}

func main() {
	arr := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create("out")
	// // stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// var arr [][]int32
	// for i := 0; i < 6; i++ {
	// 	arrRowTemp := strings.Split(readLine(reader), " ")

	// 	var arrRow []int32
	// 	for _, arrRowItem := range arrRowTemp {
	// 		arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
	// 		checkError(err)
	// 		arrItem := int32(arrItemTemp)
	// 		arrRow = append(arrRow, arrItem)
	// 	}

	// 	if len(arrRow) != int(6) {
	// 		panic("Bad input")
	// 	}

	// 	arr = append(arr, arrRow)
	// }

	result := hourglassSum(arr)
	fmt.Println(result)
	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
