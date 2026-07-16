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
		if len(parts) != 6 {
			continue
		}

		g, ok1 := new(big.Int).SetString(parts[0], 10)
		y, ok2 := new(big.Int).SetString(parts[1], 10)
		p, ok3 := new(big.Int).SetString(parts[2], 10)
		q, ok4 := new(big.Int).SetString(parts[3], 10)
		c, ok5 := new(big.Int).SetString(parts[4], 10)
		z, ok6 := new(big.Int).SetString(parts[5], 10)

		if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 {
			panic("invalid number in header")
		}
		negativeC := new(big.Int).Neg(c)
		negativeC.Mod(negativeC, q)
		gToZ := new(big.Int).Exp(g, z, p)
		yToNegativeC := new(big.Int).Exp(y, negativeC, p)
		r := new(big.Int).Mul(gToZ, yToNegativeC)
		r.Mod(r, p)
		fmt.Println(r)

	}
}
