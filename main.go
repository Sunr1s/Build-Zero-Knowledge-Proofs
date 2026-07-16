package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// TODO (what-is-zk): implement per the lesson description.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		b := math.Pow(0.5, float64(n))
		fmt.Printf("cheat_probability=%f\n", b)
	}
}
