package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// checkRange checks whether value lies in [0, 2^bits).
func checkRange(value *big.Int, bits int) {
	limit := new(big.Int).Lsh(big.NewInt(1), uint(bits))
	if value.Sign() >= 0 && value.Cmp(limit) < 0 {
		fmt.Printf("IN_RANGE bits=%d\n", bits)
		return
	}
	max := new(big.Int).Sub(limit, big.NewInt(1))
	fmt.Printf("OUT_OF_RANGE max=%s\n", max.String())
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) != 3 || parts[0] != "RANGE" {
			continue
		}

		value, ok := new(big.Int).SetString(parts[1], 10)
		if !ok {
			continue
		}

		bits, err := strconv.Atoi(parts[2])
		if err != nil || bits < 0 {
			continue
		}

		checkRange(value, bits)
	}
}
