package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) != 3 || parts[0] != "EVAL" {
			continue
		}

		x, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		coefficientStrings := strings.Split(parts[1], ",")
		coefficients := make([]int, len(coefficientStrings))
		valid := true

		for i, text := range coefficientStrings {
			coefficients[i], err = strconv.Atoi(text)
			if err != nil {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		result := 0
		for i := len(coefficients) - 1; i >= 0; i-- {
			result = result*x + coefficients[i]
		}

		fmt.Println(result)
	}
}
