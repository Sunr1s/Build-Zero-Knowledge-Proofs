package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseCSV(text string) ([]int, error) {
	parts := strings.Split(text, ",")
	intArray := make([]int, 0, len(parts))

	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		intArray = append(intArray, number)
	}
	return intArray, nil

}

func dotProduct(coefficients, witness []int) int {
	result := 0
	for i := 0; i < len(witness); i++ {
		result += coefficients[i] * witness[i]
	}
	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var witness []int
	constraintIndex := 0
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) == 0 {
			continue
		}

		if parts[0] == "WITNESS" {
			if len(parts) != 2 {
				continue
			}

			parsedWitness, err := parseCSV(parts[1])
			if err != nil {
				continue
			}

			witness = parsedWitness
			continue
		}

		if parts[0] == "CONSTRAINT" {
			if len(parts) != 4 || witness == nil {
				continue
			}

			a, err := parseCSV(parts[1])
			if err != nil {
				continue
			}

			b, err := parseCSV(parts[2])
			if err != nil {
				continue
			}

			c, err := parseCSV(parts[3])
			if err != nil {
				continue
			}

			if len(a) != len(witness) ||
				len(b) != len(witness) ||
				len(c) != len(witness) {
				continue
			}
			aResult := dotProduct(a, witness)
			bResult := dotProduct(b, witness)
			cResult := dotProduct(c, witness)

			if aResult*bResult != cResult {
				fmt.Printf("BAD %d\n", constraintIndex)
				return
			}
			// если не равно:  и

			constraintIndex++
		}
	}

	fmt.Println("OK")
}
