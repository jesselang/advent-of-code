package main

import "fmt"
import "os"
import "log"
import "bufio"
import "strconv"

func fuel_for_mass(mass int) int {
	return mass/3.0 - 2
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		mass = fuel_for_mass(mass)
		for mass > 0 {
			total += mass
			mass = fuel_for_mass(mass)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}
