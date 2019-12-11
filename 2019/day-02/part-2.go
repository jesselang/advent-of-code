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

	origStream := make([]int, len(intStream))

	copy(origStream, intStream)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(intStream, origStream)
			// adjust values to reflect previous state
			intStream[1] = i
			intStream[2] = j

			offset := 0
			for offset >= 0 {
				offset = compute(intStream[:], offset)
			}

			if intStream[0] == 19690720 {
				fmt.Println(intStream[1], intStream[2])
				break
			}
		}
	}
}
