package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// TODO (what-is-zk): implement per the lesson description.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) == 0 {
			continue
		}
		if parts[0] != "ROUNDS" || len(parts) != 2 {
			continue
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		b := math.Pow(0.5, float64(n))
		fmt.Printf("cheat_probability=%f\n", b)
	}
}
