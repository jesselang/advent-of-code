package main

import "fmt"
import "log"
import "os"
import "encoding/csv"
import "strconv"

const (
	op_add      = 1
	op_multiply = 2
	op_halt     = 99
)

func compute(stream []int, offset int) int {
	switch opcode := stream[offset]; opcode {
	case op_halt:
		return -1
	case op_add:
		stream[stream[offset+3]] = stream[stream[offset+1]] + stream[stream[offset+2]]
	case op_multiply:
		stream[stream[offset+3]] = stream[stream[offset+1]] * stream[stream[offset+2]]
	}

	return offset + 4
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	stringStream, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	intStream := make([]int, len(stringStream))

	for i, v := range stringStream {
		intStream[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	// adjust values to reflect previous state
	intStream[1] = 12
	intStream[2] = 2
	fmt.Println(intStream)

	offset := 0
	for offset >= 0 {
		offset = compute(intStream[:], offset)
	}

	fmt.Printf("%T\n", intStream)
	fmt.Println(intStream)
}
