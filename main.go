package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) != 7 || parts[0] != "VERIFY" {
			continue
		}

		generator, ok1 := new(big.Int).SetString(parts[1], 10)
		primeModulus, ok2 := new(big.Int).SetString(parts[2], 10)
		publicValue, ok3 := new(big.Int).SetString(parts[3], 10)
		commitment, ok4 := new(big.Int).SetString(parts[4], 10)
		challenge, ok5 := new(big.Int).SetString(parts[5], 10)
		response, ok6 := new(big.Int).SetString(parts[6], 10)
		if !(ok1 && ok2 && ok3 && ok4 && ok5 && ok6) {
			continue
		}

		left := new(big.Int).Exp(generator, response, primeModulus)
		publicPart := new(big.Int).Exp(publicValue, challenge, primeModulus)
		right := new(big.Int).Mul(commitment, publicPart)
		right.Mod(right, primeModulus)

		if left.Cmp(right) == 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("BAD")
		}
	}
}
